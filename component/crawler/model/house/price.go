package house

type Price struct {
	Rent                 float64 `bson:"rent" json:"rent"`                                   // 租金
	ElectricityPerDegree float64 `bson:"electricity_per_degree" json:"electricityPerDegree"` // 电价/度
	WaterPerCube         float64 `bson:"water_per_cube" json:"waterPerCube"`                 // 水价/立方
	ManagementPerMeter   float64 `bson:"management_per_meter" json:"managementPerMeter"`     // 管理费/米
	ManagementTotal      float64 `bson:"management_total" json:"managementTotal"`            // 管理费，一次性
}
