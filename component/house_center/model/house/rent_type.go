package house

type RentType string

const (
	RentTypeEntire RentType = "entire" // 整租
	RentTypeShared RentType = "shared" // 合租
)

func (receiver *RentType) String() string {
	return string(*receiver)
}

func (receiver *RentType) ToCn() string {
	switch *receiver {
	case RentTypeEntire:
		return "整租"
	case RentTypeShared:
		return "合租"
	}
	return "[[" + receiver.String() + "]]"
}

func (receiver *RentType) FormatRentType() {
	switch *receiver {
	case "整租":
		*receiver = RentTypeEntire
	case "合租":
		*receiver = RentTypeShared
	}
}
