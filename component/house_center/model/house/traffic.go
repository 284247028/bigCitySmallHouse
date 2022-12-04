package house

type Traffic struct {
	Type     TrafficType `bson:"type" json:"type"`
	Line     Line        `bson:"line" json:"line"`
	Station  Station     `bson:"station" json:"station"`
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

type Line string

func (receiver *Line) String() string {
	return string(*receiver)
}

type Station string

func (receiver *Station) String() string {
	return string(*receiver)
}
