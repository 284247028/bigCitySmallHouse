package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const CollectionPublish = "publish"

type Publish struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Title          string             `bson:"title" json:"title"`
	Description    string             `bson:"description" json:"description"`
	Floor          int                `bson:"floor" json:"floor"`
	Price          float64            `bson:"price" json:"price"`                     // 租金
	WaterFee       float64            `bson:"water_fee" json:"water_fee"`             // 单位：元/吨
	ElectricityFee float64            `bson:"electricity_fee" json:"electricity_fee"` // 单位：元/度
	Room           int                `bson:"room" json:"room"`
	Parlor         int                `bson:"parlor" json:"parlor"`
	RentType       string             `bson:"rent_type" json:"rent_type"`
	HouseType      string             `bson:"house_type" json:"house_type"`
	Facilities     []string           `bson:"facilities" json:"facilities"`
	Province       string             `bson:"province" json:"province"`
	City           string             `bson:"city" json:"city"`
	Region         string             `bson:"region" json:"region"`
	LocationExtra  string             `bson:"location_extra" json:"location_extra"`
	ImgUrls        []string           `bson:"img_urls" json:"img_urls"`
	UserUId        string             `bson:"user_uid" json:"user_uId"`
	Status         Status             `bson:"status"`
	CreateAt       time.Time          `bson:"create_at"`
	UpdateAt       time.Time          `bson:"update_at"`
	DeleteAt       time.Time          `bson:"delete_at"`
}

type Status string

const (
	StatusNotReview     Status = "not_review"      // 未审核
	StatusReviewPass    Status = "review_pass"     // 审核通过
	StatusReviewNotPass Status = "review_not_pass" // 审核未通过
)
