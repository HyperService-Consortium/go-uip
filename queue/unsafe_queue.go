// esQueue
package queue

import (
	"fmt"
	"runtime"
)

// UnsafeQueue
type UnsafeQueue struct {
	capaciity uint32
	capMod    uint32
	putPos    uint32
	getPos    uint32
	cache     []uint32
}

func NewUnsafeQueue(capaciity uint32) *UnsafeQueue {
	q := new(UnsafeQueue)
	q.capaciity = minQuantity(capaciity)
	q.capMod = q.capaciity - 1
	q.putPos = 0
	q.getPos = 0
	q.cache = make([]uint32, q.capaciity)
	return q
}

func (q *UnsafeQueue) String() string {
	return fmt.Sprintf("Queue{capaciity: %v, capMod: %v, putPos: %v, getPos: %v}",
		q.capaciity, q.capMod, q.getPos, q.getPos)
}

func (q *UnsafeQueue) Capaciity() uint32 {
	return q.capaciity
}

func (q *UnsafeQueue) Quantity() uint32 {
	if q.putPos >= q.getPos {
		return q.putPos - q.getPos
	}
	return q.capMod + (q.putPos - q.getPos)
}

func (q *UnsafeQueue) Size() uint32 {
	if q.putPos >= q.getPos {
		return q.putPos - q.getPos
	}
	return q.capMod + (q.putPos - q.getPos)
}

// put queue functions
func (q *UnsafeQueue) Put(val uint32) bool {

	if q.putPos+2 == q.getPos {
		runtime.Gosched()
		return false
	}
	q.cache[q.putPos&q.capMod] = val
	q.putPos++
	return true
}

// get queue functions
func (q *UnsafeQueue) Get() (val uint32, ok bool) {
	if q.putPos == q.getPos {
		runtime.Gosched()
		return 0, false
	}
	val = q.cache[q.getPos&q.capMod]
	q.getPos++
	return val, true
}
