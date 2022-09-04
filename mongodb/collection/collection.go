package collection

import (
	"bigCitySmallHouse/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
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
