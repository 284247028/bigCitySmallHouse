package house

type Traffic struct {
	Type     TrafficType `bson:"type"`
	Line     int         `bson:"line"`
	Station  string      `bson:"station"`
	Distance int         `bson:"distance"`
}

const (
	TrafficTypeSubway = "subway"
	TrafficTypeBus    = "bus"
)

type TrafficType string

func (receiver *TrafficType) String() string {
	return string(*receiver)
}
