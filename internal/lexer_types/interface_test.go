package lexer_types

import "github.com/HyperService-Consortium/go-uip/internal/token_types"

var _ Param = &StateVariable{}
var _ Param = BinaryExpression{}
var _ Param = UnaryExpression{}
var _ Param = LocalStateVariable{}

var _ token_types.Param = &StateVariable{}
var _ token_types.Param = &DeterminedBinaryExpression{}
var _ token_types.Param = &DeterminedUnaryExpression{}
var _ token_types.Param = &LocalStateVariable{}
