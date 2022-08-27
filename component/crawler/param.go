package crawler

type ListParam struct {
	Page int
}

func NewListParam() *ListParam {
	return &ListParam{Page: 1}
}

type SingleParam struct {
	Id string
}

func NewSingleParam() *SingleParam {
	return &SingleParam{}
}
