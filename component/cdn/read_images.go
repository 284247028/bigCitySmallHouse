package main

import (
	"bigCitySmallHouse/component/base/base_action"
	"bigCitySmallHouse/component/cdn/model"
	"bigCitySmallHouse/mongodb/collections"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func DisplayImage(ctx *gin.Context) {
	filename := ctx.Param("filename")
	if filename == "" {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("空文件名"))
		return
	}

	coll := collections.NewCollectionCdnImage(nil)

	filter := bson.D{
		{"filename", filename},
	}
	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{"data", 1},
	})
	singleResult := coll.MCollection().FindOne(ctx, filter, opts)
	var cdnImage model.CdnImage
	err := singleResult.Decode(&cdnImage)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "image/jpeg", cdnImage.Data)
}
