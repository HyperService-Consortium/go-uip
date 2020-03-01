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
	"math/big"
	"reflect"
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
			pc, err := isc.nextPC(0)
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
		}
	case token.UnaryExpression:
	case token.StateVariable:
		//v := v.(token.StateVariableI)
		return nil, errors.New("todo")
	}
	return nil, errors.New("param type not found")
}

func add(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _add(l, r)
}

func _add(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) + r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) + r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) + r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) + r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) + r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) + r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) + r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) + r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Add(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _sub(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) - r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) - r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) - r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) - r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) - r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) - r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) - r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) - r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Sub(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _mul(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) * r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) * r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) * r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) * r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) * r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) * r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) * r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) * r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Mul(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _div(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) / r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) / r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) / r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) / r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) / r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) / r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) / r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) / r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Div(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _quo(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) % r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) % r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) % r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) % r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) % r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) % r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) % r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) % r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Quo(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func sub(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _sub(l, r)
}

func mul(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _mul(l, r)
}

func div(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _div(l, r)
}

func quo(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _quo(l, r)
}

func rem(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _rem(l, r)
}

func le(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _le(l, r)
}

func lt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _lt(l, r)
}

func _lt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	b, err = _le(l, r)
	if err != nil {
		return nil, err
	}
	if b.GetConstant().(bool) {
		l, err = _eq(l, r)
		if err != nil {
			return nil, err
		}
		if l.GetConstant().(bool) {
			return &lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: false,
			}, nil
		} else {
			return b, nil
		}
	}
	return &lexer.ConstantVariable{
		Type:  value_type.Bool,
		Const: false,
	}, nil
}

func ge(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _ge(l, r)
}

func _ge(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	b, err = _lt(l, r)
	if err != nil {
		return nil, err
	}
	if b.GetConstant().(bool) {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: false,
		}, nil
	} else {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: true,
		}, nil
	}
}

func gt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _gt(l, r)
}

func _gt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	b, err = _le(l, r)
	if err != nil {
		return nil, err
	}
	if b.GetConstant().(bool) {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: false,
		}, nil
	} else {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: true,
		}, nil
	}
}

func eq(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _eq(l, r)
}

