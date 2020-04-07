package isc

import (
	"bytes"
	"encoding/hex"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/trap"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

type Context interface {
	Sender() []byte
	Address() []byte
	GetExternalStorageAt(chainID uip.ChainID, typeID uip.TypeID,
		contractAddress uip.ContractAddress, pos []byte, description []byte) (gvm.Ref, error)
}

type ISC struct {
	GVM     *gvm.GVM
	Storage Storage
	Ctx     Context
}

func NewISC(msg Context, storage *storage.VM) *ISC {
	isc := &ISC{
		Ctx:     msg,
		Storage: Storage{storage: storage},
	}
	isc.GVM = &libgvm.GVM{Machine: isc}
	return isc
}

/*                            isc as a gvm machine                            */

func (isc *ISC) CreateRef(t gvm.RefType, v interface{}) gvm.Ref {
	return gvm_type.CreateRef(t, v)
}

func (isc *ISC) GetExternalStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte) (gvm.Ref, error) {
	return isc.Ctx.GetExternalStorageAt(chainID, typeID, contractAddress, pos, description)
}

func (isc *ISC) DecodeRef(t gvm.RefType, b []byte) (gvm.Ref, error) {
	return gvm_type.DecodeRef(t, b)
}

func (isc *ISC) GetFunction(function string) (gvm.Function, error) {
	if function == "main" {
		return &NSBFunctionImpl{storage: isc.Storage.Instructions()}, nil
	} else {
		panic("todo")
	}
}

func (isc *ISC) Commit() error {
	return isc.Storage.storage.Commit()
}

func (isc *ISC) checkRunException(err error) (uint64, error) {
	if err == libgvm.StopUnderFlow {
		err = nil
		return isc.Storage.Instructions().Length(), nil
	}
	if err == trap.ClaimRequest {
		err = nil
	}
	if err != nil {
		return 0, err
	}
	// todo
	return libgvm.GetPC(isc, sugar.HandlerError(libgvm.GetCurrentDepth(isc)).(uint64))
}

func (isc *ISC) initPC() (uint64, error) {
	return isc.checkRunException(isc.GVM.Run("main"))
}

func (isc *ISC) nextPC() (uint64, error) {
	return isc.checkRunException(isc.GVM.Continue())
}

/*                            isc as a contract                            */

func (isc *ISC) IsOpening() bool {
	return isc.Storage.GetISCState() == StateOpening
}

func (isc *ISC) IsActive() bool {
	return isc.Storage.GetISCState() != StateClosed
}

func (isc *ISC) IsInitializing() bool {
	return isc.Storage.GetISCState() == StateInitializing
}

func (isc *ISC) IsInitialized() bool {
	return isc.Storage.GetISCState() == StateInitialized
}

func (isc *ISC) IsSettling() bool {
	return isc.Storage.GetISCState() == StateSettling
}

func (isc *ISC) maybeSwitchToStateSettling(pc uint64) Response {
	if pc >= isc.Storage.Instructions().Length() {
		isc.Storage.setISCState(StateSettling)
	} else if pc < 0 {
		return reportCode(CodePCUnderflow)
	}
	return nil
}

func (isc *ISC) resetAckState() {
	isc.Storage.setISCState(StateInitializing)
	isc.Storage.setFrozenInfoCount(0)
	isc.Storage.setUserAckCount(0)

	var AidMap = isc.Storage.AidMap()
	for idx, l := uint64(0), isc.Storage.Instructions().Length(); idx < l; idx++ {
		AidMap.Set(idx, TxState.Initing)
	}

	var owners = isc.Storage.Owners()
	acknowledged := isc.Storage.UserAcknowledged()
	for idx := uint64(0); idx < owners.Length(); idx++ {
		acknowledged.Delete(owners.Get(idx))
	}
}

type NewContractReply struct {
	Address []byte `json:"address"`
}

