package isc

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	merkle_proof "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	error2 "github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/mock"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"math/big"
	"reflect"
	"testing"
)

var user0 = []byte{1}

func funcSetA() []uip.Instruction {
	return []uip.Instruction{
		&instruction.SetState{
			Target:          "a",
			RightExpression: lexer_types.Bool(true),
		},
	}
}

func encodeInstructions(is []uip.Instruction) (bs [][]byte) {
	return sugar.HandlerError(instruction.EncodeInstructions(is)).([][]byte)
}

type StorageKey struct {
	chainID         uip.ChainID
	typeID          uip.TypeID
	contractAddress string
	pos             string
	description     string
}

type ContextImpl struct {
	s, a            []byte
	externalStorage map[StorageKey]gvm.Ref
}

func (c *ContextImpl) GetExternalStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte) (gvm.Ref, error) {
	if c.externalStorage == nil {
		c.externalStorage = make(map[StorageKey]gvm.Ref)
	}
	if x, ok := c.externalStorage[StorageKey{
		chainID:         chainID,
		typeID:          typeID,
		contractAddress: string(contractAddress),
		pos:             string(pos),
		description:     string(description),
	}]; ok {
		return x, nil
	} else {
		return nil, errors.New("no found")
	}
}

func (c *ContextImpl) ProvideExternalStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte, ref gvm.Ref) {
	if c.externalStorage == nil {
		c.externalStorage = make(map[StorageKey]gvm.Ref)
	}
	c.externalStorage[StorageKey{
		chainID:         chainID,
		typeID:          typeID,
		contractAddress: string(contractAddress),
		pos:             string(pos),
		description:     string(description),
	}] = ref
}

func (c ContextImpl) Sender() []byte {
	return c.s
}

func (c ContextImpl) Address() []byte {
	return c.a
}

func TestISC_NewContract(t *testing.T) {
	type fields struct {
		Storage storage.LocalStorage
		msg     Context
	}

	ctx := &ContextImpl{
		s: user0,
		a: []byte{2},
	}

	type args struct {
		iscOwners       [][]byte
		funds           []uint64
		instructions    []uip.Instruction
		rawInstructions [][]byte
	}

	type testCase struct {
		name     string
		fields   fields
		args     args
		want     Response
		wantCode Code
		callback func(t *testing.T, tt testCase)
	}

	tests := []testCase{
		{name: Description[CodeEmptyOwners], fields: fields{mock.NewLocalStorage(), ctx}, args: args{}, want: &ResponseError{
			Code: CodeEmptyOwners,
			Err:  "",
		}, wantCode: CodeEmptyOwners},
		{name: Description[CodeFirstOwnerNotBeSender], fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{append(ctx.s, 1)},
			funds:     []uint64{0},
		}, want: reportCode(CodeFirstOwnerNotBeSender), wantCode: CodeFirstOwnerNotBeSender},
		{name: Description[CodeNotEqualLengthOfOwnersAndFunds], fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{ctx.s},
			funds:     []uint64{},
		}, want: reportCode(CodeNotEqualLengthOfOwnersAndFunds), wantCode: CodeNotEqualLengthOfOwnersAndFunds},
		{name: Description[CodeDuplicateOwner], fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{ctx.s, ctx.s},
			funds:     []uint64{0, 0},
		}, want: reportString(CodeDuplicateOwner, fmt.Sprint(hex.EncodeToString(ctx.s))), wantCode: CodeDuplicateOwner},
		{name: "good pure", fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{ctx.s},
			funds:     []uint64{0},
		}, want: reply().Param(NewContractReply{Address: ctx.a})},
		{name: "good with instruction", fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners:       [][]byte{ctx.s},
			funds:           []uint64{0},
			instructions:    funcSetA(),
			rawInstructions: encodeInstructions(funcSetA()),
		}, want: reply().Param(NewContractReply{Address: ctx.a}), callback: func(t *testing.T, tt testCase) {
			isc := NewISC(tt.fields.msg, storage.NewVM(tt.fields.Storage))

			assert.EqualValues(t, encodeInstructions(funcSetA())[0], isc.Storage.Instructions().Get(0))
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			isc := NewISC(tt.fields.msg, storage.NewVM(tt.fields.Storage))
			defer func() {
				if err := recover(); err != nil {
					switch e := err.(type) {
					case Response:

						if e.GetCode() != tt.wantCode {
							t.Errorf("NewContract().Code = %v, want %v", e.(*ResponseError).Code, tt.wantCode)
						}
						if IsOK(e) {
							t.Error("err is ok", err)
						}
						fmt.Printf(
							`<code:%v, desc:%v, err: %v>
`,
							e.(*ResponseError).Code, Description[e.(*ResponseError).Code], e.(*ResponseError).Err)

						if got := e; !assert.EqualValues(t, tt.want, got) {
							t.Errorf("NewContract() = %v, want %v", got, tt.want)
						}
					default:
						panic(err)
					}
				}
			}()

			if got := isc.NewContract(tt.args.iscOwners, tt.args.funds, tt.args.instructions, tt.args.rawInstructions); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("NewContract() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, tt)
			}
		})
	}
}

