package house

import (
	house2 "bigCitySmallHouse/component/crawler/model/house"
	"time"
)

const CollectionHouse = "house"

type House struct {
	House    house2.House `bson:"house" json:"house"`
	UpdateAt time.Time    `bson:"update_at" json:"update_at"`
	Shelve   bool         `bson:"shelve" json:"shelve"` // 上架
}
