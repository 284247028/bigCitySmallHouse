package collection

import (
	"bigCitySmallHouse/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	DB         string
	Collection string
}

func NewOptions() *Options {
	return &Options{}
}

type Collection struct {
	mCollection *mongo.Collection
}

func NewCollection(opts *Options) *Collection {
	mCollection := mongodb.GetClient().Database(opts.DB).Collection(opts.Collection)
	return &Collection{
		mCollection: mCollection,
	}
}

func (receiver *Collection) MCollection() *mongo.Collection {
	return receiver.mCollection
}

func (receiver *Collection) UpsertMany(filter, object interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	update := bson.D{
		{"$set", object},
	}
	opts.SetUpsert(true)
	return receiver.mCollection.UpdateOne(context.TODO(), filter, update, opts)
}
