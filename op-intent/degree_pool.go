package opintent

import (
	"sync"
)

const (
	smalll             = 30
	smalle             = 31 * 31 / 2
	smalllCount        = 5000
	medium             = 250
	mediue             = 252 * 252 / 2
	mediumCount        = 100
	larger             = 65000
	largerCount        = 10
	maxIntentSizeLimit = 4 * 1024 * 1024
)

var (
	zeroVector [65537]uint16
	smalllPool [smalllCount][smalll]uint16
	mediumPool [mediumCount][medium]uint16
	largerPool [largerCount][larger]uint16

	smallePool [smalllCount][smalle]uint16
	mediuePool [mediumCount][mediue]uint16
	// (400000 + 650000) * 2
	//
)

type subDegreePool struct {
	resrc chan []uint16
}

func (s *subDegreePool) newVector(n uint16) []uint16 {
	var ret = <-s.resrc
	copy(ret, zeroVector[:n])
	return ret
}

func (s *subDegreePool) freeVector(vec []uint16) {
	if cap(s.resrc) > len(s.resrc) {
		s.resrc <- vec
	}
	return
}

var smalllDegPool *subDegreePool
var smalllOnce sync.Once

func newSmalllDegreePool() *subDegreePool {
	smalllOnce.Do(func() {
		smalllDegPool = &subDegreePool{resrc: make(chan []uint16, smalllCount)}
		for idx := uint16(0); idx < smalllCount; idx++ {
			smalllDegPool.resrc <- smalllPool[idx][:]
		}
	})

	return smalllDegPool
}

var mediumDegPool *subDegreePool
var mediumOnce sync.Once

func newMediumDegreePool() *subDegreePool {
	mediumOnce.Do(func() {
		mediumDegPool = &subDegreePool{resrc: make(chan []uint16, mediumCount)}
		for idx := uint16(0); idx < mediumCount; idx++ {
			mediumDegPool.resrc <- mediumPool[idx][:]
		}
	})
	return mediumDegPool
}

var largerDegPool *subDegreePool
var largerOnce sync.Once

func newLargerDegreePool() *subDegreePool {
	largerOnce.Do(func() {
		largerDegPool = &subDegreePool{resrc: make(chan []uint16, largerCount)}
		for idx := uint16(0); idx < largerCount; idx++ {
			largerDegPool.resrc <- largerPool[idx][:]
		}
	})
	return largerDegPool
}

var smalleDegPool *subDegreePool
var smalleOnce sync.Once

func newSmalleDegreePool() *subDegreePool {
	smalleOnce.Do(func() {
		smalleDegPool = &subDegreePool{resrc: make(chan []uint16, smalllCount)}
		for idx := uint16(0); idx < smalllCount; idx++ {
			smalleDegPool.resrc <- smallePool[idx][:]
		}
	})
	return smalleDegPool
}

var mediueDegPool *subDegreePool
var mediueOnce sync.Once

func newMediueDegreePool() *subDegreePool {
	mediueOnce.Do(func() {
		mediueDegPool = &subDegreePool{resrc: make(chan []uint16, mediumCount)}
		for idx := uint16(0); idx < mediumCount; idx++ {
			mediueDegPool.resrc <- mediuePool[idx][:]
		}
	})
	return mediueDegPool
}

type DegreePool struct {
	sp *subDegreePool
	se *subDegreePool
	mp *subDegreePool
	me *subDegreePool
	lp *subDegreePool
}

func (degPool *DegreePool) NewGSpace(degSize, edgeSize uint16) ([]uint16, []uint16, []uint16, []uint16) {
	switch {
	case degSize < smalll:
		return degPool.sp.newVector(degSize), degPool.sp.newVector(degSize), degPool.se.newVector(edgeSize), degPool.se.newVector(edgeSize)
	case smalll <= degSize && degSize < medium:
		return degPool.mp.newVector(degSize), degPool.mp.newVector(degSize), degPool.me.newVector(edgeSize), degPool.me.newVector(edgeSize)
	case medium <= degSize && degSize < larger:
		return degPool.lp.newVector(degSize), degPool.lp.newVector(degSize), make([]uint16, edgeSize), make([]uint16, edgeSize)
	default:
		// warning: the deg of degSize >= max_uint16
		return make([]uint16, degSize), make([]uint16, degSize), make([]uint16, edgeSize), make([]uint16, edgeSize)
	}
}

func (degPool *DegreePool) FreeVector(degVec []uint16) {
	var degSize = len(degVec)
	switch {
	case degSize < smalll:
	case smalll == degSize:
		degPool.sp.freeVector(degVec)
	case medium == degSize:
		degPool.mp.freeVector(degVec)
	case larger == degSize:
		degPool.lp.freeVector(degVec)
	case smalle == degSize:
		degPool.se.freeVector(degVec)
	case mediue == degSize:
		degPool.me.freeVector(degVec)
	case smalll < degSize && degSize < medium:
		degPool.sp.freeVector(degVec)
	case medium < degSize && degSize < larger:
		degPool.mp.freeVector(degVec)
	case larger < degSize && degSize < smalle:
		degPool.lp.freeVector(degVec)
	case smalle < degSize && degSize < mediue:
		degPool.se.freeVector(degVec)
	case mediue < degSize:
		degPool.me.freeVector(degVec)
	default:
		degVec = nil
		_ = degVec
	}
}

var degPool *DegreePool
var degOnce sync.Once

func newDegreePool() *DegreePool {
	degOnce.Do(func() {
		degPool = &DegreePool{
			sp: newSmalllDegreePool(),
			mp: newMediumDegreePool(),
			lp: newLargerDegreePool(),
			se: newSmalleDegreePool(),
			me: newMediueDegreePool(),
		}
	})
	return degPool
}
