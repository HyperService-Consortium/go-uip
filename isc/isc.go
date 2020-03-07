package isc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Context interface {
	Sender() []byte
	Address() []byte
}

type ISC struct {
	Storage Storage
	Msg     Context
}

func NewISC(msg Context, storage *storage.VM) *ISC {
	return &ISC{
		Msg:     msg,
		Storage: Storage{storage: storage},
	}
}

func (isc *ISC) IsOpening() bool {
	return isc.Storage.getISCState() == StateOpening
}

func (isc *ISC) IsActive() bool {
	return isc.Storage.getISCState() != StateClosed
}

func (isc *ISC) IsInitializing() bool {
	return isc.Storage.getISCState() == StateInitializing
}

func (isc *ISC) IsInitialized() bool {
	return isc.Storage.getISCState() == StateInitialized
}

func (isc *ISC) IsSettling() bool {
	return isc.Storage.getISCState() == StateSettling
}

type NewContractReply struct {
	Address []byte `json:"address"`
}

func (isc *ISC) NewContract(iscOwners [][]byte,
	funds []uint64,
	instructions []uip.Instruction,
	rawInstructions [][]byte) Response {
	assertTrue(len(iscOwners) != 0, CodeEmptyOwners)
	assertTrue(bytes.Equal(iscOwners[0], isc.Msg.Sender()), CodeFirstOwnerNotBeSender)
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
	return reply().Param(NewContractReply{Address: isc.Msg.Address()})
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

func (isc *ISC) maybeSwitchToStateSettling(pc uint64) Response {
	if pc >= isc.Storage.Instructions().Length() {
		isc.Storage.setISCState(StateSettling)
	} else if pc < 0 {
		return reportCode(CodePCUnderflow)
	}
	return nil
}

func (isc *ISC) UserAck(fr, signature []byte) Response {
	assertTrue(isc.IsInitialized(), CodeIsNotInitialized)
	acknowledged := isc.Storage.UserAcknowledged()
	if acknowledged.Get(fr) == nil {
		acknowledged.Set(fr, signature)
		uac := isc.Storage.getUserAckCount() + 1
		isc.Storage.setUserAckCount(uac)
		if uac == isc.Storage.Owners().Length() {
			pc, err := isc.initPC(0)
			if err != nil {
				return report(CodeIteratePCError, err)
			}
			isc.Storage.setPC(pc)
			isc.Storage.setISCState(StateOpening)
			if r := isc.maybeSwitchToStateSettling(pc); r != nil {
				return r
			}
		}
	}

	return OK
}

func (isc *ISC) UserRefuse(signature []byte) Response {
	assertTrue(isc.IsInitialized(), CodeIsNotInitialized)
	isc.resetAckState()
	return OK
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

func (isc *ISC) InsuranceClaim(tid, aid uint64) Response {
	assertTrue(isc.IsInitialized(), CodeIsNotOpening)
	var pc = isc.Storage.getPC()
	assertTrue(pc == tid, CodeTransactionNotActive)
	var AidMap = isc.Storage.AidMap()
	var miuPC = AidMap.Get(tid) + 1
	assertTrue(miuPC == aid, CodeActionNotActive)
	AidMap.Set(pc, miuPC)
	if miuPC == TxState.Closed {
		var err error
		pc, err = isc.nextPC(pc)
		if err != nil {
			return report(CodeIteratePCError, err)
		}
		isc.Storage.setPC(pc)
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

func (isc *ISC) initPC(pc uint64) (uint64, error) {
	panic("")
}

func (isc *ISC) nextPC(pc uint64) (uint64, error) {
	panic("")
}

func checkValueType(t value_type.Type) error {

	if t >= value_type.Length || t <= value_type.Unknown {
		return fmt.Errorf("unknown value_type: %v", t)
	}
	return nil
}

func checkEvaluableTokenType(t token.Type) error {
	if !token.IsEvaluable(t) {
		return fmt.Errorf("token type %v is not evaluable", t)
	}
	return nil
}
