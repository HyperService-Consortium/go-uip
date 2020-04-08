package lexer

import "github.com/HyperService-Consortium/go-uip/internal/lexer_types"

const (
	FieldOpIntents     = "op-intents"
	FieldOpIntentsName = "name"

	FieldOpIntentsInvoker      = "invoker"
	FieldOpIntentsContract     = "contract"
	FieldOpIntentsContractAddr = "contract_addr"
	FieldOpIntentsContractCode = "contract_code"
	FieldOpIntentsFunc         = "func"
	FieldOpIntentsParameters   = "parameters"
	FieldOpIntentsAmount       = "amount"
	FieldOpIntentsMeta         = "meta"

	FieldOpIntentsSrc  = "src"
	FieldOpIntentsDst  = "dst"
	FieldOpIntentsUnit = "unit"

	FieldOpIntentsDependencies = "dependencies"

	FieldOpIntentsDomain   = lexer_types.FieldOpIntentsDomain
	FieldOpIntentsUserName = "user_name"
	FieldKeyType           = "type"
	FieldOpIntentsSign     = "sign"
	FieldOpIntentsValue    = "value"

	FieldKeyLeft       = "left"
	FieldKeyRight      = "right"
	FieldDependencyDep = "dep"

	FieldContractPos     = "pos"
	FieldContractField   = "field"
	FieldContractAccount = "contract"

	FieldValueConstant = "constant"

	FieldKeyAddress = lexer_types.FieldKeyAddress
)
