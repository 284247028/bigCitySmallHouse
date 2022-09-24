package house

type Location struct {
	Province string `bson:"province"`
	City     string `bson:"city"`
	Region   string `bson:"region"`
	Extra    string `bson:"extra"`
}

const (
	CityGuangZhou = "广州"
)

const (
	ProvinceGuangDong = "广东"
)

const (
	RegionYueXiu    = "越秀"
	RegionHuangPu   = "黄埔"
	RegionHaiZhu    = "海珠"
	RegionTianHe    = "天河"
	RegionLiWan     = "荔湾"
	RegionPanYu     = "番禺"
	RegionZengCheng = "增城"
	RegionCongHua   = "从化"
	RegionHuaDu     = "花都"
	RegionNanSha    = "南沙"
)
