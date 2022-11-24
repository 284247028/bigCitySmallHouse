package main

import (
	house2 "bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/component/house_center/model/house"
	"bigCitySmallHouse/component/user_center/model/collect"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

const (
	SortPriceAsc  = "priceAsc"
	SortPriceDesc = "priceDesc"
)

type ParamReadHouse struct {
	Page      int    `form:"page"`
	HouseType string `form:"house_type"`
	Sort      string `form:"sort"`
	UserUId   string `form:"user_uid"`
}

type RHouse struct {
	RentType    string             `json:"rent_type"`
	Composition house2.Composition `json:"composition"`
	Price       house2.Price       `json:"price"`
	Location    house2.Location    `json:"location"`
	Name        string             `json:"name"`
	CoverUrl    string             `json:"cover_url"`
	IdHex       string             `json:"id_hex"`
}

func ReadHouse(ctx *gin.Context) {
	baseResponse := NewBaseResponse(ctx)
	var param ParamReadHouse
	err := ctx.ShouldBind(&param)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	coll := collections.NewCollectionHouseCenter(nil)
	size := 20
	filter, err := getFilter(&param)
	if err != nil {
		baseResponse.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	sort := getSort(&param)
	opts := options.Find()
	opts.SetSort(sort)
	cursor, err := coll.Pagination(param.Page, size, context.TODO(), filter, opts)
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

	ctx.JSON(http.StatusOK, houses2rHouse(tHouses))
}

func getFilter(param *ParamReadHouse) (bson.D, error) {
	filter := bson.D{
		{"shelve", true},
		{"house.source", "anjuke"}, // todo delete
	}
	switch param.HouseType {
	case house2.RentTypeEntire:
		filter = append(filter, bson.E{Key: "house.rentType", Value: house2.RentTypeEntire})
	case house2.RentTypeShared:
		filter = append(filter, bson.E{Key: "house.rentType", Value: house2.RentTypeShared})
	}

	if param.UserUId != "" {
		opts := options.Find()
		opts.SetProjection(bson.D{{"house_uid", 1}})
		cursor, err := collections.NewCollectionCollect(nil).MCollection().Find(context.TODO(), bson.D{{"user_uid", param.UserUId}})
		if err != nil {
			return nil, err
		}
		var collects []collect.Collect
		err = cursor.All(context.TODO(), &collects)
		if err != nil {
			return nil, err
		}
		var houseIds []string
		for _, tCollect := range collects {
			houseIds = append(houseIds, tCollect.HouseUId)
		}
		filter = append(filter, bson.E{Key: "house.uid", Value: bson.D{{"$in", houseIds}}})
	}

	return filter, nil
}

func getSort(param *ParamReadHouse) bson.D {
	sort := bson.D{}
	switch param.Sort {
	case SortPriceAsc:
		sort = append(sort, bson.E{Key: "house.price.rent", Value: 1})
	case SortPriceDesc:
		sort = append(sort, bson.E{Key: "house.price.rent", Value: -1})
	}
	return sort
}

func houses2rHouse(tHouses []house.House) []RHouse {
	var rHouses []RHouse
	for _, tHouse := range tHouses {
		rHouse := RHouse{}
		rHouse.Composition = tHouse.House.Composition
		rHouse.Price = tHouse.House.Price
		rHouse.Location = tHouse.House.Location
		rHouse.Name = tHouse.House.Name
		rHouse.RentType = tHouse.House.RentType.ToCn()
		rHouse.IdHex = tHouse.Id.Hex()
		if len(tHouse.House.ImgUrls) > 0 {
			rHouse.CoverUrl = tHouse.House.ImgUrls[0]
		}
		rHouses = append(rHouses, rHouse)
	}
	return rHouses
}
