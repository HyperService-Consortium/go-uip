package lexer_types

import "github.com/HyperService-Consortium/go-uip/internal/token_types"

var _ Param = &StateVariable{}
var _ Param = BinaryExpression{}
var _ Param = UnaryExpression{}
var _ Param = LocalStateVariable{}

var _ token_types.StateVariableI = &StateVariable{}
var _ token_types.BinaryExpressionI = &DeterminedBinaryExpression{}
var _ token_types.UnaryExpressionI = &DeterminedUnaryExpression{}
var _ token_types.LocalStateVariableI = LocalStateVariable{}

var _ token_types.NamespacedRawAccountI = &NamespacedRawAccount{}
var _ token_types.NamespacedNameAccountI = &NamespacedNameAccount{}
var _ token_types.RawAccountI = &RawAccount{}
var _ token_types.NameAccountI = NameAccount{}
var _ token_types.FullAccountI = FullAccount{}

var _ token_types.Param = &StateVariable{}
var _ token_types.Param = &DeterminedBinaryExpression{}
var _ token_types.Param = &DeterminedUnaryExpression{}
var _ token_types.Param = &LocalStateVariable{}
