package house

import "time"

type House struct {
	UId            string    // 唯一id
	Id             string    // 来源的房子id
	Source         Source    // 来源
	Type           Type      // 房子类型， 公寓/住宅
	Name           string    // 小区名字/公寓名字/...
	ImgUrls        []string  // 图片
	Area           float64   // 面积 单位/m²
	price          float64   // 价格
	Floor          int       // 楼层
	Elevator       bool      // 是否有电梯
	Location       Location  // 地点
	BuildDate      time.Time // 建造日期
	Furniture      []string  // 配套设施
	SubwayStation  string    // 地铁站	todo 定义新类型
	SubwayDistance int       // 地铁距离
	BusStation     string    // 公交站 		todo 定义新类型
	BusDistance    int       // 公交站距离
	Parlor         int       // 厅
	Room           int       // 房间
}
