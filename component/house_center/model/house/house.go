package house

import (
	house2 "bigCitySmallHouse/component/crawler/model/house"
	"time"
)

const CollectionHouse = "house"

type House struct {
	House    house2.House
	UpdateAt time.Time
	Shelve   bool // 上架
}
