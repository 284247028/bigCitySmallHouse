package model

type House struct {
	Source      int     // 来源 具体见常量定义
	HouseType   string  // 户型
	Area        float64 // 面积 单位/m²
	price       float64 // 价格
	Floor       int     // 楼层
	HasElevator bool    // 是否有电梯
}
