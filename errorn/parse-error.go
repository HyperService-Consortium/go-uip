package errorn

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/catcher"
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
	return &ParseError{Err: catcher.WrapN(catcher.BaseSkip+2, 1, err), ErrType: errType}
}

func NewSortError(err error) *ParseError {
	return _newParseError(err, "sort error")
}

const ErrorTypeUnmarshalError = "unmarshal error"

func NewUnmarshalError(err error) *ParseError {
	return _newParseError(err, ErrorTypeUnmarshalError)
}

const ErrorTypeMarshalError = "marshal error"

func NewMarshalError(err error) *ParseError {
	return _newParseError(err, ErrorTypeMarshalError)
}

const ErrorTypeInvalidField = "invalid field error"

func NewInvalidFieldError(err error) *ParseError {
	return _newParseError(err, ErrorTypeInvalidField)
}

func NewGetAccountFailed(err error) *ParseError {
	return _newParseError(err, "get account failed")
}

func NewGenPaymentError(err error) *ParseError {
	return _newParseError(err, "generate payment tx error")
}

func NewGetTranslatorError(err error) *ParseError {
	return _newParseError(err, "get translator error")
}

func NewGetTransactionProofType(err error) *ParseError {
	return _newParseError(err, "get transaction proof error")
}

func NewNoDeterminedAccount() *ParseError {
	return _newParseError(errors.New("nda"), "no determined account")
}

func NewAccountNotFound(name string, chainID uip.ChainIDUnderlyingType) *ParseError {
	return _newParseError(AccountNotFound{name, chainID}, "account not found")
}

func NewParseTransactionIntentError(err error) *ParseError {
	return _newParseError(err, "translator parse transaction intent error")
}

func NewChainIDNotEqual(u, v uip.ChainIDUnderlyingType) *ParseError {
	return _newParseError(ChainIDNotEqual{u, v}, "chain id not equal")
}

type AccountIndexConflict struct {
	Name    string
	ChainID uip.ChainIDUnderlyingType
}

func (a AccountIndexConflict) Error() string {
	return fmt.Sprintf("account indexing conflict:at <name:%v, chain_id:%v>", a.Name, a.ChainID)
}

func NewAccountIndexConflict(name string, chainID uip.ChainIDUnderlyingType) *ParseError {
	return _newParseError(AccountIndexConflict{name, chainID}, "account indexing conflict")
}

func NewNoDeterminedChainID() *ParseError {
	return _newParseError(ErrNoDeterminedChainID, ErrNoDeterminedChainID.Error())
}

const ErrorTypeValueTypeNotFound = "value type not found"

func NewValueTypeNotFound(valueType string) *ParseError {
	return _newParseError(ValueTypeNotFound{ValueType: valueType}, ErrorTypeValueTypeNotFound)
}

const ErrorTypeTokenTypeNotFound = "token type not found"

func NewTokenTypeNotFound(tokenType int) *ParseError {
	return _newParseError(TokenTypeNotFound{TokenType: tokenType}, ErrorTypeTokenTypeNotFound)
}

const ErrorTypeGVMTypeNotFound = "gvm type not found"

func NewGVMTypeNotFound(gvmType int) *ParseError {
	return _newParseError(GVMTypeNotFound{GVMType: gvmType}, ErrorTypeGVMTypeNotFound)
}

const ErrorTypeAccountTypeNotFound = "account type not found"

func NewAccountTypeNotFound(t int) *ParseError {
	return _newParseError(AccountTypeNotFound{AccountType: t}, ErrorTypeAccountTypeNotFound)
}

const ErrorTypeFieldNotFound = "field not found"

func NewFieldNotFound(field string) *ParseError {
	return _newParseError(FieldNotFound{Field: field}, ErrorTypeFieldNotFound)
}

func NewSignTypeNotFound(sign string) *ParseError {
	return _newParseError(SignNotFound{Sign: sign}, "sign not found")
}

func NewOpNameNotFound(opName string) *ParseError {
	return _newParseError(OpNameNotFound{OpName: opName}, "op name not found")
}

const ErrorTypeDecodeDomainError = "decode domain error"

func NewDecodeDomainError(err error) *ParseError {
	return _newParseError(err, ErrorTypeDecodeDomainError)
}

func NewInvalidDocumentType(hint string) *ParseError {
	return _newParseError(fmt.Errorf("invalid document type: %v", hint), "invalid document type")
}

func NewGetDomainError(err error) *ParseError {
	return _newParseError(err, "get domain error")
}

const ErrorTypeDecodeAddressError = "decode address error"

func NewDecodeAddressError(err error) *ParseError {
	return _newParseError(err, ErrorTypeDecodeAddressError)
}

const ErrorTypeDecodeHexError = "decode hex error"

func NewDecodeHexError(err error) *ParseError {
	return _newParseError(err, ErrorTypeDecodeHexError)
}

const ErrorTypeDecodeContractPosError = "decode contract pos error"

func NewDecodeContractPosError(err error) *ParseError {
	return _newParseError(err, ErrorTypeDecodeContractPosError)
}

func NewNotEnoughParamInformation() *ParseError {
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

type AtOpIntentParameterPos struct{ Pos int }

func (a AtOpIntentParameterPos) String() string {
	return fmt.Sprintf("at parameter pos %d", a.Pos)
}

type WithAccountName struct{ Acc string }

func (a WithAccountName) String() string {
	return fmt.Sprintf("with account name %s", a.Acc)
}

//func (a AtOpIntentsPos) String() string {
//	return fmt.Sprintf("at op intents pos %d", a.pos)
//}
