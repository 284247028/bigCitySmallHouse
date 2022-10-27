package house

import "bigCitySmallHouse/mongodb/collection"

type Collection struct {
	*collection.Collection
}

func NewCollection(opts *collection.Options) *Collection {
	collection.NewCollection(opts)
	return &Collection{}
}
