package libgvm

import (
	"bytes"
	"encoding/binary"
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
		return 0, err
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
		return "", err
	}
	return r.Unwrap().(string), nil
}

// GetCurrentDepth return the value of current depth
func GetCurrentDepth(g abstraction.Machine) (uint64, error) {
	r, err := g.Load(fDepth, RefUint64)
	if err != nil {
		return 0, err
	}
	if r == Undefined {
		return 0, nil
	}
	return r.Unwrap().(uint64), nil
}

func setPC(g abstraction.Machine, dep, pc uint64) error {
	return g.Save(fPCPrefix+formatUintSuffix(dep), Uint64(pc))
}

func deletePC(g abstraction.Machine, dep uint64) error {
	return g.Delete(fPCPrefix + formatUintSuffix(dep))
}

func setCurrentPC(g abstraction.Machine, pc uint64) error {
	return g.Save(fCurrentPC, Uint64(pc))
}

func setFN(g abstraction.Machine, dep uint64, fn string) error {
	return g.Save(fFNPrefix+formatUintSuffix(dep), String(fn))
}

func deleteFN(g abstraction.Machine, dep uint64) error {
	return g.Delete(fFNPrefix + formatUintSuffix(dep))
}

func setCurrentFN(g abstraction.Machine, fn string) error {
	return g.Save(fCurrentFN, String(fn))
}

func setCurrentDepth(g abstraction.Machine, dep uint64) error {
	return g.Save(fDepth, Uint64(dep))
}

func setCurrentState(g *abstraction.ExecCtx) error {
	err := setCurrentDepth(g, g.Depth)
	if err != nil {
		return err
	}
	err = setCurrentFN(g, g.FN)
	if err != nil {
		return err
	}
	err = setCurrentPC(g, g.PC)
	if err != nil {
		return err
	}
	return nil
}

func loadLocals(g abstraction.Machine, d uint64) (abstraction.Locals, error) {

	r, err := g.Load("_gvm_locals_"+strconv.FormatUint(d, 16), RefBytes)
	if err != nil {
		return nil, err
	}
	return deserializeMapSR(g, r.Unwrap().([]byte))
}

func deleteLocals(g abstraction.Machine, d uint64) error {

	return g.Delete("_gvm_locals_" + strconv.FormatUint(d, 16))
}

func saveLocals(g abstraction.Machine, d uint64, locals abstraction.Locals) error {
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
