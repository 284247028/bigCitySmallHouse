package collections

import (
	"bigCitySmallHouse/component/cdn/model"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
)

type CollectionCdnImage struct {
	*collection.Collection
}

func NewCollectionCdnImage(opts *collection.Options) *CollectionHouse {
	if opts == nil {
		opts = &collection.Options{}
	}
	opts.DB = mongodb.DBCdn
	opts.Collection = model.CollectionCdnImage
	coll := collection.NewCollection(opts)
	return &CollectionHouse{
		Collection: coll,
	}
}