func createTestContract() Storage {
	var vm = storage.NewVM(mock.NewLocalStorage())

	ctx := &ContextImpl{
		s: user0,
		a: []byte{2},
	}
	isc := NewISC(ctx, vm)
	var response = isc.NewContract([][]byte{ctx.s}, []uint64{0}, funcSetA(), encodeInstructions(funcSetA()))
	if response.GetCode() != CodeOK {
		panic(response)
	}
	sugar.HandlerError0(isc.Storage.storage.Commit())
	return Storage{vm}
}

func TestISC_FreezeInfo(t *testing.T) {

	var c1 = createTestContract()

	type fields struct {
		Storage Storage
		Msg     Context
	}
	type args struct {
		tid uint64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     Response
		callback func(t *testing.T, isc *ISC)
	}{
		{name: "freeze-01", fields: fields{
			Storage: c1,
			Msg: &ContextImpl{
				s: user0,
				a: []byte{2},
			},
		}, args: args{
			0,
		}, want: OK, callback: func(t *testing.T, isc *ISC) {
			assert.EqualValues(t, StateInitialized, isc.Storage.GetISCState())
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isc := &ISC{
				Storage: tt.fields.Storage,
				Ctx:     tt.fields.Msg,
			}
			isc.GVM = gvm.Wrap(isc)

			if got := isc.FreezeInfo(tt.args.tid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreezeInfo() = %v, want %v", got, tt.want)
			}
			sugar.HandlerError0(isc.Storage.storage.Commit())
			if tt.callback != nil {
				tt.callback(t, isc)
			}
		})
	}
}

func TestISC_UserAck(t *testing.T) {

	var c1 = createTestContract()
	var _isc = &ISC{Storage: c1,
		Ctx: &ContextImpl{
			s: user0,
			a: []byte{2},
		}}

	assert.EqualValues(t, OK, _isc.FreezeInfo(0))
	sugar.HandlerError0(_isc.Storage.storage.Commit())
	assert.EqualValues(t, StateInitialized, _isc.Storage.GetISCState())

	type fields struct {
		Storage Storage
		Msg     Context
	}
	type args struct {
		fr        []byte
		signature []byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     Response
		callback func(t *testing.T, isc *ISC)
	}{
		{name: "freeze-01", fields: fields{
			Storage: c1,
			Msg: &ContextImpl{
				s: user0,
				a: []byte{2},
			},
		}, args: args{
			user0, []byte{1},
		}, want: OK, callback: func(t *testing.T, isc *ISC) {
			assert.EqualValues(t, StateSettling, isc.Storage.GetISCState())
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isc := &ISC{
				Storage: tt.fields.Storage,
				Ctx:     tt.fields.Msg,
			}
			isc.GVM = gvm.Wrap(isc)

			if got := isc.UserAck(tt.args.fr, tt.args.signature); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserAck() = %v, want %v", got, tt.want)
			}
			sugar.HandlerError0(isc.Storage.storage.Commit())
			if tt.callback != nil {
				tt.callback(t, isc)
			}
		})
	}
}

