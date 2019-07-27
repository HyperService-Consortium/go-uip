package opintent

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	txI := make([][]*TransactionIntent, 3)
	txI[0] = make([]*TransactionIntent, 0)
	txI[1] = make([]*TransactionIntent, 1)
	txI[2] = make([]*TransactionIntent, 2)
	Dep := []Dependency{
		Dependency{1, 0},
		Dependency{2, 0},
	}
	ier := NewOpIntentInitializer()
	ier.TopologicalSort(txI, Dep)
	for _, tx := range txI {
		fmt.Println(len(tx))
	}
}

func TestSort2(t *testing.T) {
	txI := make([][]*TransactionIntent, 4)
	txI[0] = make([]*TransactionIntent, 0)
	txI[1] = make([]*TransactionIntent, 1)
	txI[2] = make([]*TransactionIntent, 2)
	txI[3] = make([]*TransactionIntent, 3)
	Dep := []Dependency{
		Dependency{1, 0},
		Dependency{2, 0},
		Dependency{2, 1},
		Dependency{3, 2},
	}
	ier := NewOpIntentInitializer()
	ier.TopologicalSort(txI, Dep)
	for _, tx := range txI {
		fmt.Println(len(tx))
	}
}
