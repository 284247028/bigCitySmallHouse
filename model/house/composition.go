package house

type Composition struct {
	Parlor  int `bson:"parlor"` // 客厅
	Room    int `bson:"room"`
	Toilet  int `bson:"toilet"`
	Kitchen int `bson:"kitchen"`
	Balcony int `bson:"balcony"` // 阳台
}
