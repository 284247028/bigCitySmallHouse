package house

import (
	house2 "bigCitySmallHouse/component/crawler/model/house"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const CollectionHouse = "house"

type House struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	House    house2.House       `bson:"house" json:"house"`
	UpdateAt time.Time          `bson:"update_at" json:"update_at"`
	Shelve   bool               `bson:"shelve" json:"shelve"` // 上架
}
