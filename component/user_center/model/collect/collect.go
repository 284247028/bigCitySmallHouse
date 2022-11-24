package collect

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollectionCollect = "collect"

type Collect struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserUId  string             `bson:"user_uid"`
	HouseUId string             `bson:"house_uid"`
}
