package main

import (
	"bigCitySmallHouse/component/base/base_action"
	"bigCitySmallHouse/component/user_center/model/user"
	"bigCitySmallHouse/constant/wx"
	"bigCitySmallHouse/mongodb/collections"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type ParamLogin struct {
	LoginCode string `form:"login_code"`
}

type WxLoginData struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type Resp struct {
	UId string `json:"uid"`
}

func Login(ctx *gin.Context) {
	var param ParamLogin
	err := ctx.ShouldBind(&param)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	appId := wx.AppId
	secret := wx.Secret
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + secret + "&js_code=" + param.LoginCode + "&grant_type=authorization_code"
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	var respData WxLoginData
	err = json.Unmarshal(body, &respData)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if respData.ErrCode != 0 {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("请求登录小程序失败，错误码 %d 错误信息 %s", respData.ErrCode, respData.ErrMsg))
		return
	}

	coll := collections.NewCollectionUser(nil)
	filter := bson.D{
		{"open_id", respData.Openid},
	}

	tUser := user.User{
		OpenId:     respData.Openid,
		SessionKey: respData.SessionKey,
		LoginTime:  time.Now(),
	}
	tUser.UId = tUser.GetUid()
	_, err = coll.UpsertOne(filter, tUser, options.Update())
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, Resp{
		UId: tUser.UId,
	})
}
