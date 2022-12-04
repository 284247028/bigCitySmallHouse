package house

type Location struct {
	Province string `bson:"province" json:"province"`
	City     string `bson:"city" json:"city"`
	Region   string `bson:"region" json:"region"`
	Extra    string `bson:"extra" json:"extra"`
}

const (
	CityGuangZhou = "广州市"
)

const (
	ProvinceGuangDong = "广东省"
)

const (
	RegionYueXiu    = "越秀区"
	RegionHuangPu   = "黄埔区"
	RegionHaiZhu    = "海珠区"
	RegionTianHe    = "天河区"
	RegionLiWan     = "荔湾区"
	RegionPanYu     = "番禺区"
	RegionZengCheng = "增城区"
	RegionCongHua   = "从化区"
	RegionHuaDu     = "花都区"
	RegionNanSha    = "南沙区"
)
