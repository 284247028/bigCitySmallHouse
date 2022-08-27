package house

const (
	SourceLeyoujia = "leyoujia"
)

type Source string

func (receiver *Source) String() string {
	return string(*receiver)
}
