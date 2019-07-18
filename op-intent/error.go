package opintent

import (
	"errors"
	"fmt"
)

var (
	invalidOpType = errors.New("there is at least an unexpected op_type in OpIntents")
)

func initializeError(keyAttr string) error {
	return fmt.Errorf("the attribute %v must be included in the Payment intent", keyAttr)
}
