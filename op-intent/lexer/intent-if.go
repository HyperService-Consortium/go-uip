package lexer

import "github.com/HyperService-Consortium/go-uip/op-intent/lexer/internal"

type IfIntent struct {
	*IntentImpl
	If        *RootIntents   `json:"if"`        // key
	Condition internal.Param `json:"condition"` // key
	Else      *RootIntents   `json:"else"`      // option
}
