package collections

import (
	"bigCitySmallHouse/component/publish/model"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CollectionPublish struct {
	*collection.Collection
}

func NewCollectionPublish(opts *collection.Options) *CollectionPublish {
	if opts == nil {
		opts = &collection.Options{}
	}
	opts.DB = mongodb.DBPublish
	opts.Collection = model.CollectionPublish
	coll := collection.NewCollection(opts)
	return &CollectionPublish{
		Collection: coll,
	}
}

func (receiver *CollectionPublish) UpsertOnePublish(filter bson.D, publish model.Publish, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	publish.UpdateAt = time.Now()
	if len(filter) == 0 {
		publish.CreateAt = time.Now()
	}
	return receiver.Collection.UpsertOne(filter, publish, opts)
}
