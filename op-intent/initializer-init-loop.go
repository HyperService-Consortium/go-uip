package opintent

type BaseLoopStatementOpIntent struct {
	Loop  RawIntent `json:"loop"`     // key
	Times int64     `json:"loopTime"` // key
}

func (ier *Initializer) initLoopStatement(info *RawIntent, content []byte) (sub interface{}, err error) {
	res, err := NewGJSONResult(content)
	if err != nil {
		return nil, err
	}

	var intent = new(BaseLoopStatementOpIntent)
	sub = intent
	err = ier.InitContentR(&intent.Loop, res.Get("loop"))
	if err != nil {
		return nil, err.(*ParseError).Desc(AtOpIntentField{"loop"})
	}
	intent.Times = res.Get("loopTime").Int()
	return
}
