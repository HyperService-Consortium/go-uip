package lexer

import (
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/standard"
)

type DocumentLexer struct {
	BaseLexer
}

func (l *DocumentLexer) InitContents(source document.Document) (r *RootIntents, err error) {
	rawContents := source.Get(FieldOpIntents)
	if !rawContents.Exists() {
		return nil, errorn.NewFieldNotFound(FieldOpIntents)
	}

	return l.InitContents_(rawContents)
}

func (l *DocumentLexer) InitContents_(source document.Document) (r *RootIntents, err error) {
	contents := source.Array()
	r = &RootIntents{}

	r.Infos = make([]Intent, contents.Len())
	r.NameMap = make(map[string]int)

	for idx := 0; idx < contents.Len(); idx++ {
		r.Infos[idx], err = l.InitContent(contents.Index(idx))
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentsPos{Pos: idx})
		}
		r.NameMap[r.Infos[idx].GetName()] = idx
		//infos[idx].OpTypeString
	}
	return
}

func (l *DocumentLexer) InitContent(content document.Document) (i Intent, err error) {
	intent := &IntentImpl{}
	err = intent.UnmarshalDocument(content)
	if err != nil {
		return nil, err
	}

	switch intent.OpType {
	case token_type.Pay:
		return l.initPayment(intent, content)
	case token_type.Invoke:
		return l.initContractInvocation(intent, content)
	case token_type.If:
		return l.initIfStatement(intent, content)
	case token_type.Loop:
		return l.initLoopStatement(intent, content)
	default:
		return nil, errorn.NewInvalidFieldError(errorn.InvalidOpType)
	}
}

func (l *DocumentLexer) initContractInvocation(info *IntentImpl, content document.Document) (sub Intent, err error) {
	var intent = &InvokeIntent{IntentImpl: info}
	sub = intent
	err = intent.UnmarshalDocument(content)
	if err != nil {
		return nil, err
	}
	return l.checkContractInvocation(intent)
}

func (l *DocumentLexer) initPayment(info *IntentImpl, content document.Document) (sub Intent, err error) {
	var intent = &PaymentIntent{IntentImpl: info}
	sub = intent
	err = intent.UnmarshalResult(content)
	if err != nil {
		return nil, err
	}

	if intent.Src == nil {
		return nil, errorn.NewFieldNotFound("src")
	}
	if intent.Dst == nil {
		return nil, errorn.NewFieldNotFound("dst")
	}

	// cinfo, err = SearchChainId(paymentIntent.Src.ChainId)
	// if err != nil {
	// 	return nil, err
	// }
	// var processor ProcessorInterface
	// switch cinfo.GetChainType() {
	// case chain_type.Ethereum:
	// 	processor = eth_processor.GetProcessor()
	// default:
	// 	return nil, errors.New("unsupport chain_type")
	// }
	// if !processor.CheckAddress(paymentIntent.Src.)

	if err = standard.CheckValidHexString(intent.Amount); err != nil {
		return nil, errorn.NewInvalidFieldError(err).Desc(errorn.AtOpIntentField{Field: "amount"})
	}
	if t, ok := UnitType.Mapping[intent.UnitString]; !ok {
		return nil, errorn.NewInvalidFieldError(errorn.UnknownUnitType).Desc(errorn.AtOpIntentField{Field: "unit"})
	} else {
		intent.Unit = t
	}

	return
}

func (l *DocumentLexer) initIfStatement(info *IntentImpl, content document.Document) (sub Intent, err error) {

	var intent = &IfIntent{IntentImpl: info}
	sub = intent
	intent.If, err = l.InitContents_(content.Get("if"))
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "if"})
	}
	intent.Else, err = l.InitContents_(content.Get("else"))
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "else"})
	}
	intent.Condition, err = ParamUnmarshalResult(content.Get("condition"))
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "condition"})
	}
	return
}

func (l *DocumentLexer) initLoopStatement(info *IntentImpl, content document.Document) (sub Intent, err error) {

	var intent = &LoopIntent{IntentImpl: info}
	intent.Loop, err = l.InitContents_(content.Get("loop"))
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "loop"})
	}
	intent.Times = content.Get("loopTime").Int()
	return intent, nil
}

type RawDependencies struct {
	dependencies []RawDependency
}

func (r RawDependencies) Len() int {
	return len(r.dependencies)
}

func (r RawDependencies) GetDependencies(i int) RawDependencyI {
	return &r.dependencies[i]
}

type RawDependenciesI interface {
	Len() int
	GetDependencies(i int) RawDependencyI
}

type RawDependencyI interface {
	GetSrc() string
	GetDst() string
}

func (l *DocumentLexer) InitDependencies(source document.Document) (deps *RawDependencies, err error) {
	if source.Exists() && !source.IsArray() {
		return nil, errorn.NewInvalidFieldError(errorn.ErrTypeError).Desc(errorn.AtOpIntentField{Field: FieldOpIntentsParameters})
	}
	rawDeps := source.Array()

	deps = &RawDependencies{
		dependencies: make([]RawDependency, rawDeps.Len()),
	}
	for idx := 0; idx < rawDeps.Len(); idx++ {
		err = deps.dependencies[idx].UnmarshalResult(rawDeps.Index(idx))
		if err != nil {
			return
		}
	}
	return
}

func (l *DocumentLexer) initAccounts(nameKey string, source document.Document) (accounts []lexer_types.FullAccount, err error) {
	if source.Exists() && !source.IsArray() {
		return nil, errorn.NewInvalidFieldError(errorn.ErrTypeError).Desc(errorn.AtOpIntentField{Field: FieldOpIntentsParameters})
	}
	rawAccounts := source.Array()

	accounts = make([]lexer_types.FullAccount, rawAccounts.Len())
	for idx := 0; idx < rawAccounts.Len(); idx++ {
		err = accounts[idx].UnmarshalResult(nameKey, rawAccounts.Index(idx))
		if err != nil {
			return
		}
	}
	return
}

func (l *DocumentLexer) InitContracts_(source document.Document) (accounts []lexer_types.FullAccount, err error) {
	return l.initAccounts("contractName", source)
}

func (l *DocumentLexer) InitAccounts_(source document.Document) (accounts []lexer_types.FullAccount, err error) {
	return l.initAccounts("userName", source)
}

func (l *DocumentLexer) InitContracts(source document.Document) (lexer_types.AccountMap, error) {
	r, err := l.InitContracts_(source)
	if err != nil {
		return nil, err
	}
	return lexer_types.BuildAccountMap(r)
}

func (l *DocumentLexer) InitAccounts(source document.Document) (lexer_types.AccountMap, error) {
	r, err := l.InitAccounts_(source)
	if err != nil {
		return nil, err
	}
	return lexer_types.BuildAccountMap(r)
}
