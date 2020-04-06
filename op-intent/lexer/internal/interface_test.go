package internal

import "github.com/HyperService-Consortium/go-uip/op-intent/token"

var _ Param = ConstantVariable{}
var _ Param = StateVariable{}
var _ Param = BinaryExpression{}
var _ Param = UnaryExpression{}
var _ Param = LocalStateVariable{}

var _ token.Param = DeterminedBinaryExpression{}
var _ token.Param = DeterminedUnaryExpression{}
