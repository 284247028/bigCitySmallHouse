package house

type Composition struct {
	Parlor int `bson:"parlor"`
	Room   int `bson:"room"`
	Toilet int `bson:"toilet"`
}
