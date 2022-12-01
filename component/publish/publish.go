package main

import (
	"bigCitySmallHouse/component/base/base_action"
	"bigCitySmallHouse/component/publish/model"
	"bigCitySmallHouse/mongodb/collections"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type ParamPublish struct {
	ImageNameList []string `json:"image_name_list"`
	model.Publish
}

func Publish(ctx *gin.Context) {
	var param ParamPublish
	err := ctx.ShouldBind(&param)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	for _, imageName := range param.ImageNameList {
		imgUrl := "http://localhost:10003/api/image/" + imageName
		param.ImgUrls = append(param.ImgUrls, imgUrl)
	}
	coll := collections.NewCollectionPublish(nil)
	filter := bson.D{}
	_, err = coll.UpsertOnePublish(filter, param.Publish, nil)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
}
