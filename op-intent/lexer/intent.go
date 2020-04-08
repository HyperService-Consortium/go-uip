package lexer

import (
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
)

type IntentImpl struct {
	Name   string          `json:"name"`
	OpType token_type.Type `json:"op_type"`
}

func (r IntentImpl) GetType() token_type.Type {
	return r.OpType
}

func (r IntentImpl) GetName() string {
	return r.Name
}

func (r *IntentImpl) UnmarshalDocument(content document.Document) error {

	name, opType := content.Get(FieldOpIntentsName), content.Get(FieldKeyType)
	if !name.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsName)
	}
	r.Name = name.String()

	if !opType.Exists() {
		return errorn.NewFieldNotFound(FieldKeyType)
	}

	switch opType.String() {
	case "Payment":
		r.OpType = token_type.Pay
	case "ContractInvocation":
		r.OpType = token_type.Invoke
	case "IfStatement":
		r.OpType = token_type.If
	case "loopFunction":
		r.OpType = token_type.Loop
	default:
		return errorn.NewInvalidFieldError(errorn.InvalidOpType)
	}
	return nil
}
