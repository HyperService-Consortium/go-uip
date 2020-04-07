package internal_test

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/instruction/internal"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestSerialize(t *testing.T) {
	type testCase struct {
		name      string
		marshal   func(r io.Writer, b interface{})
		unmarshal func(r io.Reader) interface{}
		encoding  uip.Serializable
		wantErr   error
	}

	var inheritInstructionTestCase = func(name string, encoding uip.Serializable, wantErr ...error) testCase {
		if len(wantErr) > 1 {
			panic("bad want param")
		}
		var we error
		if len(wantErr) == 1 {
			we = wantErr[0]
		}

		return testCase{
			name:      name,
			encoding:  encoding,
			marshal:   encodeInstruction,
			unmarshal: unmarshalInstruction,
			wantErr:   we,
		}
	}

	for _, tt := range []testCase{
		inheritInstructionTestCase("encode-goto", internal.Goto{Index: 1}),
		inheritInstructionTestCase("encode-condition-goto", internal.ConditionGoto{
			Condition: lexer.Bool(true), Index: 1}),
		inheritInstructionTestCase("encode-set-state", internal.SetState{
			Type: 1, Target: "target", RightExpression: lexer.Bool(true)}),
		inheritInstructionTestCase("encode-condition-set-state", internal.ConditionSetState{
			Type: 1, Target: "target",
			RightExpression: lexer.Bool(true), Condition: lexer.Bool(true)}),
	} {
		t.Run(tt.name, func(t *testing.T) {
			var b = bytes.NewBuffer(nil)
			if tt.wantErr != nil {
				defer func() {
					assert.EqualValues(t, tt.wantErr, recover())
				}()
			}
			tt.marshal(b, tt.encoding)
			assert.EqualValues(t, tt.encoding, tt.unmarshal(b))
			assert.EqualValues(t, 0, b.Len())
		})
	}

}

func encodeInstruction(r io.Writer, b interface{}) {
	sugar.HandlerError0(instruction.EncodeInstruction(b.(uip.Instruction), r))
	return
}

func unmarshalInstruction(r io.Reader) interface{} {
	return sugar.HandlerError(instruction.DecodeInstruction(r))
}
