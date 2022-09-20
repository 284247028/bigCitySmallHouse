package house

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const CollectionName = "house"

type House struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"` // mongo id
	UId         string             `bson:"uid"`           // 怕不同房源source_id重复导致数据覆盖, source + _ + source_id
	SourceId    string             `bson:"source_id"`     // 来源的房子id
	Source      Source             `bson:"source"`        // 来源
	Type        Type               `bson:"type"`          // 房子类型， 公寓/住宅
	Name        string             `bson:"name"`          // 小区名字/公寓名字/...
	ImgUrls     []string           `bson:"img_urls"`      // 图片
	VideoUrls   []string           `bson:"video_urls"`    // 视频
	Area        float64            `bson:"area"`          // 面积 单位/m²
	Price       float64            `bson:"price"`         // 价格
	Floor       int                `bson:"floor"`         // 楼层
	Elevator    bool               `bson:"elevator"`      // 是否有电梯
	Location    Location           `bson:"location"`      // 地点
	BuildTime   time.Time          `bson:"build_time"`    // 建造日期
	Furniture   []string           `bson:"furniture"`     // 家具
	Facility    []string           `bson:"facility"`      // 设施
	Traffic     []Traffic          `bson:"traffic"`       // 交通
	Composition Composition        `bson:"composition"`   // n厅n房...组成
	UpdateAt    time.Time          `bson:"update_at"`     // 数据更新时间
}

const (
	TypeApartment = "apartment" // 公寓
	TypeResidence = "residence" // 住宅
	TypeVilla     = "villa"     // 别墅
)

type Type string

func (receiver *Type) String() string {
	return string(*receiver)
}
