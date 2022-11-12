package collections

import (
	"bigCitySmallHouse/component/user_center/model/user"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
)

type CollectionUser struct {
	*collection.Collection
}

func NewCollectionUser(opts *collection.Options) *CollectionHouse {
	if opts == nil {
		opts = &collection.Options{}
	}
	opts.DB = mongodb.DBUser
	opts.Collection = user.CollectionUser
	coll := collection.NewCollection(opts)
	return &CollectionHouse{
		Collection: coll,
	}
}
