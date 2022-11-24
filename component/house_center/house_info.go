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
	ImgUrls     []string           `json:"img_urls"`
	Name        string             `json:"name"`
	PriceRent   float64            `json:"price_rent"`
	Composition house2.Composition `json:"composition"`
	Location    house2.Location    `json:"location"`
	Traffic     []Traffic          `json:"traffic"`
	Facilities  []string           `json:"facilities"`
	HouseUId    string             `json:"house_uid"`
	Total       int64              `json:"total"`
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
	for _, t := range tHouse.House.Traffic {
		tTraffic := Traffic{}
		if t.Type == house2.TrafficTypeBus {
			tTraffic.Station = "公交" + t.Line + "路" + t.Station + "站"
		}
		if t.Type == house2.TrafficTypeSubway {
			tTraffic.Station = "地铁" + t.Line + "号线" + t.Station
		}
		tTraffic.Distance = t.Distance
		traffic = append(traffic, tTraffic)
	}

	filter = bson.D{
		{"user_uid", param.UserUId},
		{"house_uid", tHouse.House.UId},
	}
	coll = collections.NewCollectionCollect(nil)
	total, err := coll.MCollection().CountDocuments(context.TODO(), filter)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	houseIngo := HouseInfo{
		ImgUrls:     tHouse.House.ImgUrls,
		Name:        tHouse.House.Name,
		PriceRent:   tHouse.House.Price.Rent,
		Composition: tHouse.House.Composition,
		Location:    tHouse.House.Location,
		Traffic:     traffic,
		Facilities:  tHouse.House.Facilities,
		HouseUId:    tHouse.House.UId,
		Total:       total,
	}

	ctx.JSON(http.StatusOK, houseIngo)
}
