package lexer

import "github.com/HyperService-Consortium/go-uip/internal/lexer_types"

type IfIntent struct {
	*IntentImpl
	If        *RootIntents      `json:"if"`        // key
	Condition lexer_types.Param `json:"condition"` // key
	Else      *RootIntents      `json:"else"`      // option
}
