package isc

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/mock"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/instruction"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var user0 = []byte{1}

func funcSetA() []uip.Instruction {
	return []uip.Instruction{
		&instruction.SetState{
			Target:          "a",
			RightExpression: lexer.Bool(true),
		},
	}
}

func encodeInstructions(is []uip.Instruction) (bs [][]byte) {
	return sugar.HandlerError(instruction.EncodeInstructions(is)).([][]byte)
}

type ContextImpl struct {
	s, a []byte
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

	ctx := ContextImpl{
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

	ctx := ContextImpl{
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
			assert.EqualValues(t, StateInitialized, isc.Storage.getISCState())
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isc := &ISC{
				Storage: tt.fields.Storage,
				Msg:     tt.fields.Msg,
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
		Msg: &ContextImpl{
			s: user0,
			a: []byte{2},
		}}

	assert.EqualValues(t, OK, _isc.FreezeInfo(0))
	sugar.HandlerError0(_isc.Storage.storage.Commit())
	assert.EqualValues(t, StateInitialized, _isc.Storage.getISCState())

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
			assert.EqualValues(t, StateSettling, isc.Storage.getISCState())
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isc := &ISC{
				Storage: tt.fields.Storage,
				Msg:     tt.fields.Msg,
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
				Msg:     tt.fields.Msg,
			}
			isc.GVM = gvm.Wrap(isc)

			if got := isc.InsuranceClaim(tt.args.tid, tt.args.aid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsuranceClaim() = %v, want %v", got, tt.want)
			}
			sugar.HandlerError0(isc.Storage.storage.Commit())
		})
	}
}
