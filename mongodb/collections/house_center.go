package collections

import (
	"bigCitySmallHouse/component/house_center/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
)

type CollectionHouse struct {
	*collection.Collection
}

func NewCollectionHouseCenter(opts *collection.Options) *CollectionHouse {
	if opts == nil {
		opts = &collection.Options{}
	}
	opts.DB = mongodb.DBHouseCenter
	opts.Collection = house.CollectionHouse
	coll := collection.NewCollection(opts)
	return &CollectionHouse{
		Collection: coll,
	}
}
