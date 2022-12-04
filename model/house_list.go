package model

import (
	"bigCitySmallHouse/component/house_center/model/house"
	"bigCitySmallHouse/component/publish/model"
	"strconv"
)

type HouseInfo struct {
	IdHex string // mongo唯一id十六进制
	house.House
	//RentType    house.RentType    // 整租/合租
	//Composition house.Composition // 2房2厅1卫
	//Price       house.Price       // ￥400/月
	//Location    house.Location    // 广州市白云区白云大道北1689号
	//ImgUrls     []string          // 图片地址
}

type HouseInfoList []HouseInfo

type HouseDisplay struct {
	IdHex       string `json:"id_hex"`
	RentType    string `json:"rent_type"`
	Composition string `json:"composition"`
	RentPrice   string `json:"rent_price"`
	Location    string `json:"location"`
	CoverUrl    string `json:"cover_url"`
}

type HouseDisplayList []HouseDisplay

func (receiver HouseInfoList) ToHouseDisplayList() HouseDisplayList {
	houseDisplayList := make(HouseDisplayList, 0, len(receiver))
	for _, houseInfo := range receiver {
		composition := strconv.Itoa(houseInfo.Composition.Room) + "房" + strconv.Itoa(houseInfo.Composition.Parlor) + "厅" + strconv.Itoa(houseInfo.Composition.Toilet) + "卫"
		rentPrice := "￥" + strconv.FormatFloat(houseInfo.Price.Rent, 'f', 0, 64) + "/月"
		location := ""
		if houseInfo.Location.Province != "" {
			location += houseInfo.Location.Province.String()
		}
		if houseInfo.Location.City != "" {
			location += houseInfo.Location.City.String()
		}
		if houseInfo.Location.Region != "" {
			location += houseInfo.Location.Region.String()
		}
		if houseInfo.Location.Extra != "" {
			location += houseInfo.Location.Extra
		}
		imgUrl := ""
		if len(houseInfo.ImgUrls) > 0 {
			imgUrl = houseInfo.ImgUrls[0]
		}
		houseDisplayList = append(houseDisplayList, HouseDisplay{
			IdHex:       houseInfo.IdHex,
			RentType:    houseInfo.RentType.ToCn(),
			Composition: composition,
			RentPrice:   rentPrice,
			Location:    location,
			CoverUrl:    imgUrl,
		})
	}
	return houseDisplayList
}

func HouseCenter2HouseInfoList(tHouses []house.House) HouseInfoList {
	houseInfoList := make(HouseInfoList, 0, len(tHouses))
	for _, tHouse := range tHouses {
		houseInfoList = append(houseInfoList, HouseInfo{
			IdHex: tHouse.Id.Hex(),
			House: tHouse,
		})
	}
	return houseInfoList
}

func Publishes2HouseList(publishes []model.Publish) HouseInfoList {
	houseInfoList := make(HouseInfoList, 0, len(publishes))
	for _, publish := range publishes {
		houseInfoList = append(houseInfoList, HouseInfo{
			IdHex: publish.Id.Hex(),
			House: house.House{
				RentType:    publish.RentType,
				Composition: publish.Composition,
				Price:       publish.Price,
				Location:    publish.Location,
				ImgUrls:     publish.ImgUrls,
			},
		})

	}
	return houseInfoList
}
