package house

import "time"

type House struct {
	UId         string      // 唯一id
	Id          string      // 来源的房子id
	Source      Source      // 来源
	Type        Type        // 房子类型， 公寓/住宅
	Name        string      // 小区名字/公寓名字/...
	ImgUrls     []string    // 图片
	VideoUrls   []string    // 视频
	Area        float64     // 面积 单位/m²
	Price       float64     // 价格
	Floor       int         // 楼层
	Elevator    bool        // 是否有电梯
	Location    Location    // 地点
	BuildTime   time.Time   // 建造日期
	Furniture   []string    // 家具
	Facility    []string    // 设施
	Traffic     []Traffic   // 交通
	Composition Composition // n厅n房...组成
}

const (
	TypeApartment = "apartment"
	TypeResidence = "residence"
)

type Type string

func (receiver *Type) String() string {
	return string(*receiver)
}
