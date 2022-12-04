package house

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const CollectionHouse = "house"

type House struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Source   Source             `bson:"source" json:"source"`       // 来源 大城小屋/乐有家/贝壳/...
	SourceId primitive.ObjectID `bson:"source_id" json:"source_id"` // 来源的唯一id 自己平台/第三方平台
	Type     Type               `bson:"type" json:"type"`           // 房子类型， 公寓/住宅/小区
	// 需要加个复式吗？
	Name        string      `bson:"name" json:"name"`               // 小区名字/公寓名字/...
	Title       string      `bson:"title" json:"title"`             // 标题
	Description string      `bson:"description" json:"description"` // 描述
	ImgUrls     []string    `bson:"img_urls" json:"img_urls"`       // 图片
	VideoUrls   []string    `bson:"video_urls" json:"video_urls"`   // 视频
	Area        float64     `bson:"area" json:"area"`               // 面积 单位/m²
	Price       Price       `bson:"price" json:"price"`             // 价格
	Floor       int         `bson:"floor" json:"floor"`             // 楼层
	Location    Location    `bson:"location" json:"location"`       // 地点
	RentType    RentType    `bson:"rent_type" json:"rent_type"`     // 租住类型, 合租/整租
	BuildTime   time.Time   `bson:"build_time" json:"build_time"`   // 建造日期
	Facilities  []Facility  `bson:"facilities" json:"facilities"`   // 设施   床、桌子、电梯、跑步机...
	Traffic     []Traffic   `bson:"traffic" json:"traffic"`         // 交通
	Composition Composition `bson:"composition" json:"composition"` // n厅n房...组成
	CreateAt    time.Time   `bson:"create_at" json:"create_at"`     // 文档创建时间
	UpdateAt    time.Time   `bson:"update_at" json:"update_at"`     // 数据更新时间
	DeleteAt    time.Time   `bson:"delete_at" json:"delete_at"`     // 删除时间
}
