package opintent

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

type Ints []int

func (is Ints) Len() int {
	return len(is)
}

func (is Ints) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func TestSort(t *testing.T) {
	var is = Ints{0, 1, 2}
	Dep := []Dependency{
		{1, 0},
		{2, 0},
	}
	ier := sugar.HandlerError(NewInitializer(
		uip.BlockChainGetterNilImpl{}, uip.AccountBaseNilImpl{})).(*Initializer)
	sugar.HandlerError0(ier.TopologicalSort(is, Dep))

	for _, i := range is {
		fmt.Println(i)
	}
}

func TestSort2(t *testing.T) {
	var is = Ints{0, 1, 2, 3}
	Dep := []Dependency{
		{1, 0},
		{2, 0},
		{2, 1},
		{3, 2},
	}
	ier := sugar.HandlerError(NewInitializer(
		uip.BlockChainGetterNilImpl{}, uip.AccountBaseNilImpl{})).(*Initializer)
	sugar.HandlerError0(ier.TopologicalSort(is, Dep))

	for _, i := range is {
		fmt.Println(i)
	}
}
