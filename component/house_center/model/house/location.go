package house

type Location struct {
	Province Province `bson:"province" json:"province"`
	City     City     `bson:"city" json:"city"`
	Region   Region   `bson:"region" json:"region"`
	Extra    string   `bson:"extra" json:"extra"`
}

type Province string

func (receiver *Province) String() string {
	return string(*receiver)
}

const (
	CityGuangZhou Province = "广州市"
)

type City string

func (receiver *City) String() string {
	return string(*receiver)
}

const (
	ProvinceGuangDong City = "广东省"
)

type Region string

func (receiver *Region) String() string {
	return string(*receiver)
}

const (
	RegionYueXiu    Region = "越秀区"
	RegionHuangPu   Region = "黄埔区"
	RegionHaiZhu    Region = "海珠区"
	RegionTianHe    Region = "天河区"
	RegionLiWan     Region = "荔湾区"
	RegionPanYu     Region = "番禺区"
	RegionZengCheng Region = "增城区"
	RegionCongHua   Region = "从化区"
	RegionHuaDu     Region = "花都区"
	RegionNanSha    Region = "南沙区"
)
