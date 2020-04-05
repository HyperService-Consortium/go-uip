package isc

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
	"math/big"
)

type Context interface {
	Sender() []byte
	Address() []byte
}

type ISC struct {
	GVM     *gvm.GVM
	Storage Storage
	Msg     Context
}

func (isc *ISC) CreateRef(t gvm.RefType, v interface{}) gvm.Ref {
	return gvm_type.CreateRef(t, v)
}

func (isc *ISC) DecodeRef(t gvm.RefType, b []byte) (gvm.Ref, error) {
	return gvm_type.DecodeRef(t, b)
}

func (isc *ISC) GetFunction(function string) (gvm.Function, error) {
	if function == "main" {
		return createSerialFunction(isc.Storage.Instructions())
	} else {
		panic("todo")
	}
}

type NSBFunctionImpl struct {
	storage *storage.BytesArray
}

type NSBInstructionImpl struct {
	storage []byte
}

type SetState struct {
	Target          string   `json:"target"`
	RightExpression uip.VTok `json:"expression"`
}

func (s *SetState) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	write(w, s.Target, err)
	EncodeVTok(w, s.RightExpression, err)
}

func (s *SetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	read(r, &s.Target, err)
	DecodeVTok(r, &s.RightExpression, err)
	*i = s
}

func (s SetState) Exec(g *gvm.ExecCtx) error {
	return execSetState(g, s.Target, s.RightExpression)
}

func (s SetState) GetType() instruction_type.Type {
	return instruction_type.SetState
}

func EncodeInstructions(is []uip.Instruction) (bs [][]byte, err error) {
	bs = make([][]byte, len(is))
	var b = bytes.NewBuffer(nil)
	for i := range is {
		b.Reset()
		err = EncodeInstruction(is[i], b)
		if err != nil {
			return nil, err
		}
		bs[i] = make([]byte, b.Len())
		copy(bs[i], b.Bytes())
	}
	return bs, nil
}

func write(w io.Writer, v interface{}, err *error) {
	if *err != nil {
		return
	}
	switch v := v.(type) {
	case string:
		*err = binary.Write(w, binary.BigEndian, uint64(len(v)))
		if *err != nil {
			return
		}

		*err = binary.Write(w, binary.BigEndian, []byte(v))
		return
	case []byte:
		*err = binary.Write(w, binary.BigEndian, uint64(len(v)))
	}
	if *err != nil {
		return
	}

	*err = binary.Write(w, binary.BigEndian, v)
	return
}

func read(r io.Reader, target interface{}, err *error) {
	if *err != nil {
		return
	}
	switch t := target.(type) {
	case *string:
		var dataLength uint64
		*err = binary.Read(r, binary.BigEndian, &dataLength)
		if *err != nil {
			return
		}
		var b = make([]byte, dataLength)
		*err = binary.Read(r, binary.BigEndian, &b)
		if *err != nil {
			return
		}
		*t = string(b)
		return
	case *[]byte:
		var dataLength uint64
		*err = binary.Read(r, binary.BigEndian, &dataLength)
		if *err != nil {
			return
		}
		*t = make([]byte, dataLength)
		*err = binary.Read(r, binary.BigEndian, t)
		if *err != nil {
			return
		}
		return
	}
	*err = binary.Read(r, binary.BigEndian, target)
	if *err != nil {
		return
	}
}

func EncodeInstruction(i uip.Instruction, w io.Writer) (err error) {
	write(w, i.GetType(), &err)
	if err != nil {
		return err
	}
	i.Marshal(w, &err)
	return err
}

func EncodeVTok(w io.Writer, v uip.VTok, err *error) {
	if *err != nil {
		return
	}
	write(w, v.GetGVMTok(), err)
	if v.GetGVMTok() == gvm_type.TokConstant {
		write(w, v.GetGVMType(), err)
	}
	v.Marshal(w, err)
	return
}

func DecodeVTok(r io.Reader, v *uip.VTok, err *error) {
	if *err != nil {
		return
	}
	var t gvm.TokType
	*err = binary.Read(r, binary.BigEndian, &t)
	DecodeVTokWithType(r, v, t, err)
}

