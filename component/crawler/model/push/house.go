package push

import (
	"bigCitySmallHouse/component/crawler/model/house"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	StatusPushValid   = "push_valid"
	StatusPushInvalid = "push_invalid"
)

type Push struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Status string             `bson:"status"`
	House  house.House        `bson:"house"`
}

func Packs2pushes(packs []house.Pack) ([]Push, error) {
	tPushes := make([]Push, 0, len(packs))
	for _, pack := range packs {
		push, err := Pack2push(&pack)
		if err != nil {
			return nil, err
		}
		tPushes = append(tPushes, *push)
	}

	return tPushes, nil
}

func Pack2push(pack *house.Pack) (*Push, error) {
	if pack == nil {
		return nil, fmt.Errorf("nil pointer pack in Pack2push")
	}
	status, err := PackStatus2pushStatus(pack.Status)
	if err != nil {
		return nil, err
	}
	return &Push{
		Id:     pack.Id,
		Status: status,
		House:  pack.House,
	}, nil
}

func PackStatus2pushStatus(status string) (string, error) {
	switch status {
	case house.PackStatusSingle:
		return StatusPushValid, nil
	case house.PackStatusList:
		return StatusPushInvalid, nil
	}
	return "", fmt.Errorf("传入错误的pack status: %s", status)
}
