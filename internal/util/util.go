package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func BytesToBytes(bt []byte) []byte {
	return bt
}

func BytesToBytesHelper(bt interface{}) []byte {
	return bt.([]byte)
}

func BytesToBytesHelperR(bt []byte) interface{} {
	return bt
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

func Int32ToBytesHelper(i interface{}) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i.(int32)))
	return buf
}

func BytesToInt32Helper(buf []byte) interface{} {
	return int32(binary.BigEndian.Uint32(buf))
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToUint64(buf []byte) uint64 {
	return uint64(binary.BigEndian.Uint64(buf))
}

func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func Uint16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func BytesToUint16(buf []byte) uint16 {
	return binary.BigEndian.Uint16(buf)
}


func Uint32ToBytesHelper(i interface{}) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i.(uint32)))
	return buf
}

func Int16ToBytes(i int16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func BytesToInt16(buf []byte) int16 {
	return int16(binary.BigEndian.Uint16(buf))
}

func BytesToUint32Helper(buf []byte) interface{} {
	return uint32(binary.BigEndian.Uint32(buf))
}

func StringToBytes(str string) []byte {
	return []byte(str)
}

func BytesToString(bt []byte) string {
	return string(bt)
}

func ErrorToString(err error) string {
	return fmt.Sprintf("%v", err)
}

func ConcatBytes(dat ...[]byte) []byte {
	var buff bytes.Buffer
	var totlen int
	for _, btdat := range dat {
		buff.Write(btdat)
		totlen += len(btdat)
	}
	return buff.Next(totlen)
}