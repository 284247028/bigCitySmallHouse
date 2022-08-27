package crawler

type SingleParser struct {
	*Parser
	Param *SingleParam
}

func NewSingleParser(param *SingleParam) *SingleParser {
	return &SingleParser{
		Parser: NewParser(),
		Param:  param,
	}
}
