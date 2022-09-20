package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var _once sync.Once
var _mongoClient *mongo.Client

const LocalUri = "mongodb://localhost:27017"

const DBCrawler = "crawler"

type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

type Options struct {
	Uri string
}

func NewOptions() *Options {
	return &Options{
		Uri: LocalUri,
	}
}

func (receiver *DB) ConnectMongodb(opts *Options) error {
	var err error
	clientOptions := options.Client()
	clientOptions.ApplyURI(opts.Uri)
	_once.Do(func() {
		_mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	})
	return err
}

func GetClient() *mongo.Client {
	return _mongoClient
}
