package lexer

type RootIntents struct {
	Infos   []Intent
	NameMap map[string]int
}

func (r RootIntents) Len() int {
	return len(r.Infos)
}

func (r RootIntents) GetRawIntent(idx int) Intent {
	return r.Infos[idx]
}