func _eq(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64,
		value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64, value_type.Bool, value_type.String:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant() == r.GetConstant(),
		}, nil
	case value_type.Bytes:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: bytes.Equal(r.GetConstant().([]byte), r.GetConstant().([]byte)),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(*big.Int).Cmp(r.GetConstant().(*big.Int)) == 0,
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _le(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator le")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint8) <= r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint16) <= r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint32) <= r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint64) <= r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int8) <= r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int16) <= r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int32) <= r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int64) <= r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(*big.Int).Cmp(r.GetConstant().(*big.Int)) <= 0,
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func convErr(l, r token.ConstantI) (err error) {
	return fmt.Errorf("cant conv param type %v and %v", l.GetParamType(), r.GetParamType())
}

func getBasicTraits(l token.ConstantI) (integral bool, unsigned bool, k int) {
	integral, unsigned, k = true, false, 0
	switch l.GetParamType() {
	case value_type.Bool, value_type.String, value_type.Bytes:
		integral = false
	case value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64, value_type.Uint128, value_type.Uint256:
		unsigned = false
		switch l.GetParamType() {
		case value_type.Uint8:
			k = 8
		case value_type.Uint16:
			k = 16
		case value_type.Uint32:
			k = 32
		case value_type.Uint64:
			k = 64
		case value_type.Uint128:
			k = 128
		case value_type.Uint256:
			k = 256
		}
	case value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64, value_type.Int128, value_type.Int256:
		unsigned = true
		switch l.GetParamType() {
		case value_type.Int8:
			k = 8
		case value_type.Int16:
			k = 16
		case value_type.Int32:
			k = 32
		case value_type.Int64:
			k = 64
		case value_type.Int128:
			k = 128
		case value_type.Int256:
			k = 256
		}
	default:
		panic(errors.New("value type not found"))
	}
	return
}

func tryConv(l token.ConstantI, r token.ConstantI) (_ token.ConstantI, _ token.ConstantI, err error) {
	if l.GetParamType() == r.GetParamType() {
		return l, r, nil
	}
	li, lu, lk := getBasicTraits(l)
	ri, ru, rk := getBasicTraits(r)
	if li && ri {
		if lu && !ru {
			if r, err = tryConvUnsigned(r); err != nil {
				return nil, nil, err
			}
			ru = true
		}
		if ru && !lu {
			if l, err = tryConvUnsigned(l); err != nil {
				return nil, nil, err
			}
			lu = true
		}
		if lk != rk {
			if lk < rk {
				return extentBit(l, rk), r, nil
			} else {
				return l, extentBit(r, lk), nil
			}
		}

	} else {
		return nil, nil, convErr(l, r)
	}

	return l, r, nil
}

func extentBit(l token.ConstantI, rk int) token.ConstantI {
	switch l.GetParamType() {
	case value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64:
		k := reflect.ValueOf(l.GetConstant()).Uint()
		switch rk {
		case 128:
			return &lexer.ConstantVariable{
				Type:  value_type.Uint128,
				Const: big.NewInt(0).SetUint64(k),
			}
		case 256:
			return &lexer.ConstantVariable{
				Type:  value_type.Uint256,
				Const: big.NewInt(0).SetUint64(k),
			}
		default:
			switch rk {
			case 8:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint8,
					Const: uint8(k),
				}
			case 16:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint16,
					Const: uint16(k),
				}
			case 32:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint32,
					Const: uint32(k),
				}
			case 64:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint64,
					Const: uint64(k),
				}
			default:
				panic(fmt.Errorf("bit not found: %v", rk))
			}
		}
	case value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64:
		k := reflect.ValueOf(l.GetConstant()).Int()
		switch rk {
		case 128:
			return &lexer.ConstantVariable{
				Type:  value_type.Int128,
				Const: big.NewInt(k),
			}
		case 256:
			return &lexer.ConstantVariable{
				Type:  value_type.Int256,
				Const: big.NewInt(k),
			}
		default:
			switch rk {
			case 8:
				return &lexer.ConstantVariable{
					Type:  value_type.Int8,
					Const: int8(k),
				}
			case 16:
				return &lexer.ConstantVariable{
					Type:  value_type.Int16,
					Const: int16(k),
				}
			case 32:
				return &lexer.ConstantVariable{
					Type:  value_type.Int32,
					Const: int32(k),
				}
			case 64:
				return &lexer.ConstantVariable{
					Type:  value_type.Int64,
					Const: int64(k),
				}
			default:
				panic(fmt.Errorf("bit not found: %v", rk))
			}
		}
	case value_type.Bytes, value_type.Bool, value_type.String:
		panic(fmt.Errorf("wrong extendBit on non integral type"))
	case value_type.Uint128, value_type.Uint256:
		return &lexer.ConstantVariable{
			Type:  value_type.Uint256,
			Const: l.GetConstant(),
		}
	case value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  value_type.Int256,
			Const: l.GetConstant(),
		}
	default:
		panic(errors.New("value type not found"))
	}
}

func tryConvUnsigned(k token.ConstantI) (token.ConstantI, error) {
	switch k.GetParamType() {
	case value_type.Int8:
		if k.GetConstant().(int8) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint8,
			Const:uint8( k.GetConstant().(int8)),
		}, nil
	case value_type.Int16:
		if k.GetConstant().(int16) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint16,
			Const: uint16(k.GetConstant().(int16)),
		}, nil
	case value_type.Int32:
		if k.GetConstant().(int32) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint32,
			Const: uint32(k.GetConstant().(int32)),
		}, nil
	case value_type.Int64:
		if k.GetConstant().(int64) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint64,
			Const: uint64(k.GetConstant().(int64)),
		}, nil
	case value_type.Int128, value_type.Int256:
		if k.GetConstant().(*big.Int).Sign() < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
	}
	return nil, errors.New("value type invalid")
}

func (isc *ISC) save(variable token.ConstantI) ([]byte, error) {
	return storage.Encode(variable)
}

func (isc *ISC) load(field []byte, t value_type.Type) (token.ConstantI, error) {
	return storage.Decode(isc.Storage.storage.GetBytes(string(field)), t)
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