func TestISC_InsuranceClaim(t *testing.T) {
	type fields struct {
		Storage Storage
		Msg     Context
	}
	type args struct {
		tid uint64
		aid uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isc := &ISC{
				Storage: tt.fields.Storage,
				Ctx:     tt.fields.Msg,
			}
			isc.GVM = gvm.Wrap(isc)

			if got := isc.InsuranceClaim(tt.args.tid, tt.args.aid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsuranceClaim() = %v, want %v", got, tt.want)
			}
			sugar.HandlerError0(isc.Storage.storage.Commit())
		})
	}
}

var c2 = obj{
	"contractName": "c2",
	"domain":       2,
	"address":      "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
}

type obj map[string]interface{}

var opIntents = obj{
	"op-intents": []obj{
		{
			"name": "op1",
			"type": "Payment",
			"src": obj{
				"domain":    1,
				"user_name": "a1",
			},
			"dst": obj{
				"domain":    2,
				"user_name": "a2",
			},
			"amount": "1a",
			"unit":   "ether",
		},
		{
			"name":    "op2",
			"type":    "ContractInvocation",
			"invoker": "a2",
			"func":    "vote",
			"contract": obj{
				"domain":  2,
				"address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
			},
			"parameters": []obj{},
		},
		{
			"name": "if-op",
			"type": "IfStatement",
			"if": []obj{
				{
					"name":    "op3",
					"type":    "ContractInvocation",
					"invoker": "a2",
					"func":    "vote",
					"contract": obj{
						"address": "0x3723261b2a5a62b778b5c74318d34d7fdbadb38e",
					},
					"parameters": []obj{},
				},
				{
					"name": "op4",
					"type": "Payment",
					"src": obj{
						"domain":    1,
						"user_name": "a1",
					},
					"dst": obj{
						"domain":    2,
						"user_name": "a2",
					},
					"amount": "aa",
					"unit":   "ether",
				},
			},
			"else": []obj{
				{
					"name":    "op5",
					"type":    "ContractInvocation",
					"invoker": "a2",
					"func":    "vote",
					"contract": obj{
						"domain":  2,
						"address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
					},
					"parameters": []obj{},
				},
			},
			"condition": obj{
				"left": obj{
					"type": "uint256",
					"value": obj{
						"contract": "c2",
						"field":    "num_count",
						"pos":      "00",
					},
				},
				"right": obj{
					"type": "uint256",
					"value": obj{
						"contract": "c2",
						"field":    "totalVotes",
						"pos":      "01",
					},
				},
				"sign": "Greater",
			},
		},
		{
			"name": "loop",
			"type": "loopFunction",
			"loop": []obj{
				{
					"name":    "op6",
					"type":    "ContractInvocation",
					"invoker": "a2",
					"func":    "vote",
					"contract": obj{
						"domain":  2,
						"address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
					},
					"parameters": []obj{},
				},
			},
			"loopTime": "5",
		},
	},
	"dependencies": []obj{},
	"contracts": []obj{
		{
			"contractName": "c1",
			"domain":       1,
			"address":      "0xafc7d2959e72081770304f6474151293be1fbba7",
		},
		c2,
		{
			"contractName": "c3",
			"domain":       3,
			"address":      "0x3723261b2a5a62b778b5c74318d34d7fdbadb38e",
		},
	},
	"accounts": []obj{
		{
			"userName": "a1",
			"domain":   1,
			"address":  "0x7019fa779024c0a0eac1d8475733eefe10a49f3b",
		},
		{
			"userName": "a2",
			"domain":   2,
			"address":  "0x47a1cdb6594d6efed3a6b917f2fbaa2bbcf61a2e",
		},
		{
			"userName": "a3",
			"domain":   3,
			"address":  "0x47a1cdb6559d6efed3a6b917f2fbaa2bbcf61a2e",
		},
	},
}

