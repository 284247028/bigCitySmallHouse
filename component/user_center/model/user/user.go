package user

import (
	"crypto/md5"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const CollectionUser = "user"

type User struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	OpenId     string             `bson:"open_id"`     // 小程序用户唯一标识
	SessionKey string             `bson:"session_key"` // 小程序数据签名用
	UId        string             `bson:"uid"`         // 根据 OpenId 计算，保护 open id 不暴露
	LoginTime  time.Time          `bson:"login_time"`
}

func (receiver *User) GetUid() string {
	sumByte := md5.Sum([]byte(receiver.OpenId))
	return fmt.Sprintf("%x", sumByte)
}
