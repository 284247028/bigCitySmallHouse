package house

type Location struct {
	Province string `bson:"province"`
	City     string `bson:"city"`
	Region   string `bson:"region"`
	Extra    string `bson:"extra"`
}
