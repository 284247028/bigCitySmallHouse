package collection

import (
	"bigCitySmallHouse/mongodb"
	"context"
	"fmt"
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

func (receiver *Collection) UpsertOne(filter, object interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	opts.SetUpsert(true)
	update := bson.D{
		{"$set", object},
	}
	return receiver.mCollection.UpdateOne(context.TODO(), filter, update, opts)
}

func (receiver *Collection) Pagination(page, size int, ctx context.Context, filter interface{},
	opts *options.FindOptions) (*mongo.Cursor, error) {
	if page <= 0 {
		return nil, fmt.Errorf("分页的page小于等于0： %d", page)
	}
	if size <= 0 {
		return nil, fmt.Errorf("分页的size小于等于0：%d", size)
	}
	skipNum := (page - 1) * size
	limitNum := page * size

	opts.SetLimit(int64(limitNum))
	opts.SetSkip(int64(skipNum))
	return receiver.MCollection().Find(ctx, filter, opts)
}
