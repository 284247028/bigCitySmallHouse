package house

const (
	TypeApartment = "apartment"
	TypeResidence = "residence"
)

type Type string

func (receiver *Type) String() string {
	return string(*receiver)
}
