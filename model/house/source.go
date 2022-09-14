package house

const (
	SourceLeyoujia = "leyoujia"
	SourceAnjuke   = "anjuke"
)

type Source string

func (receiver *Source) String() string {
	return string(*receiver)
}