func setupOpIntent(t *testing.T) (ctx *ContextImpl, isc *ISC) {
	t.Helper()
	var intents parser.TxIntents

	ier, err := opintent.NewInitializer(uip.BlockChainGetterNilImpl{}, mAccountProvider{})
	if err != nil {
		t.Error(err)
		return
	}

	p := packet{
		content: sugar.HandlerError(json.Marshal(opIntents)).([]byte),
	}

	intents, err = ier.ParseR(p)
	if err != nil {
		t.Error(err)
		pe := err.(*error2.ParseError)
		fmt.Println(string(sugar.HandlerError(pe.Serialize()).([]byte)))
		return
	}
	var txIntents = intents.GetTxIntents()
	var instructions []uip.Instruction
	for i := range txIntents {
		fmt.Println(i, txIntents[i].GetName(), txIntents[i].GetInstruction().GetType())
		instructions = append(instructions, txIntents[i].GetInstruction())
	}

	ctx = &ContextImpl{
		s: user0,
		a: []byte{2},
	}
	isc = NewISC(ctx, storage.NewVM(mock.NewLocalStorage()))

	var newContractReply NewContractReply

	unpack(isc.NewContract([][]byte{ctx.s}, []uint64{0}, instructions, encodeInstructions(instructions)), &newContractReply)
	commit(t, isc)
	fmt.Println(newContractReply)
	assert.EqualValues(t, StateInitializing, isc.Storage.GetISCState())

	for i := range instructions {
		assert.EqualValues(t, OK, isc.FreezeInfo(uint64(i)))
		commit(t, isc)
	}
	assert.EqualValues(t, StateInitialized, isc.Storage.GetISCState())

	assert.EqualValues(t, OK, isc.UserAck(user0, []byte("todo")))
	assert.EqualValues(t, StateOpening, isc.Storage.GetISCState())
	return
}

func TestIfScenario_IfYes(t *testing.T) {
	ctx, isc := setupOpIntent(t)
	//0 op1.cna 0
	doTransaction(t, isc, uint64(0))
	//1 op1.cnb 0
	doTransaction(t, isc, uint64(1))
	//2 op2 1

	BranchIfTest0(ctx, true)
	doTransaction(t, isc, uint64(2))
	//3 if-op.goto.if 3
	//4 op5 1
	//5 if-op.goto.else 2
	//6 op3 1
	doTransaction(t, isc, uint64(6))
	//7 op4.cna 0
	doTransaction(t, isc, uint64(7))
	//8 op4.cnb 0
	doTransaction(t, isc, uint64(8))
	//9 loop.loopBegin 3
	//10 op6 1
	for i := 0; i < 5; i++ {
		doTransaction(t, isc, uint64(10))
	}
	//11 loop.addLoopVar 4
	//12 loop.loopEnd 2
	//13 loop.resetLoopVar 4

	assert.EqualValues(t, StateSettling, isc.Storage.GetISCState())

	assert.EqualValues(t, OK, isc.SettleContract())
	commit(t, isc)

	assert.EqualValues(t, StateClosed, isc.Storage.GetISCState())
}

func TestIfScenario_IfNo(t *testing.T) {
	ctx, isc := setupOpIntent(t)
	//0 op1.cna 0
	doTransaction(t, isc, uint64(0))
	//1 op1.cnb 0
	doTransaction(t, isc, uint64(1))
	//2 op2 1

	BranchIfTest0(ctx, false)
	doTransaction(t, isc, uint64(2))
	//3 if-op.goto.if 3
	//4 op5 1
	doTransaction(t, isc, uint64(4))
	//5 if-op.goto.else 2
	//6 op3 1
	//7 op4.cna 0
	//8 op4.cnb 0
	//9 loop.loopBegin 3
	//10 op6 1
	for i := 0; i < 5; i++ {
		doTransaction(t, isc, uint64(10))
	}
	//11 loop.addLoopVar 4
	//12 loop.loopEnd 2
	//13 loop.resetLoopVar 4

	assert.EqualValues(t, StateSettling, isc.Storage.GetISCState())

	assert.EqualValues(t, OK, isc.SettleContract())
	commit(t, isc)

	assert.EqualValues(t, StateClosed, isc.Storage.GetISCState())
}

