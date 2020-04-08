package lexer_types

import (
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type ChainMap map[uip.ChainIDUnderlyingType]*FullAccount

type AccountMap map[string]ChainMap

func BuildAccountMap(accounts []FullAccount) (res AccountMap, err error) {
	res = make(AccountMap)
	var c ChainMap
	for i := range accounts {
		a := &accounts[i]

		if res[a.Name] == nil {
			res[a.Name] = make(ChainMap)
		}
		c = res[a.Name]

		if c[a.ChainID] != nil {
			return nil, errorn.NewAccountIndexConflict(a.Name, a.ChainID)
		}
		c[a.ChainID] = a

		if c[0] == nil {
			c[0] = a
		}
	}
	for _, c := range res {
		if len(c) > 2 {
			delete(c, 0)
		}
	}
	return
}
