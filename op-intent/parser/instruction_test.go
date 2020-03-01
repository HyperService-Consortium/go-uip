package parser

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/standard"
	"github.com/HyperService-Consortium/go-uip/uip"
	"testing"
)

func TestInstructions(t *testing.T) {
	for _, tt := range []struct {
		name string
		impl uip.Instruction
	} {
		//{"raw goto", &RawGoto{}},
		//{"raw condition goto", &RawConditionGoto{}},
		//{"raw set state", &RawSetState{}},
		{"goto", NewGoto(0)},
		{"condition goto", NewConditionGoto(0, nil)},
		{"set state", NewSetState(0, nil, nil)},
		{"condition set state", NewConditionSetState(0, nil, nil, nil)},
		{"transaction intent", &TransactionIntent{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := standard.IsValidInstructionImpl(tt.impl); err != nil {
				t.Error(err)
			}
			fmt.Print(tt.impl)
		})
	}
}

