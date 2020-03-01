package isc

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
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

func (isc *ISC) UserAck(fr, signature []byte) Response {
	assertTrue(isc.IsInitialized(), CodeIsNotInitialized)
	acknowledged := isc.Storage.UserAcknowledged()
	if acknowledged.Get(fr) == nil {
		acknowledged.Set(fr, signature)
		uac := isc.Storage.getUserAckCount() + 1
		if uac == isc.Storage.Owners().Length() {
			pc, err := isc.initPC(0)
			if err != nil {
				return report(CodeIteratePCError, err)
			}
			isc.Storage.setPC(pc)
			isc.Storage.setISCState(StateOpening)
		}
		isc.Storage.setUserAckCount(uac)
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
		if pc >= isc.Storage.Instructions().Length() {
			isc.Storage.setISCState(StateSettling)
		} else if pc < 0 {
			return reportCode(CodePCUnderflow)
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
	instruction, err := opintent.HandleInstruction(isc.Storage.Instructions().Get(pc))
	if err != nil {
		return 0, err
	}
	switch instruction.GetType() {
	case instruction_type.Payment, instruction_type.ContractInvoke:
		return 0, nil
	default:
		return pc, nil
	}
}

func (isc *ISC) nextPC(pc uint64) (uint64, error) {
	instruction, err := opintent.HandleInstruction(isc.Storage.Instructions().Get(pc))
	if err != nil {
		return 0, err
	}
	return isc._nextPC(pc, instruction)
}

func (isc *ISC) _nextPC(pc uint64, instruction opintent.LazyInstruction) (uint64, error) {
	switch instruction.GetType() {
	case instruction_type.Goto:
		i, err := instruction.Deserialize()
		if err != nil {
			return 0, err
		}
		//todo
		return uint64(i.(*parser.Goto).Index), nil
	case instruction_type.ConditionGoto:
		i, err := instruction.Deserialize()
		if err != nil {
			return 0, err
		}
		v, err := isc.evalBytes(i.(*parser.ConditionGoto).Condition)
		if err != nil {
			return 0, err
		}
		if v.GetParamType() != value_type.Bool {
			return 0, errors.New("not bool value")
		}
		if v.GetConstant().(bool) {
			return uint64(i.(*parser.ConditionGoto).Index), nil
		}
		return pc+1, nil
	case instruction_type.ConditionSetState:
		i, err := instruction.Deserialize()
		if err != nil {
			return 0, err
		}
		inst := i.(*parser.ConditionSetState)

		v, err := isc.evalBytes(inst.Condition)
		if err != nil {
			return 0, err
		}
		if v.GetParamType() != value_type.Bool {
			return 0, errors.New("not bool value")
		}

		if v.GetConstant().(bool) {
			k, err := isc.evalBytes(inst.RightExpression)
			if err != nil {
				return 0, err
			}
			x, err := isc.save(k)
			if err != nil {
				return 0, err
			}
			isc.Storage.storage.SetBytes(string(inst.Target), x)
		}

		return pc+1, nil
	case instruction_type.SetState:
		i, err := instruction.Deserialize()
		if err != nil {
			return 0, err
		}
		inst := i.(*parser.SetState)

		k, err := isc.evalBytes(inst.RightExpression)
		if err != nil {
			return 0, err
		}
		x, err := isc.save(k)
		if err != nil {
			return 0, err
		}
		isc.Storage.storage.SetBytes(string(inst.Target), x)
		return pc+1, nil
	default:
		return pc+1, nil
	}
}

func (isc *ISC) evalBytes(variable []byte) (token.ConstantI, error) {
	v, err := lexer.ParamUnmarshalJSON(variable)
	if err != nil {
		return nil, err
	}
	return isc.eval(v)
}

func (isc *ISC) eval(v token.Param) (token.ConstantI, error) {
	switch v.GetType(){
	case token.Constant:
		return v.(token.ConstantI), nil
	case token.LocalStateVariable:
		v := v.(token.LocalStateVariableI)
		return isc.load(v.GetField(), v.GetParamType())
	case token.BinaryExpression:
		v := v.(token.BinaryExpressionI)
		l, err := isc.eval(v.GetLeft())
		if err != nil {
			return nil, err
		}
		r, err := isc.eval(v.GetRight())
		if err != nil {
			return nil, err
		}
		switch v.GetSign() {
		case sign_type.EQ:
			return eq(l, r)
		case sign_type.LE:
			return le(l, r)
		case sign_type.LT:
			return lt(l, r)
		case sign_type.GE:
			return ge(l, r)
		case sign_type.GT:
			return gt(l, r)
		case sign_type.LAnd:
			return land(l, r)
		case sign_type.LOr:
			return lor(l, r)
		case sign_type.ADD:
			return add(l, r)
		case sign_type.SUB:
			return sub(l, r)
		case sign_type.MUL:
			return mul(l, r)
		case sign_type.QUO:
			return quo(l, r)
		case sign_type.REM:
			return rem(l, r)
		default:
			return nil, fmt.Errorf("unknown sign_type: %v", v.GetSign())
		}
	case token.UnaryExpression:
		v := v.(token.BinaryExpressionI)
		l, err := isc.eval(v.GetLeft())
		if err != nil {
			return nil, err
		}
		switch v.GetSign() {
		case sign_type.LNot:
			return lnot(l)
		}
	case token.StateVariable:
		//v := v.(token.StateVariableI)
		return nil, errors.New("todo")
	}
	return nil, errors.New("param type not found")
}

func (isc *ISC) save(variable token.ConstantI) ([]byte, error) {
	return storage.Encode(variable)
}

func (isc *ISC) load(field []byte, t value_type.Type) (token.ConstantI, error) {
	return storage.Decode(isc.Storage.storage.GetBytes(string(field)), t)
}