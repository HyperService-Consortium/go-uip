package libgvm

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

func ConvertError(fr, to abstraction.Ref) (err error) {
	return fmt.Errorf("cant convert %s and %s to the same type", ExplainGVMType(fr.GetGVMType()), ExplainGVMType(to.GetGVMType()))
}

func ConvertUnsignedError(k abstraction.Ref) error {
	return fmt.Errorf("cant convert %s to the unsigned type", ExplainGVMType(k.GetGVMType()))
}

func InvalidTypeError(k abstraction.Ref) error {
	return fmt.Errorf("invalid type: %v", ExplainGVMType(k.GetGVMType()))
}

var OutOfRange = errors.New("gvm stopped")

var StopUnderFlow = errors.New("depth underflow")
