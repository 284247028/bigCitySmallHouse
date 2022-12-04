package house

type Type string

const (
	TypeUnknown   Type = "unknown"   // 未知
	TypeApartment Type = "apartment" // 公寓
	TypeCommunity Type = "community" // 小区
	TypeResidence Type = "residence" // 住宅
	TypeVilla     Type = "villa"     // 别墅
	TypeShop      Type = "shop"      // 商铺
	TypeParking   Type = "parking"   // 停车位
	TypeOffice    Type = "office"    // 办公楼
)

func (receiver *Type) FormatType() {
	switch *receiver {
	case "公寓":
		*receiver = TypeApartment
	case "住宅":
		*receiver = TypeResidence
	case "小区":
		*receiver = TypeCommunity
	}
}
