package house

type Composition struct {
	Parlor  int `bson:"parlor" json:"parlor"` // 客厅
	Room    int `bson:"room" json:"room"`
	Toilet  int `bson:"toilet" json:"toilet"`
	Kitchen int `bson:"kitchen" json:"kitchen"`
	Balcony int `bson:"balcony" json:"balcony"` // 阳台
}
