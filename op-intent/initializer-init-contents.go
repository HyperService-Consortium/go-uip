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
		if intent.Sub, err = ier.initPayment(intent, content); err != nil {
			return err
		}

	case "ContractInvocation":
		if intent.Sub, err = ier.initContractInvocation(intent, content); err != nil {
			return err
		}

	default:
		return newInvalidFieldError(invalidOpType)
	}
	return nil
}
