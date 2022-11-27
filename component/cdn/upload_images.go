package main

import (
	"bigCitySmallHouse/component/base/base_action"
	"bigCitySmallHouse/component/cdn/model"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UploadImages(ctx *gin.Context) {
	var cdnImage model.CdnImage
	err := ctx.ShouldBind(&cdnImage)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	fileReader, err := ctx.FormFile("image_upload")
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	file, err := fileReader.Open()
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	imageBytes, err := ioutil.ReadAll(file)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	cdnImage.Data = imageBytes
	cdnImage.Filename = cdnImage.Group + "-" + strconv.Itoa(cdnImage.Index)
	coll := collections.NewCollectionCdnImage(nil)
	_, err = coll.MCollection().InsertOne(context.TODO(), cdnImage)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
