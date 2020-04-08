package isc

import (
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/lib/math"
)

type Code int

const (
	CodeOK Code = iota
	CodeEncodeDataError
	CodeEmptyOwners
	CodeFirstOwnerNotBeSender
	CodeDuplicateOwner
	CodeNotEqualLengthOfOwnersAndFunds
	CodeTransactionNotActive
	CodeActionNotActive
	CodeGetError
	CodeSetError
	CodeMarshalError

	CodeIsNotInitializing
	CodeIsNotInitialized
	CodeIsNotOpening
	CodeIsNotSettling

	CodePCUnderflow
	CodeIteratePCError
)

var Description = map[Code]string{
	CodeOK:                             "ok",
	CodeEncodeDataError:                "encode data error",
	CodeEmptyOwners:                    "len(owners) should not be 0",
	CodeFirstOwnerNotBeSender:          "first owner must be sender",
	CodeDuplicateOwner:                 "the address is already the isc owner",
	CodeNotEqualLengthOfOwnersAndFunds: "the length of owners is not equal to that of funds",
	CodeTransactionNotActive:           "this transaction is not active",
	CodeActionNotActive:                "this action is not active",
	CodeGetError:                       "get error",
	CodeSetError:                       "set error",
	CodeMarshalError:                   "marshal error",
	CodeIsNotInitializing:              "the contract is not initializing",
	CodeIsNotInitialized:               "the contract is not initialized",
	CodeIsNotOpening:                   "the contract is not opening",
	CodeIsNotSettling:                  "the contract is not settling",
	CodePCUnderflow:                    "pc pointer underflow",
	CodeIteratePCError:                 "iterate pc error",
}

type Response interface {
	GetCode() Code
}

func reply() *ResponseData {
	return &ResponseData{}
}

func report(code Code, err error) Response {
	return &ResponseError{Code: code, Err: err.Error()}
}

func reportString(code Code, err string) Response {
	return &ResponseError{Code: code, Err: err}
}

func reportCode(code Code) Response {
	return &ResponseError{Code: code, Err: ""}
}

func IsOK(r Response) bool {
	return r.GetCode() == CodeOK
}

func IsErr(r Response) bool {
	return r.GetCode() != CodeOK
}

type ResponseData struct {
	Data    []byte
	Value   *math.Uint256
	OutFlag bool
}

func (r *ResponseData) ImplISCResponse() Response {
	return r
}

// Account to Contract
func (r *ResponseData) TransferToC(v *math.Uint256) *ResponseData {
	r.Value = v
	r.OutFlag = false
	return r
}

// Contract to Account
func (r *ResponseData) TransferToA(v *math.Uint256) *ResponseData {
	r.Value = v
	r.OutFlag = true
	return r
}

// Data
func (r *ResponseData) Param(v interface{}) Response {
	var err error
	r.Data, err = json.Marshal(v)
	if err != nil {
		return report(CodeEncodeDataError, err)
	}
	return r
}

func (r ResponseData) GetCode() Code {
	return CodeOK
}

type ResponseError struct {
	Code Code
	Err  string
}

func (r ResponseError) ImplISCResponse() Response {
	return r
}

func (r ResponseError) GetCode() Code {
	return r.Code
}

func (r ResponseError) Error() string {
	return r.Err
}

type nilResponse struct{}

func (r *nilResponse) ImplISCResponse() Response {
	return r
}

func (r *nilResponse) GetCode() Code {
	return 0
}

func assertTrue(k bool, hintCode Code) {
	if !k {
		panic(reportCode(hintCode))
	}
}

func assertTrueH(k bool, hintCode Code, hint ...interface{}) {
	if !k {
		panic(reportString(hintCode, fmt.Sprint(hint...)))
	}
}

var Nil *nilResponse
var OK = new(ResponseData)
