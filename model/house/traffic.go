package house

type Traffic struct {
	Type     TrafficType
	Line     string
	Station  string
	Distance int
}

const (
	TrafficTypeSubway = "subway"
	TrafficTypeBus    = "bus"
)

type TrafficType string

func (receiver *TrafficType) String() string {
	return string(*receiver)
}
