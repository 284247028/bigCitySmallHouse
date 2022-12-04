package main

import (
	house2 "bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/component/house_center/model/house"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type ParamHouseInfo struct {
	IdHex   string `form:"id_hex"`
	UserUId string `form:"user_uid"`
}

type HouseInfo struct {
	ImgUrls     []string          `json:"img_urls"`
	Name        string            `json:"name"`
	PriceRent   float64           `json:"price_rent"`
	Composition house.Composition `json:"composition"`
	Location    house.Location    `json:"location"`
	Traffic     []Traffic         `json:"traffic"`
	Facilities  []house.Facility  `json:"facilities"`
	IdHex       string            `json:"id_hex"`
	Total       int64             `json:"total"`
}

type Traffic struct {
	Station  string `json:"station"`
	Distance int    `json:"distance"`
}

func GetHouseInfo(ctx *gin.Context) {
	baseResponse := NewBaseResponse(ctx)
	var param ParamHouseInfo
	err := ctx.ShouldBind(&param)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	id, err := primitive.ObjectIDFromHex(param.IdHex)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	coll := collections.NewCollectionHouseCenter(nil)

	filter := bson.D{
		{"_id", id},
	}

	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{"house.img_urls", 1},
		{"house.name", 1},
		{"house.price", 1},
		{"house.composition", 1},
		{"house.location", 1},
		{"house.traffic", 1},
		{"house.facility", 1},
		{"house.uid", 1},
	})
	var tHouse house.House
	singleResult := coll.MCollection().FindOne(context.TODO(), filter, opts)
	err = singleResult.Decode(&tHouse)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	var traffic []Traffic
	for _, t := range tHouse.Traffic {
		tTraffic := Traffic{}
		if t.Type == house2.TrafficTypeBus {
			tTraffic.Station = "公交" + t.Line.String() + "路" + t.Station.String() + "站"
		}
		if t.Type == house2.TrafficTypeSubway {
			tTraffic.Station = "地铁" + t.Line.String() + "号线" + t.Station.String()
		}
		tTraffic.Distance = t.Distance
		traffic = append(traffic, tTraffic)
	}

	filter = bson.D{
		{"user_uid", param.UserUId},
		{"_id", tHouse.Id},
	}
	coll = collections.NewCollectionCollect(nil)
	total, err := coll.MCollection().CountDocuments(context.TODO(), filter)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	houseInfo := HouseInfo{
		ImgUrls:     tHouse.ImgUrls,
		Name:        tHouse.Name,
		PriceRent:   tHouse.Price.Rent,
		Composition: tHouse.Composition,
		Location:    tHouse.Location,
		Traffic:     traffic,
		Facilities:  tHouse.Facilities,
		IdHex:       tHouse.Id.Hex(),
		Total:       total,
	}

	ctx.JSON(http.StatusOK, houseInfo)
}
