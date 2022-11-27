package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollectionCdnImage = "cdn_image"

type CdnImage struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Group    string             `bson:"group" form:"group"` // house_uid
	Index    int                `bson:"index" form:"index"` // 该分类下的第几张
	Filename string             `bson:"filename"`
	Data     []byte             `bson:"data"`
}
