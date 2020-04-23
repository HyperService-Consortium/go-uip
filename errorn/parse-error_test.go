package errorn

import (
	"errors"
	"github.com/Myriad-Dreamin/catcher"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
)

func TestParseError_Error(t *testing.T) {
	x, e := NewInvalidFieldError(InvalidOpType).Desc(AtOpIntentsPos{Pos: 2}).ToError()
	assert.NoError(t, e)
	assert.EqualValues(t,
		errors.New(`{"error":"\u003c84f4446f\u003epos:\u003c\u003cgithub.com/HyperService-Consortium/go-uip/errorn.TestParseError_Error,C:/work/go/src/github.com/HyperService-Consortium/go-uip/errorn/parse-error_test.go:11\u003e,C:/work/go/src/github.com/HyperService-Consortium/go-uip/errorn/parse-error_test.go:12\u003e,\u003c84f4446f\u003ecode:1,\u003c84f4446f\u003eerr:there is at least an unexpected op_type in OpIntents","type":"invalid field error","description":["at op intents pos 2"]}`), x)
	stack, ok := catcher.StackFromString(gjson.Get(x.Error(), "error").String())
	assert.EqualValues(t, true, ok)
	assert.EqualValues(t, `1 <- <pos:<<github.com/HyperService-Consortium/go-uip/errorn.TestParseError_Error,C:/work/go/src/github.com/HyperService-Consortium/go-uip/errorn/parse-error_test.go:11>,C:/work/go/src/github.com/HyperService-Consortium/go-uip/errorn/parse-error_test.go:12>,code:1,err:there is at least an unexpected op_type in OpIntents>
`, stack.String())
}
