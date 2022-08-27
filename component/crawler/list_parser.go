package crawler

type ListParser struct {
	*Parser
	Param *ListParam
}

func NewListParser(param *ListParam) *ListParser {
	checkParam(param)
	return &ListParser{
		Parser: NewParser(),
		Param:  param,
	}
}

func checkParam(param *ListParam) {
	if param.Page <= 0 {
		param.Page = 1
	}
}