func doTransaction(t *testing.T, isc *ISC, pc uint64) {
	t.Helper()
	assert.EqualValues(t, pc, isc.GetPC())
	assert.EqualValues(t, OK, isc.InsuranceClaim(pc, TxState.Instantiating))
	commit(t, isc)

	assert.EqualValues(t, pc, isc.GetPC())
	assert.EqualValues(t, OK, isc.InsuranceClaim(pc, TxState.Open))
	commit(t, isc)

	assert.EqualValues(t, pc, isc.GetPC())
	assert.EqualValues(t, OK, isc.InsuranceClaim(pc, TxState.Opened))
	commit(t, isc)

	assert.EqualValues(t, pc, isc.GetPC())
	assert.EqualValues(t, OK, isc.InsuranceClaim(pc, TxState.Closed))
	commit(t, isc)
}

func BranchIfTest0(ctx *ContextImpl, ifOrNot bool) {
	//"left": obj{
	//	"type": "uint256",
	//	"value": obj{
	//		"contract": "c2",
	//		"field":    "num_count",
	//		"pos":      "00",
	//	},
	//},
	var l0 *lexer_types.Uint256

	if ifOrNot {
		l0 = (*lexer_types.Uint256)(big.NewInt(2))
	} else {
		l0 = (*lexer_types.Uint256)(big.NewInt(1))
	}
	ctx.ProvideExternalStorageAt(
		uip.ChainIDUnderlyingType(c2["domain"].(int)), value_type.Uint256,
		sugar.HandlerError(hex.DecodeString(c2["address"].(string)[2:])).([]byte), []byte{0}, []byte("num_count"),
		l0)
	//	"right": obj{
	//	"type": "uint256",
	//	"value": obj{
	//		"contract": "c2",
	//		"field":    "totalVotes",
	//		"pos":      "01",
	//	},
	//},
	ctx.ProvideExternalStorageAt(
		uip.ChainIDUnderlyingType(c2["domain"].(int)), value_type.Uint256,
		sugar.HandlerError(hex.DecodeString(c2["address"].(string)[2:])).([]byte), []byte{1}, []byte("totalVotes"),
		(*lexer_types.Uint256)(big.NewInt(1)))
	//	"sign": "Greater",
}

func commit(t *testing.T, isc *ISC) {
	t.Helper()
	sugar.HandlerError0(isc.Storage.storage.Commit())
}

func unpack(response Response, n *NewContractReply) {
	var data = response.(*ResponseData).Data
	//fmt.Println(string(data))
	sugar.HandlerError0(json.Unmarshal(data, n))
}

type packet struct {
	content []byte
}

func (p packet) GetContent() (content []byte) {
	return p.content
}

type mAccountProvider struct {
}

func (a mAccountProvider) AccountBase() uip.AccountBase {
	return a
}

func (mAccountProvider) Get(_ string, chainId uint64) (uip.Account, error) {
	return &uip.AccountImpl{
		ChainId: chainId,
		Address: []byte("121313212313133123333333333333333313"),
	}, nil
}

func (mAccountProvider) GetRelay(domain uint64) (uip.Account, error) {
	return &uip.AccountImpl{
		ChainId: domain,
		Address: []byte("99999"),
	}, nil
}

func (mAccountProvider) GetTransactionProofType(_ uint64) (uip.MerkleProofType, error) {
	return merkle_proof.MerklePatriciaTrieUsingKeccak256, nil
}
