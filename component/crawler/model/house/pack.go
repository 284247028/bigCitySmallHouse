package house

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollectionPack = "pack"

type Pack struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Status string             `bson:"status" json:"status"`
	House  House              `bson:"house" json:"house"`
}

const (
	PackStatusList    = "pack_status_list"
	PackStatusSingle  = "pack_status_single"
	PackStatusSuccess = "pack_status_success"
	PackStatusFail    = "pack_status_fail"
)
