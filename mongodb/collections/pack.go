package collections

import (
	"bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionPack struct {
	*collection.Collection
}

func NewCollectionPack(opts *collection.Options) *CollectionPack {
	if opts == nil {
		opts = &collection.Options{}
	}
	opts.DB = mongodb.DBCrawler
	opts.Collection = house.CollectionPack
	coll := collection.NewCollection(opts)
	return &CollectionPack{
		Collection: coll,
	}
}

func (receiver *CollectionPack) Houses2Packs(houses []house.House) []house.Pack {
	packs := make([]house.Pack, 0, len(houses))
	for _, tHouse := range houses {
		pack := house.Pack{
			Status: house.PackStatusList,
			House:  tHouse,
		}
		packs = append(packs, pack)
	}
	return packs
}

func (receiver *CollectionPack) PackUpsertMany(packs []house.Pack) ([]*mongo.UpdateResult, error) {
	var results []*mongo.UpdateResult
	for _, pack := range packs {
		uid := pack.House.Source.String() + "-" + pack.House.SourceId
		pack.House.UId = uid
		filter := bson.D{
			{"house.uid", uid},
		}
		opts := &options.UpdateOptions{}
		result, err := receiver.UpsertOne(filter, pack, opts)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
