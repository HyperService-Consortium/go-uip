package gvm

type PackPrototype interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

func Encode(variable Ref) ([]byte, error) { return EncodeG(instance, variable) }

func DecodeR(x []byte, t RefType) (interface{}, error) { return DecodeGR(instance, x, t) }

func Decode(x []byte, t RefType) (Ref, error) { return DecodeG(instance, x, t) }

func EncodeG(g GVM, variable Ref) ([]byte, error) {
	if i, err := g.GetPackPrototype(variable.GetGVMType()); err != nil {
		return nil, err
	} else {
		return i.Encode(variable.Unwrap())
	}
}

func DecodeGR(g GVM, x []byte, t RefType) (interface{}, error) {
	if i, err := g.GetPackPrototype(t); err != nil {
		return nil, err
	} else {
		return i.Decode(x)
	}
}

func DecodeG(g GVM, x []byte, t RefType) (Ref, error) {
	if b, err := DecodeGR(g, x, t); err != nil {
		return nil, err
	} else {
		return ref{
			t:  RefType(t),
			v: b,
		}, nil
	}
}

