package libgvm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"strconv"
)

func formatUintSuffix(v uint64) string {
	return strconv.FormatUint(v, 16)
}

const (
	fCurrentPC = "_gvm_cpc"
	fPCPrefix  = "_gvm_pc"
	fCurrentFN = "_gvm_cfn"
	fFNPrefix  = "_gvm_fn"
	fDepth     = "_gvm_dep"
)

// GetPC return the value of PC at depth `dep`
func GetPC(g abstraction.Machine, dep uint64) (uint64, error) {
	r, err := g.Load(fPCPrefix+formatUintSuffix(dep), RefUint64)
	if err != nil {
		return 0, err
	}
	return r.Unwrap().(uint64), nil
}

// GetCurrentPC return the value of current PC
func GetCurrentPC(g abstraction.Machine) (uint64, error) {
	r, err := g.Load(fCurrentPC, RefUint64)
	if err != nil {
		return 0, nil
	}
	return r.Unwrap().(uint64), nil
}

// GetFN return the function name at depth `dep`
func GetFN(g abstraction.Machine, dep uint64) (string, error) {
	r, err := g.Load(fFNPrefix+formatUintSuffix(dep), RefString)
	if err != nil {
		return "", err
	}
	return r.Unwrap().(string), nil
}

// GetCurrentFN return the value of current function name
func GetCurrentFN(g abstraction.Machine) (string, error) {
	r, err := g.Load(fCurrentFN, RefString)
	if err != nil {
		return "", nil
	}
	return r.Unwrap().(string), nil
}

// GetCurrentDepth return the value of current depth
func GetCurrentDepth(g abstraction.Machine) (uint64, error) {
	r, err := g.Load(fDepth, RefUint64)
	if err != nil {
		return 0, nil
	}
	return r.Unwrap().(uint64), nil
}

func setPC(g abstraction.Machine, dep, pc uint64) error {
	return g.Save(fPCPrefix+formatUintSuffix(dep), Uint64(pc))
}

func setCurrentPC(g abstraction.Machine, pc uint64) error {
	return g.Save(fCurrentPC, Uint64(pc))
}

func setFN(g abstraction.Machine, dep uint64, fn string) error {
	return g.Save(fFNPrefix+formatUintSuffix(dep), String(fn))
}

func setCurrentFN(g abstraction.Machine, fn string) error {
	return g.Save(fCurrentFN, String(fn))
}

func setCurrentDepth(g abstraction.Machine, dep uint64) error {
	return g.Save(fDepth, Uint64(dep))
}

func fetch(g abstraction.Machine, mainFunc string, pc uint64) (f abstraction.Function, inst abstraction.Instruction, err error) {

	f, err = g.GetFunction(mainFunc)
	if err != nil {
		return nil, nil, err
	}
	inst, err = f.Fetch(pc)
	if err != nil {
		return nil, nil, err
	}
	return
}

func loadLocal(g abstraction.Machine, d uint64) (map[string]abstraction.Ref, error) {
	r, err := g.Load("_gvm_locals_"+strconv.FormatUint(d, 16), RefBytes)
	if err != nil {
		return nil, err
	}
	return deserializeMapSR(g, r.Unwrap().([]byte))
}

func saveLocals(g abstraction.Machine, d uint64, locals map[string]abstraction.Ref) error {
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

func pushFrame(g abstraction.Machine, fn string) error {
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

func popFrame(g abstraction.Machine) (uint64, string, error) {
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

func serializeMapSR(g abstraction.Machine, mp map[string]abstraction.Ref) (_ []byte, err error) {
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

func deserializeMapSR(g abstraction.Machine, rb []byte) (mp map[string]abstraction.Ref, err error) {
	var b = bytes.NewBuffer(rb)
	var mpLen, dataLen int32
	err = binary.Read(b, binary.BigEndian, &mpLen)
	if err != nil {
		return
	}

	mp = make(map[string]abstraction.Ref)

	var x []byte
	var pType abstraction.RefType
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

		r, err = g.DecodeRef(pType, x[:dataLen])
		if err != nil {
			return
		}

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

		mp[string(x[:dataLen])] = r
	}
	return
}
