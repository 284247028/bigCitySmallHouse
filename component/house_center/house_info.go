package main

import (
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
	IdHex string `form:"id_hex"`
}

type HouseInfo struct {
	ImgUrls []string `json:"img_urls"`
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
	})
	var tHouse house.House
	singleResult := coll.MCollection().FindOne(context.TODO(), filter, opts)
	err = singleResult.Decode(&tHouse)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	houseIngo := HouseInfo{
		ImgUrls: tHouse.House.ImgUrls,
	}

	ctx.JSON(http.StatusOK, houseIngo)
}
