package queue

import (
	"fmt"
	"runtime"
)

// UnsafeQueue16
type UnsafeQueue16 struct {
	capaciity uint32
	capMod    uint16
	putPos    uint16
	getPos    uint16
	cache     []uint16
}

func NewUnsafeQueue16(capaciity uint16) *UnsafeQueue16 {
	q := new(UnsafeQueue16)
	q.capaciity = minQuantity(uint32(capaciity))
	q.capMod = uint16(q.capaciity - 1)
	q.putPos = 0
	q.getPos = 0
	q.cache = make([]uint16, q.capaciity)
	return q
}

func (q *UnsafeQueue16) String() string {
	return fmt.Sprintf("Queue{capaciity: %v, capMod: %v, putPos: %v, getPos: %v}",
		q.capaciity, q.capMod, q.getPos, q.getPos)
}

func (q *UnsafeQueue16) Capaciity() uint32 {
	return q.capaciity
}

func (q *UnsafeQueue16) Quantity() uint16 {
	if q.putPos >= q.getPos {
		return q.putPos - q.getPos
	}
	return q.capMod + (q.putPos - q.getPos)
}

func (q *UnsafeQueue16) Size() uint16 {
	if q.putPos >= q.getPos {
		return q.putPos - q.getPos
	}
	return q.capMod + (q.putPos - q.getPos)
}

// put queue functions
func (q *UnsafeQueue16) Put(val uint16) bool {

	if q.putPos+2 == q.getPos {
		runtime.Gosched()
		return false
	}
	q.cache[q.putPos&q.capMod] = val
	q.putPos++
	return true
}

// get queue functions
func (q *UnsafeQueue16) Get() (val uint16, ok bool) {
	if q.putPos == q.getPos {
		runtime.Gosched()
		return 0, false
	}
	val = q.cache[q.getPos&q.capMod]
	q.getPos++
	return val, true
}
