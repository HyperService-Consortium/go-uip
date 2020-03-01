package isc

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/mock"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		s: []byte{1},
		a: []byte{2},
	}

	type args struct {
		iscOwners       [][]byte
		funds           []uint64
		instructions    []uip.Instruction
		rawInstructions [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Response
		wantCode Code
	}{
		{name: Description[CodeEmptyOwners], fields: fields{mock.NewLocalStorage(), ctx}, args: args{}, want: &ResponseError{
			Code: CodeEmptyOwners,
			Err:  "",
		}, wantCode: CodeEmptyOwners},
		{name: Description[CodeFirstOwnerNotBeSender], fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{append(ctx.s, 1)},
			funds: []uint64{0},
		}, want: reportCode(CodeFirstOwnerNotBeSender), wantCode: CodeFirstOwnerNotBeSender},
		{name: Description[CodeNotEqualLengthOfOwnersAndFunds], fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{ctx.s},
			funds: []uint64{},
		}, want: reportCode(CodeNotEqualLengthOfOwnersAndFunds), wantCode: CodeNotEqualLengthOfOwnersAndFunds},
		{name: Description[CodeDuplicateOwner], fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{ctx.s, ctx.s},
			funds: []uint64{0, 0},
		}, want: reportString(CodeDuplicateOwner, fmt.Sprint(hex.EncodeToString(ctx.s))), wantCode: CodeDuplicateOwner},
		{name: "good pure", fields: fields{mock.NewLocalStorage(), ctx}, args: args{
			iscOwners: [][]byte{ctx.s},
			funds: []uint64{0},
		}, want: reply().Param(NewContractReply{Address: ctx.a})},
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
							e.(*ResponseError).Code, Description[e.(*ResponseError).Code],e.(*ResponseError).Err)

						if got := e;
							!assert.EqualValues(t, tt.want, got) {
							t.Errorf("NewContract() = %v, want %v", got, tt.want)
						}
					default:
						panic(err)
					}
				}
			}()

			if got := isc.NewContract(tt.args.iscOwners, tt.args.funds, tt.args.instructions, tt.args.rawInstructions);
				!assert.EqualValues(t, tt.want, got) {
				t.Errorf("NewContract() = %v, want %v", got, tt.want)
			}
		})
	}
}