func (isc *ISC) NewContract(iscOwners [][]byte,
	funds []uint64,
	instructions []uip.Instruction,
	rawInstructions [][]byte) Response {
	assertTrue(len(iscOwners) != 0, CodeEmptyOwners)
	assertTrue(bytes.Equal(iscOwners[0], isc.Ctx.Sender()), CodeFirstOwnerNotBeSender)
	assertTrue(len(iscOwners) == len(funds), CodeNotEqualLengthOfOwnersAndFunds)
	owners := isc.Storage.Owners()
	mustFunds := isc.Storage.MustFunds()
	isOwner := isc.Storage.IsOwner()
	var AidMap = isc.Storage.AidMap()
	for idx, iscOwner := range iscOwners {
		assertTrueH(!isOwner.Get(iscOwner), CodeDuplicateOwner, hex.EncodeToString(iscOwner))
		owners.Append(iscOwner)
		mustFunds.Set(iscOwner, funds[idx])
		isOwner.Set(iscOwner, true)
	}
	//err = isc.Storage.UserAcknowledged().Set(0, vesSig)
	//if err != nil {
	//	return report(CodeSetError, err)
	//}
	//isc.Storage.SetUint64("userAckCount", 1)

	is := isc.Storage.Instructions()
	for idx, inst := range instructions {
		_ = inst
		is.Append(rawInstructions[idx])
		AidMap.Set(uint64(idx), TxState.Initing)
	}
	return reply().Param(NewContractReply{Address: isc.Ctx.Address()})
}

func (isc *ISC) FreezeInfo(tid uint64) Response {
	assertTrue(isc.IsInitializing(), CodeIsNotInitializing)
	var AidMap = isc.Storage.AidMap()
	if AidMap.Get(tid) == TxState.Initing {
		AidMap.Set(tid, TxState.Inited)
		fc := isc.Storage.getFrozenInfoCount() + 1
		if fc == isc.Storage.Instructions().Length() {
			isc.Storage.setISCState(StateInitialized)
		}
		isc.Storage.setFrozenInfoCount(fc)
	}
	return OK
}

func (isc *ISC) UserAck(fr, signature []byte) Response {
	assertTrue(isc.IsInitialized(), CodeIsNotInitialized)
	acknowledged := isc.Storage.UserAcknowledged()
	if acknowledged.Get(fr) == nil {
		acknowledged.Set(fr, signature)
		uac := isc.Storage.getUserAckCount() + 1
		isc.Storage.setUserAckCount(uac)
		if uac == isc.Storage.Owners().Length() {
			pc, err := isc.initPC()
			if err != nil {
				return report(CodeIteratePCError, err)
			}
			isc.Storage.SetPC(pc)
			isc.Storage.SetMuPC(TxState.Inited)
			isc.Storage.setISCState(StateOpening)
			if r := isc.maybeSwitchToStateSettling(pc); r != nil {
				return r
			}
		}
	}

	return OK
}

//noinspection GoUnusedParameter
func (isc *ISC) UserRefuse(signature []byte) Response {
	assertTrue(isc.IsInitialized(), CodeIsNotInitialized)
	isc.resetAckState()
	return OK
}

func (isc *ISC) InsuranceClaim(tid, aid uint64) Response {
	assertTrue(isc.IsOpening(), CodeIsNotOpening)
	var pc = isc.Storage.GetPC()
	assertTrue(pc == tid, CodeTransactionNotActive)
	var muPC = isc.Storage.GetMuPC() + 1
	assertTrue(muPC == aid, CodeActionNotActive)
	isc.Storage.SetMuPC(muPC)
	if muPC == TxState.Closed {
		var err error
		pc, err = isc.nextPC()
		if err != nil {
			return report(CodeIteratePCError, err)
		}
		isc.Storage.SetPC(pc)
		isc.Storage.SetMuPC(TxState.Inited)
		if r := isc.maybeSwitchToStateSettling(pc); r != nil {
			return r
		}
	}
	return OK
}

func (isc *ISC) SettleContract() Response {
	assertTrue(isc.IsSettling(), CodeIsNotSettling)
	isc.Storage.setISCState(StateClosed)
	return OK
}

func (isc *ISC) GetPC() uint64 {
	return isc.Storage.GetPC()
}

func (isc *ISC) SetPC(u uint64) {
	isc.Storage.SetPC(u)
}

func (isc *ISC) GetMuPC() uint64 {
	return isc.Storage.GetMuPC()
}

func (isc *ISC) SetMuPC(u uint64) {
	isc.Storage.SetMuPC(u)
}

type NSBFunctionImpl struct {
	storage *storage.BytesArray
}

type NSBInstructionImpl struct {
	storage []byte
}

func (N *NSBInstructionImpl) Exec(g *gvm.ExecCtx) error {
	i, err := instruction.DecodeInstruction(bytes.NewReader(N.storage))
	if err != nil {
		return err
	}
	return i.Exec(g)
}

func (N *NSBFunctionImpl) Fetch(pc uint64) (gvm.Instruction, error) {
	b := N.storage.Get(pc)
	return &NSBInstructionImpl{b}, nil
}

func (N *NSBFunctionImpl) Len() int {
	return int(N.storage.Length())
}
