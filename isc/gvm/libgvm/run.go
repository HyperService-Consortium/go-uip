package libgvm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"strconv"
)

func Step(g abstraction.GVM, fn string, pc uint64) (uint64, error) {
	f, inst, err := fetch(g, fn, pc)
	if err != nil {
		return 0, err
	}
	c := abstraction.ExecCtx{GVM: g, Function: f, PC: pc}
	return c.PC, inst.Exec(&c)
}

func Continue(g abstraction.GVM, fn string, pc uint64) (uint64, error) {

	f, err := g.GetFunction(fn)
	c := abstraction.ExecCtx{GVM: g, Function: f, PC: pc}
	for ; err == nil; err = Iter(&c) {
	}
	return c.PC, err
}

func loadLocal(g abstraction.GVM, d uint64) (map[string]abstraction.Ref, error) {
	r, err := g.Load("_gvm_locals_"+strconv.FormatUint(d, 16), RefBytes)
	if err != nil {
		return nil, err
	}
	return deserializeMapSR(g, r.Unwrap().([]byte))
}

func saveLocals(g abstraction.GVM, d uint64, locals map[string]abstraction.Ref) error {
	b, err := serializeMapSR(g, locals)
	if err != nil {
		return err
	}
	err = g.Save("_gvm_locals_"+strconv.FormatUint(d, 16), g.CreateRef(RefBytes, b))
	if err != nil {
		return err
	}
	return nil
}

func serializeMapSR(g abstraction.GVM, mp map[string]abstraction.Ref) (_ []byte, err error) {
	var b = bytes.NewBuffer(make([]byte, 0, 50))
	err = binary.Write(b, binary.BigEndian, int32(len(mp)))
	if err != nil {
		return
	}

	var x []byte

	for k, v := range mp {
		x, err = v.Encode()
		if err != nil {
			return
		}

		err = binary.Write(b, binary.BigEndian, v.GetGVMType())
		if err != nil {
			return
		}
		err = binary.Write(b, binary.BigEndian, int32(len(x)))
		if err != nil {
			return
		}
		err = binary.Write(b, binary.BigEndian, x)
		if err != nil {
			return
		}

		err = binary.Write(b, binary.BigEndian, int32(len(k)))
		if err != nil {
			return
		}
		err = binary.Write(b, binary.BigEndian, []byte(k))
		if err != nil {
			return
		}
	}
	return b.Bytes(), nil
}

func deserializeMapSR(g abstraction.GVM, rb []byte) (mp map[string]abstraction.Ref, err error) {
	var b = bytes.NewBuffer(rb)
	var mpLen, dataLen int32
	err = binary.Read(b, binary.BigEndian, &mpLen)
	if err != nil {
		return
	}

	mp = make(map[string]abstraction.Ref)

	var x []byte
	var pType uint64
	var v []byte
	var r abstraction.Ref

	for i := int32(0); i < mpLen; i++ {
		err = binary.Read(b, binary.BigEndian, &pType)
		if err != nil {
			return
		}

		err = binary.Read(b, binary.BigEndian, &dataLen)
		if err != nil {
			return
		}
		if cap(x) < int(dataLen) {
			x = make([]byte, dataLen, dataLen+50)
		}
		if len(x) < int(dataLen) {
			x = x[:dataLen]
		}

		_, err = b.Read(x)
		if err != nil {
			return
		}

		v = x[:dataLen]

		err = binary.Read(b, binary.BigEndian, &dataLen)
		if err != nil {
			return
		}
		if cap(x) < int(dataLen) {
			x = append(x, make([]byte, int(dataLen)-cap(x))...)
		}
		if len(x) < int(dataLen) {
			x = x[:dataLen]
		}
		_, err = b.Read(x)
		if err != nil {
			return
		}

		r, err = g.DecodeRef(RefBytes, v)
		if err != nil {
			return
		}

		mp[string(x[:dataLen])] = r
	}
	return
}

func pushFrame(g abstraction.GVM, fn string) error {
	d, err := GetCurrentDepth(g)
	if err != nil {
		return err
	}
	err = setCurrentDepth(g, d+1)
	if err != nil {
		return err
	}
	pc, err := GetCurrentPC(g)
	if err != nil {
		return err
	}
	ofn, err := GetCurrentFN(g)
	if err != nil {
		return err
	}
	err = setPC(g, d, pc)
	if err != nil {
		return err
	}
	err = setFN(g, d, ofn)
	if err != nil {
		return err
	}
	err = setCurrentFN(g, fn)
	if err != nil {
		return err
	}
	err = setCurrentPC(g, 0)
	if err != nil {
		return err
	}

	return nil
}

func popFrame(g abstraction.GVM) (uint64, string, error) {
	d, err := GetCurrentDepth(g)
	if err != nil {
		return 0, "", err
	}
	if d == 0 {
		return 0, "", errors.New("depth underflow")
	}
	d--
	err = setCurrentDepth(g, d)
	if err != nil {
		return 0, "", err
	}
	pc, err := GetPC(g, d)
	if err != nil {
		return 0, "", err
	}
	fn, err := GetFN(g, d)
	if err != nil {
		return 0, "", err
	}
	err = setCurrentFN(g, fn)
	if err != nil {
		return 0, "", err
	}
	err = setCurrentPC(g, pc)
	if err != nil {
		return 0, "", err
	}
	return pc, fn, nil
}

type Trap interface {
	DoTrap(g abstraction.GVM, fn *string, pc *uint64) error
}

type TrapCallFunc struct {
	newFn string
}

func (c TrapCallFunc) DoTrap(g abstraction.GVM, fn *string, pc *uint64) {

}

//TrapCallFunc
func Run(g abstraction.GVM, fn string) (uint64, error) {
	err := pushFrame(g, fn)
	var pc uint64 = 0
	for err == nil {
		pc, err = Continue(g, fn, pc)
		if err != nil {
			return 0, err
		}

		if err == OutOfRange {
			pc, fn, err = popFrame(g)
		} else if trap, ok := err.(Trap); ok {
			err = trap.DoTrap(g, &fn, &pc)
		}
	}
	return pc, err
}

func Iter(g *abstraction.ExecCtx) (err error) {
	if err = validate(g); err != nil {
		return
	}
	inst, err := g.Fetch(g.PC)
	if err != nil {
		return err
	}
	return inst.Exec(g)
}

//var Trap = errors.New("gvm trapped")

var OutOfRange = errors.New("gvm stopped")

func validate(g *abstraction.ExecCtx) error {
	if g.PC >= uint64(g.Len()) {
		return OutOfRange
	}
	return nil
}
