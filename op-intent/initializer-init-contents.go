package opintent

import "github.com/HyperService-Consortium/go-uip/const/trans_type"

type RawIntent struct {
	BaseOpIntent
	Sub    interface{}
	OpType trans_type.Type
}

func (r *RawIntent) GetType() trans_type.Type {
	return r.OpType
}

func (r *RawIntent) GetSub() interface{} {
	return r.Sub
}

type RawIntents struct {
	infos   []RawIntent
	nameMap map[string]int
}

func (r RawIntents) Len() int {
	return len(r.infos)
}

func (r RawIntents) GetRawIntent(idx int) RawIntentI {
	return &r.infos[idx]
}

func (ier *Initializer) InitContents(contents [][]byte) (intents *RawIntents, err error) {
	intents = &RawIntents{
		infos:   make([]RawIntent, len(contents)),
		nameMap: make(map[string]int),
	}

	for idx, content := range contents {
		err := ier.InitContent(&intents.infos[idx], content)
		if err != nil {
			return nil, err.(*ParseError).Desc(AtOpIntentsPos{Pos: idx})
		}
		intents.nameMap[intents.infos[idx].Name] = idx
		//infos[idx].OpTypeString
	}
	return
}

func (ier *Initializer) InitContent(intent *RawIntent, content []byte) (err error) {
	err = ier.unmarshal(content, &intent.BaseOpIntent)
	if err != nil {
		return err
	}
	//infos[idx].OpTypeString

	switch intent.OpTypeString {
	case "Payment":
		intent.OpType = trans_type.Payment
		if intent.Sub, err = ier.initPayment(intent, content); err != nil {
			return err
		}

	case "ContractInvocation":
		intent.OpType = trans_type.ContractInvoke
		if intent.Sub, err = ier.initContractInvocation(intent, content); err != nil {
			return err
		}

	case "IfStatement":
		intent.OpType = trans_type.IfStatement
		if intent.Sub, err = ier.initIfStatement(intent, content); err != nil {
			return err
		}

	case "loopFunction":
		intent.OpType = trans_type.LoopStatement
		if intent.Sub, err = ier.initLoopStatement(intent, content); err != nil {
			return err
		}

	default:
		return newInvalidFieldError(invalidOpType)
	}
	return nil
}

func (ier *Initializer) InitContentsR(source ResultI) (intents *RawIntents, err error) {
	rawContents := source.Get(FieldOpIntents)
	if ! rawContents.Exists() {
		return nil, newFieldNotFound(FieldOpIntents)
	}
	contents := rawContents.Array()

	intents = &RawIntents{
		infos:   make([]RawIntent, contents.Len()),
		nameMap: make(map[string]int),
	}

	for idx := 0; idx < contents.Len(); idx++ {
		err := ier.InitContentR(&intents.infos[idx], contents.Index(idx))
		if err != nil {
			return nil, err.(*ParseError).Desc(AtOpIntentsPos{Pos: idx})
		}
		intents.nameMap[intents.infos[idx].Name] = idx
		//infos[idx].OpTypeString
	}
	return
}

func (ier *Initializer) InitContentR(intent *RawIntent, content ResultI) (err error) {
	name, opType := content.Get(FieldOpIntentsName), content.Get(FieldOpIntentsOpType)
	if !name.Exists() {
		return newFieldNotFound(FieldOpIntentsName)
	}
	if !opType.Exists() {
		return newFieldNotFound(FieldOpIntentsOpType)
	}
	intent.Name = name.String()
	intent.OpTypeString = opType.String()

	switch intent.OpTypeString {
	case "Payment":
		intent.OpType = trans_type.Payment
		if intent.Sub, err = ier.initPaymentR(intent, content); err != nil {
			return err
		}

	case "ContractInvocation":
		intent.OpType = trans_type.ContractInvoke
		if intent.Sub, err = ier.initContractInvocationR(intent, content); err != nil {
			return err
		}

	//case "IfStatement":
	//	intent.OpType = trans_type.IfStatement
	//	if intent.Sub, err = ier.initIfStatement(intent, content); err != nil {
	//		return err
	//	}

	default:
		return newInvalidFieldError(invalidOpType)
	}
	return nil
}
