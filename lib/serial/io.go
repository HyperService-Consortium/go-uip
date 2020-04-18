package serial

import (
	"encoding/binary"
	"io"
	"math/big"
)

func Write(w io.Writer, v interface{}, err *error) {
	if *err != nil {
		return
	}
	switch v := v.(type) {
	case string:
		*err = binary.Write(w, binary.BigEndian, uint64(len(v)))
		if *err != nil {
			return
		}

		*err = binary.Write(w, binary.BigEndian, []byte(v))
		return
	case []byte:
		*err = binary.Write(w, binary.BigEndian, uint64(len(v)))
	}
	if *err != nil {
		return
	}

	*err = binary.Write(w, binary.BigEndian, v)
	return
}

func Read(r io.Reader, target interface{}, err *error) {
	if *err != nil {
		return
	}
	switch t := target.(type) {
	case *string:
		var dataLength uint64
		*err = binary.Read(r, binary.BigEndian, &dataLength)
		if *err != nil {
			return
		}
		var b = make([]byte, dataLength)
		*err = binary.Read(r, binary.BigEndian, &b)
		if *err != nil {
			return
		}
		*t = string(b)
		return
	case *[]byte:
		var dataLength uint64
		*err = binary.Read(r, binary.BigEndian, &dataLength)
		if *err != nil {
			return
		}
		*t = make([]byte, dataLength)
		*err = binary.Read(r, binary.BigEndian, t)
		if *err != nil {
			return
		}
		return
	}
	*err = binary.Read(r, binary.BigEndian, target)
	if *err != nil {
		return
	}
}

func EncodeBigInt(b *big.Int) ([]byte, error) {
	if b == nil {
		return nil, nil
	}
	return b.Bytes(), nil
}

func WriteBigInt(w io.Writer, i *big.Int, err *error) {
	if i == nil {
		Write(w, []byte(nil), err)
		return
	}
	Write(w, i.Bytes(), err)
}

func ReadBigInt(r io.Reader, i *big.Int, err *error) {
	var b []byte
	Read(r, &b, err)
	i.SetBytes(b)
}
