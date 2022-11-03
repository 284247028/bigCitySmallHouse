package main

import (
	"bigCitySmallHouse/component/house_center/model/house"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func ReadHouse(ctx *gin.Context) {
	baseResponse := NewBaseResponse(ctx)
	coll := collections.NewCollectionHouseCenter(nil)
	page := 1
	size := 20
	filter := bson.D{}
	cursor, err := coll.Pagination(page, size, context.TODO(), filter, options.Find())
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	var tHouses []house.House
	err = cursor.All(context.TODO(), &tHouses)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, tHouses)
}
