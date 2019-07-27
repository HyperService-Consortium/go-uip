package opintent

import (
	"errors"
	"fmt"
	"sync/atomic"

	queue "github.com/Myriad-Dreamin/go-uip/queue"
)

func (ier *OpIntentInitializer) TopologicalSort(
	transactionIntents [][]*TransactionIntent,
	dependencies []Dependency,
) error {
	niz := len(transactionIntents)
	miz := len(dependencies)
	if niz > maxIntentSizeLimit || (miz>>2) > maxIntentSizeLimit {
		return errors.New("too large")
	}

	if niz > 65000 || miz > 64500 {
		if atomic.LoadUint32(&ier.largerThanLarger) == 1 {
			return errors.New("more than a large work in this server")
		}
		atomic.StoreUint32(&ier.largerThanLarger, 1)
		deg := make([]uint32, niz)
		hed := make([]uint32, niz)
		nxt := make([]uint32, miz+2)
		lnk := make([]uint32, miz+2)
		var mal uint32 = 1

		for _, dep := range dependencies {
			lnk[mal] = dep.Dst
			nxt[mal] = hed[dep.Src]
			hed[dep.Src] = mal
			mal++
		}

		Q := queue.NewUnsafeQueue(uint32(niz + 2))

		var succ bool
		for idx, degC := range deg {
			if degC == 0 {
				succ = succ && Q.Put(uint32(idx))
			}
		}
		if !succ {
			return errors.New("put error")
		}

		var u, v, sel uint32
		for Q.Size() != 0 {
			u, succ = Q.Get()
			if !succ {
				return errors.New("get error")
			}
			for idx := hed[u]; idx != 0; idx = nxt[idx] {
				v = lnk[idx]
				deg[v]--
				if deg[v] == 0 {
					succ = succ && Q.Put(uint32(v))
				}
			}
			hed[u] = sel
			sel++
			if !succ {
				return errors.New("put error")
			}
		}
		if int(sel) == niz {
			for idx := 0; idx < niz; idx++ {
				for hed[idx] != uint32(idx) {
					h := hed[idx]
					transactionIntents[idx], transactionIntents[h] = transactionIntents[h], transactionIntents[idx]
					hed[h], hed[idx] = hed[idx], hed[h]
				}
			}
		}
		deg = nil
		_ = deg
		hed = nil
		_ = hed
		nxt = nil
		_ = nxt
		lnk = nil
		_ = lnk
		atomic.StoreUint32(&ier.largerThanLarger, 0)

		if int(sel) != niz {
			return errors.New("sort error")
		}
		return nil
	}
	deg, hed, nxt, lnk := ier.degPool.NewGSpace(uint16(niz), uint16(miz+2))
	var mal uint16 = 1

	for _, dep := range dependencies {
		deg[dep.Dst]++
		lnk[mal] = uint16(dep.Dst)
		nxt[mal] = hed[dep.Src]
		hed[dep.Src] = mal
		mal++
	}

	Q := queue.NewUnsafeQueue16(uint16(niz + 2))

	var succ = true
	for idx := 0; idx < niz; idx++ {
		if deg[idx] == 0 {
			succ = succ && Q.Put(uint16(idx))
		}
	}
	if !succ {
		return errors.New("put error")
	}

	var u, v, sel uint16
	for Q.Size() != 0 {
		u, succ = Q.Get()
		if !succ {
			return errors.New("get error")
		}
		fmt.Println("get", u)

		for idx := hed[u]; idx != 0; idx = nxt[idx] {
			v = lnk[idx]
			// fmt.Println(v, deg[v], hed[u], idx)
			deg[v]--
			if deg[v] == 0 {
				succ = succ && Q.Put(v)
			}
		}
		hed[u] = sel
		sel++
		if !succ {
			return errors.New("put error")
		}
	}

	if int(sel) == niz {
		for idx := 0; idx < niz; idx++ {
			for hed[idx] != uint16(idx) {
				h := hed[idx]
				transactionIntents[idx], transactionIntents[h] = transactionIntents[h], transactionIntents[idx]
				hed[h], hed[idx] = hed[idx], hed[h]
			}
		}
	}

	ier.degPool.FreeVector(deg)
	ier.degPool.FreeVector(hed)
	ier.degPool.FreeVector(nxt)
	ier.degPool.FreeVector(lnk)

	if int(sel) != niz {
		return errors.New("sort error")
	}

	// var dependencyGrapher Grapher
	return nil
}
