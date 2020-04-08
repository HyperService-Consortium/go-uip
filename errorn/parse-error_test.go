package errorn

import (
	"fmt"
	"testing"
)

func TestParseError_Error(t *testing.T) {
	fmt.Println((&ParseError{Err: InvalidOpType, Descriptions: []fmt.Stringer{
		AtOpIntentsPos{Pos: 2},
	}}).ToError())
}
