package opintent

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type ParseError struct {
	Err          error
	ErrType      string
	Descriptions []fmt.Stringer
}

type JSONParseError struct {
	Err          string   `json:"error"`
	ErrType      string   `json:"type"`
	Descriptions []string `json:"description"`
}

func _newParseError(err error, errType string) *ParseError {
	return &ParseError{Err: err, ErrType: errType}
}

func newSortError(err error) *ParseError {
	return _newParseError(err, "sort error")
}

func newUnmarshalError(err error) *ParseError {
	return _newParseError(err, "unmarshal error")
}

func newMarshalError(err error) *ParseError {
	return _newParseError(err, "marshal error")
}

func newInvalidFieldError(err error) *ParseError {
	return _newParseError(err, "invalid field error")
}

func newGetAccountFailed(err error) *ParseError {
	return _newParseError(err, "get account failed")
}

func newGenPaymentError(err error) *ParseError {
	return _newParseError(err, "generate payment tx error")
}

func newGetTranslatorError(err error) *ParseError {
	return _newParseError(err, "get translator error")
}

func newGetTransactionProofType(err error) *ParseError {
	return _newParseError(err, "get transaction proof error")
}

func newParseTransactionIntentError(err error) *ParseError {
	return _newParseError(err, "translator parse transaction intent error")
}

func newValueTypeNotFound(valueType string) *ParseError {
	return _newParseError(ValueTypeNotFound{ValueType: valueType}, "value type not found")
}

func newFieldNotFound(field string) *ParseError {
	return _newParseError(FieldNotFound{Field: field}, "field not found")
}

func newOpNameNotFound(opName string) *ParseError {
	return _newParseError(OpNameNotFound{OpName: opName}, "op name not found")
}

func newDecodeDomainError(err error) *ParseError {
	return _newParseError(err, "decode domain error")
}

func newGetDomainError(err error) *ParseError {
	return _newParseError(err, "get domain error")
}


func newDecodeContractAddressError(err error) *ParseError {
	return _newParseError(err, "decode contract address error")
}

func newDecodeContractPosError(err error) *ParseError {
	return _newParseError(err, "decode contract pos error")
}

func newNotEnoughParamInformation() *ParseError {
	return _newParseError(ErrNotEnoughParamInformation, "not enough param information")
}

func (p ParseError) Error() string {
	return p.Err.Error()
}

func (p ParseError) Serialize() ([]byte, error) {
	var jpe JSONParseError
	jpe.Descriptions = make([]string, len(p.Descriptions))
	for i := range p.Descriptions {
		jpe.Descriptions[i] = p.Descriptions[i].String()
	}
	jpe.ErrType = p.ErrType
	jpe.Err = p.Err.Error()
	return json.Marshal(&jpe)
}

func (p ParseError) ToError() (pe error, err error) {
	x, err := p.Serialize()
	if err != nil {
		return
	}
	return errors.New(string(x)), nil
}

func (p *ParseError) Desc(d fmt.Stringer) *ParseError {
	p.Descriptions = append(p.Descriptions, d)
	return p
}

type AtOpIntentsPos struct{ Pos int }

func (a AtOpIntentsPos) String() string {
	return fmt.Sprintf("at op intents pos %d", a.Pos)
}

type AtChainID struct{ ChainID uip.ChainIDUnderlyingType }

func (a AtChainID) String() string {
	return fmt.Sprintf("at chain id %d", a.ChainID)
}

type AtOpIntentField struct{ Field string }

func (a AtOpIntentField) String() string {
	return fmt.Sprintf("at op intents field %s", a.Field)
}

//func (a AtOpIntentsPos) String() string {
//	return fmt.Sprintf("at op intents pos %d", a.pos)
//}
