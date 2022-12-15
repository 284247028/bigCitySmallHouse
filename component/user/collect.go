package main

import (
	"bigCitySmallHouse/component/base/base_action"
	"bigCitySmallHouse/component/user/model/collect"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type ParamCollect struct {
	UserUId  string `json:"user_uid"`
	HouseUId string `json:"house_uid"`
	Collect  bool   `json:"collect"`
}

func Collect(ctx *gin.Context) {
	var param ParamCollect
	err := ctx.ShouldBind(&param)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if param.HouseUId == "" || param.UserUId == "" {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	coll := collections.NewCollectionCollect(nil)
	filter := bson.D{
		{"user_uid", param.UserUId},
		{"house_uid", param.HouseUId},
	}
	if param.Collect { // 收藏
		tCollect := collect.Collect{}
		tCollect.UserUId = param.UserUId
		tCollect.HouseUId = param.HouseUId
		_, err := coll.UpsertOne(filter, tCollect, options.Update())
		if err != nil {
			base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
	} else { // 删除收藏
		_, err := coll.MCollection().DeleteOne(context.TODO(), filter)
		if err != nil {
			base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, nil)
}
