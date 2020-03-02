package gvm

import "testing"

type gotoImpl struct{}

func (g gotoImpl) GetGVMIType() InstructionType {
	return 1
}

func (g gotoImpl) GetGotoIndexGVMI() uint64 {
	return 2
}

func Test__Goto_Exec(t *testing.T) {
	type args struct {
		g *ExecCtx
		i Instruction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"easy", args{&ExecCtx{
			GVM:      nil,
			Function: nil,
			PC:       1,
		}, gotoImpl{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_g := &_Goto{}
			if err := _g.Exec(tt.args.g, tt.args.i); (err != nil) != tt.wantErr {
				t.Fatalf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				return
			}

			if uint64(tt.args.i.(Goto).GetGotoIndexGVMI()) != tt.args.g.PC {
				t.Errorf("Exec() get pc = %v, want: %v", tt.args.g.PC, tt.args.i.(Goto).GetGotoIndexGVMI())
			}
		})
	}
}