func DecodeVTokWithType(r io.Reader, v *uip.VTok, t gvm.TokType, err *error) {
	if *err != nil {
		return
	}
	switch t {
	case gvm_type.TokConstant:
		var vt gvm.RefType
		read(r, &vt, err)
		if *err != nil {
			return
		}
		readConstant(r, v, vt, err)
	default:
		panic("todo")
	}
	return
}

func readConstant(r io.Reader, v *uip.VTok, t gvm.RefType, err *error) {
	switch t {
	case gvm_type.RefUint8:
		*v = Uint8(0)
	case gvm_type.RefUint16:
		*v = Uint16(0)
	case gvm_type.RefUint32:
		*v = Uint32(0)
	case gvm_type.RefUint64:
		*v = Uint64(0)

	case gvm_type.RefInt8:
		*v = Int8(0)
	case gvm_type.RefInt16:
		*v = Int16(0)
	case gvm_type.RefInt32:
		*v = Int32(0)
	case gvm_type.RefInt64:
		*v = Int64(0)

	case gvm_type.RefUint128:
		*v = (*Uint128)(big.NewInt(0))
	case gvm_type.RefUint256:
		*v = (*Uint256)(big.NewInt(0))
	case gvm_type.RefInt128:
		*v = (*Int128)(big.NewInt(0))
	case gvm_type.RefInt256:
		*v = (*Int256)(big.NewInt(0))

	case gvm_type.RefString:
		*v = String("")
	case gvm_type.RefBytes:
		*v = Bytes(nil)
	case gvm_type.RefBool:
		*v = Bool(false)
	case gvm_type.RefUnknown:
		*v = Undefined
	default:
		panic(fmt.Errorf("unknown reference type: %v", gvm_type.ExplainGVMType(t)))
	}
	(*v).Unmarshal(r, v, err)
}

func DecodeInstruction(r io.Reader) (i uip.Instruction, err error) {
	var t instruction_type.Type
	err = binary.Read(r, binary.BigEndian, &t)
	if err != nil {
		return nil, err
	}
	return DecodeInstructionWithType(r, t)
}

func DecodeInstructionWithType(r io.Reader, t instruction_type.Type) (i uip.Instruction, err error) {
	switch t {
	case instruction_type.SetState:
		i = new(SetState)
	default:
		return nil, errors.New("invalid type")
	}
	i.Unmarshal(r, &i, &err)
	return
}

