package lexer


type IfIntent struct {
	*IntentImpl
	If        *RootIntents `json:"if"`        // key
	Condition Param `json:"condition"` // key
	Else      *RootIntents `json:"else"`      // option
}

