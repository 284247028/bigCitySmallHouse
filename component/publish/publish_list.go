package main

import (
	"bigCitySmallHouse/component/base/base_action"
	"bigCitySmallHouse/component/publish/model"
	model2 "bigCitySmallHouse/model"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type ParamPublishList struct {
	UserUId string `form:"user_uid"`
	Status  string `form:"status"`
}

func PublishList(ctx *gin.Context) {
	var param ParamPublishList
	err := ctx.ShouldBind(&param)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	coll := collections.NewCollectionPublish(nil)

	filter := bson.D{}
	if param.UserUId != "" {
		filter = append(filter, bson.E{Key: "user_uid", Value: param.UserUId})
	}
	if param.Status != "" {
		filter = append(filter, bson.E{Key: "status", Value: param.Status})
	}

	cursor, err := coll.MCollection().Find(context.TODO(), filter)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	var publishes []model.Publish
	err = cursor.All(context.TODO(), &publishes)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, model2.Publishes2HouseList(publishes).ToHouseDisplayList())
}