func (N *NSBInstructionImpl) Exec(g *gvm.ExecCtx) error {
	i, err := DecodeInstruction(bytes.NewReader(N.storage))
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

func createSerialFunction(i *storage.BytesArray) (gvm.Function, error) {
	return &NSBFunctionImpl{storage: i}, nil
}

func (isc *ISC) Load(field string, t gvm.RefType) (gvm.Ref, error) {
	switch t {
	case gvm_type.RefUint8:
		return isc.CreateRef(t, isc.Storage.storage.GetUint8(field)), nil
	case gvm_type.RefUint16:
		return isc.CreateRef(t, uint16(isc.Storage.storage.GetUint64(field))), nil
	case gvm_type.RefUint32:
		return isc.CreateRef(t, uint32(isc.Storage.storage.GetUint64(field))), nil
	case gvm_type.RefUint64:
		return isc.CreateRef(t, isc.Storage.storage.GetUint64(field)), nil
	case gvm_type.RefInt8:
		return isc.CreateRef(t, isc.Storage.storage.GetInt8(field)), nil
	case gvm_type.RefInt16:
		return isc.CreateRef(t, int16(isc.Storage.storage.GetInt64(field))), nil
	case gvm_type.RefInt32:
		return isc.CreateRef(t, int32(isc.Storage.storage.GetInt64(field))), nil
	case gvm_type.RefInt64:
		return isc.CreateRef(t, isc.Storage.storage.GetInt64(field)), nil
	case gvm_type.RefBool:
		return isc.CreateRef(t, isc.Storage.storage.GetBool(field)), nil
	case gvm_type.RefString:
		return isc.CreateRef(t, isc.Storage.storage.GetString(field)), nil
	case gvm_type.RefBytes, gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
		switch t {
		case gvm_type.RefBytes:
			return isc.CreateRef(t, isc.Storage.storage.GetBytes(field)), nil
		case gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
			return isc.CreateRef(t, big.NewInt(0).SetBytes(isc.Storage.storage.GetBytes(field))), nil
		default:
			panic("runtime case error")
		}
	default:
		return nil, errors.New("type error")
	}
}

func (isc *ISC) Save(field string, k gvm.Ref) error {
	switch k.GetGVMType() {
	case gvm_type.RefUint8:
		isc.Storage.storage.SetUint8(field, k.Unwrap().(uint8))
		return nil
	case gvm_type.RefUint16:
		isc.Storage.storage.SetUint64(field, uint64(k.Unwrap().(uint16)))
		return nil
	case gvm_type.RefUint32:
		isc.Storage.storage.SetUint64(field, uint64(k.Unwrap().(uint32)))
		return nil
	case gvm_type.RefUint64:
		isc.Storage.storage.SetUint64(field, k.Unwrap().(uint64))
		return nil
	case gvm_type.RefInt8:
		isc.Storage.storage.SetInt8(field, k.Unwrap().(int8))
		return nil
	case gvm_type.RefInt16:
		isc.Storage.storage.SetInt64(field, int64(k.Unwrap().(int16)))
		return nil
	case gvm_type.RefInt32:
		isc.Storage.storage.SetInt64(field, int64(k.Unwrap().(int32)))
		return nil
	case gvm_type.RefInt64:
		isc.Storage.storage.SetInt64(field, k.Unwrap().(int64))
		return nil
	case gvm_type.RefBool:
		isc.Storage.storage.SetBool(field, k.Unwrap().(bool))
		return nil
	case gvm_type.RefString:
		isc.Storage.storage.SetString(field, k.Unwrap().(string))
		return nil
	case gvm_type.RefBytes, gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
		var v = k.Unwrap()
		if v == nil && k.GetGVMType() == gvm_type.RefBytes {
			isc.Storage.storage.SetBytes(field, []byte{})
			return nil
		}
		switch k.GetGVMType() {
		case gvm_type.RefBytes:
			isc.Storage.storage.SetBytes(field, k.Unwrap().([]byte))
			return nil
		case gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
			b, err := k.Encode()
			if err != nil {
				return err
			}
			isc.Storage.storage.SetBytes(field, b)
			return nil
		default:
			panic("runtime case error")
		}
	default:
		return errors.New("type error")
	}
}

func (isc *ISC) Delete(_ string) error {
	return nil
}

func NewISC(msg Context, storage *storage.VM) *ISC {
	isc := &ISC{
		Msg:     msg,
		Storage: Storage{storage: storage},
	}
	isc.GVM = &libgvm.GVM{Machine: isc}
	return isc
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
			pc, err := isc.initPC()
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
		pc, err = isc.nextPC()
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

func (isc *ISC) initPC() (uint64, error) {
	err := isc.GVM.Run("main")
	if err == libgvm.StopUnderFlow {
		err = nil
		return isc.Storage.Instructions().Length(), nil
	}
	if err != nil {
		return 0, err
	}
	return libgvm.GetCurrentPC(isc)
}

func (isc *ISC) nextPC() (uint64, error) {
	err := isc.GVM.Continue()
	if err == libgvm.StopUnderFlow {
		err = nil
		return isc.Storage.Instructions().Length(), nil
	}
	if err != nil {
		return 0, err
	}
	return libgvm.GetCurrentPC(isc)
}

//func checkValueType(t value_type.Type) error {
//
//	if t >= value_type.Length || t <= value_type.Unknown {
//		return fmt.Errorf("unknown value_type: %v", t)
//	}
//	return nil
//}
//
//func checkEvaluableTokenType(t token.Type) error {
//	if !token.IsEvaluable(t) {
//		return fmt.Errorf("token type %v is not evaluable", t)
//	}
//	return nil
//}
