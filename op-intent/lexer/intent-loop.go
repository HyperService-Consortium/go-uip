package lexer

type LoopIntent struct {
	*IntentImpl
	Loop  *RootIntents `json:"loop"`     // key
	Times int64        `json:"loopTime"` // key
}

