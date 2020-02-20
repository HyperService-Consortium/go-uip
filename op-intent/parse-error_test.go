package opintent

import (
	"fmt"
	"testing"
)

func TestParseError_Error(t *testing.T) {
	fmt.Println((&ParseError{Err:invalidOpType, Descriptions:[]fmt.Stringer{
		AtOpIntentsPos{Pos: 2},
	}}).ToError())
}