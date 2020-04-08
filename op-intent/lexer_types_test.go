package opintent

import (
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUndefinedEqual(t *testing.T) {
	assert.EqualValues(t, true, Undefined == lexer_types.Undefined)
}
