package collections

import (
	"bigCitySmallHouse/model/house"
	"bigCitySmallHouse/mongodb/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionHouse struct {
	*collection.Collection
}

func NewCollectionHouse(opts *collection.Options) *CollectionHouse {
	coll := collection.NewCollection(opts)
	return &CollectionHouse{
		Collection: coll,
	}
}

func (receiver *CollectionHouse) HouseUpsertMany(houses []house.House) ([]*mongo.UpdateResult, error) {
	var results []*mongo.UpdateResult
	for _, tHouse := range houses {
		uid := tHouse.Source.String() + "-" + tHouse.SourceId
		tHouse.UId = uid
		filter := bson.D{
			{"uid", uid},
		}
		opts := &options.UpdateOptions{}
		result, err := receiver.UpsertOne(filter, tHouse, opts)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
