package opintent

type BoolVariable = ResultI

type BaseIfStatementOpIntent struct {
	If        RawIntent `json:"if"`        // key
	Condition BoolVariable   `json:"condition"` // key
	Else      RawIntent `json:"else"`      // option
}

func (ier *Initializer) initIfStatement(info *RawIntent, content []byte) (sub interface{}, err error) {
	res, err := NewGJSONResult(content)
	if err != nil {
		return nil, err
	}

	var intent = new(BaseIfStatementOpIntent)
	sub = intent
	err = ier.InitContentR(&intent.If, res.Get("if"))
	if err != nil {
		return nil, err.(*ParseError).Desc(AtOpIntentField{"if"})
	}
	err = ier.InitContentR(&intent.Else, res.Get("else"))
	if err != nil {
		return nil, err.(*ParseError).Desc(AtOpIntentField{"else"})
	}
	intent.Condition = res.Get("condition")
	return
}
