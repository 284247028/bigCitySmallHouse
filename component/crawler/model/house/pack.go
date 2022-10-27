package house

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollectionPack = "pack"

type Pack struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Status string             `bson:"status"`
	House  House              `bson:"house"`
}

const (
	PackStatusList    = "pack_status_list"
	PackStatusSingle  = "pack_status_single"
	PackStatusSuccess = "pack_status_success"
	PackStatusFail    = "pack_status_fail"
)
