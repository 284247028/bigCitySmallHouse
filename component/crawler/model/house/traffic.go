package house

type Traffic struct {
	Type     TrafficType `bson:"type" json:"type"`
	Line     int         `bson:"line" json:"line"`
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
