package isc

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	var x Response = Nil
	fmt.Println(x.GetCode(), Nil.GetCode())
}
