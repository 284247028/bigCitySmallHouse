package collections

import (
	"bigCitySmallHouse/component/user_center/model/collect"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
)

type CollectionCollect struct {
	*collection.Collection
}

func NewCollectionCollect(opts *collection.Options) *CollectionHouse {
	if opts == nil {
		opts = &collection.Options{}
	}
	opts.DB = mongodb.DBUser
	opts.Collection = collect.CollectionCollect
	coll := collection.NewCollection(opts)
	return &CollectionHouse{
		Collection: coll,
	}
}
