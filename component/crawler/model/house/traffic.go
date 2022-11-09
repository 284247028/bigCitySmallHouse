package house

type Traffic struct {
	Type     TrafficType `bson:"type" json:"type"`
	Line     string      `bson:"line" json:"line"`
	Station  string      `bson:"station" json:"station"`
	Distance int         `bson:"distance" json:"distance"`
}

const (
	TrafficTypeSubway = "subway"
	TrafficTypeBus    = "bus"
)

type TrafficType string

func (receiver *TrafficType) String() string {
	return string(*receiver)
}

func (receiver *TrafficType) ToCn() string {
	switch *receiver {
	case TrafficTypeSubway:
		return "地铁"
	case TrafficTypeBus:
		return "公交"
	}
	return receiver.String()
}
