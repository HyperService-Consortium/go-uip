package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
)

type BoolVariable = document.Document

type IfIntent struct {
	*IntentImpl
	If        *RootIntents `json:"if"`        // key
	Condition BoolVariable `json:"condition"` // key
	Else      *RootIntents `json:"else"`      // option
}

