package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
)

type IntentImpl struct {
	Name   string     `json:"name"`
	OpType token.Type `json:"op_type"`
}

func (r IntentImpl) GetType() token.Type {
	return r.OpType
}

func (r IntentImpl) GetName() string {
	return r.Name
}

func (r *IntentImpl) UnmarshalDocument(content document.Document) error {

	name, opType := content.Get(FieldOpIntentsName), content.Get(FieldOpIntentsOpType)
	if !name.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsName)
	}
	r.Name = name.String()

	if !opType.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsOpType)
	}

	switch opType.String() {
	case "Payment":
		r.OpType = token.Pay
	case "ContractInvocation":
		r.OpType = token.Invoke
	case "IfStatement":
		r.OpType = token.If
	case "loopFunction":
		r.OpType = token.Loop
	default:
		return errorn.NewInvalidFieldError(errorn.InvalidOpType)
	}
	return nil
}

