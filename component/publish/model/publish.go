package model

import (
	"bigCitySmallHouse/component/house_center/model/house"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const CollectionPublish = "publish"

type Publish struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	UserUId     string             `bson:"user_uid" json:"user_uid"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Floor       int                `bson:"floor" json:"floor"`
	house.Price
	house.Composition
	RentType   house.RentType   `bson:"rent_type" json:"rent_type"`
	HouseType  house.Type       `bson:"house_type" json:"house_type"`
	Facilities []house.Facility `bson:"facilities" json:"facilities"`
	house.Location
	ImgUrls  []string  `bson:"img_urls" json:"img_urls"`
	Status   Status    `bson:"status"`
	CreateAt time.Time `bson:"create_at"`
	UpdateAt time.Time `bson:"update_at"`
	DeleteAt time.Time `bson:"delete_at"`
}

type Status string

const (
	StatusNotReview     Status = "not_review"      // 未审核
	StatusReviewPass    Status = "review_pass"     // 审核通过
	StatusReviewNotPass Status = "review_not_pass" // 审核未通过
)
