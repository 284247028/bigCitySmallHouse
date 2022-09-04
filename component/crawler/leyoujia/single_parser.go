package leyoujia

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const HostSingle = "https://steward.leyoujia.com"
const PathSingle = "/stewardnew/zf/queryZfDetailV2"
const HostImage = "https://steward.leyoujia.com"
const PathImage = "/stewardnew/house/queryHouseImage"

type SingleParser struct {
	*crawler.SingleParser
	Single *Single
	House  *house.House
}

func NewSingleParser(param *crawler.SingleParam) *SingleParser {
	singleParser := crawler.NewSingleParser(param)
	return &SingleParser{
		SingleParser: singleParser,
		House:        &house.House{},
	}
}

func (receiver *SingleParser) init() error {
	request, err := receiver.newRequest()
	if err != nil {
		return err
	}
	bs, err := receiver.Do(http.DefaultClient, request)
	if err != nil {
		return err
	}
	var single Single
	err = json.Unmarshal(bs, &single)
	if err != nil {
		return err
	}

	if !single.Success {
		return crawler.InitErr
	}

	receiver.Single = &single

	return nil
}

func (receiver *SingleParser) newRequest() (*http.Request, error) {
	uri := HostSingle + PathSingle

	body := receiver.buildBody()
	reader := bytes.NewBufferString(body)

	request, err := http.NewRequest(http.MethodPost, uri, reader)
	if err != nil {
		return nil, err
	}
	receiver.setHeader(request)

	return request, nil
}

func (receiver *SingleParser) setHeader(req *http.Request) {
	uuid := "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66"
	timestamp := "1661523334649"

	str := uuid + PathSingle + timestamp + "houseid" + receiver.Param.Id
	clientSign := sign(str)

	req.Header.Add("ssid", "0000000063560a9f2070e2f300000000")
	req.Header.Add("androidid", "8874965091b33125")
	req.Header.Add("longitude", "114.05288999999999")
	req.Header.Add("uuid", "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66")
	req.Header.Add("mac", "08:00:27:6B:9E:8C")
	req.Header.Add("timestamp", "1661523334649")
	req.Header.Add("clientSign", clientSign)
	req.Header.Add("oaid", "0000000063560a9f2070e2f300000000")
	req.Header.Add("network", "WIFI")
	req.Header.Add("clientId", "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66")
	req.Header.Add("cit", "001187")
	req.Header.Add("sid", "cb2a1a0050e8bcab2656662965d8ca05")
	req.Header.Add("phoneOS", "android")
	req.Header.Add("imei", "862641055496861")
	req.Header.Add("tgType", "0")
	req.Header.Add("version", "8.1.8")
	req.Header.Add("d", "0")
	req.Header.Add("latitude", "22.54551666666667")
	req.Header.Add("phoneModel", "MuMu")
	req.Header.Add("aid", "APP001")
	req.Header.Add("channel", "online_32")
	req.Header.Add("imsi", "")
	req.Header.Add("carries", "0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "steward.leyoujia.com")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("User-Agent", "okhttp/3.9.1")
}

func (receiver *SingleParser) buildBody() string {
	values := make(url.Values)
	values.Add("houseId", receiver.Param.Id)
	return values.Encode()
}

func (receiver *SingleParser) Parse() (*house.House, error) {
	err := receiver.init()
	if err != nil {
		return nil, err
	}

	receiver.setId()
	receiver.setSource()
	receiver.setType()
	receiver.setName()
	err = receiver.setImgUrls()
	if err != nil {
		return nil, err
	}
	receiver.setArea()
	receiver.setPrice()
	receiver.setFloor()
	receiver.setElevator()
	receiver.setLocation()
	receiver.setBuildDate()
	receiver.setFurniture()
	receiver.setFacility()
	receiver.setTraffic()
	receiver.setComposition()

	return receiver.House, nil
}

func (receiver *SingleParser) setId() {
	receiver.House.SourceId = strconv.Itoa(receiver.Single.Data.Zf.HouseId)
}

func (receiver *SingleParser) setSource() {
	receiver.House.Source = house.SourceLeyoujia
}

func (receiver *SingleParser) setType() {
	switch receiver.Single.Data.Zf.PropertyType {
	case "公寓":
		receiver.House.Type = house.TypeApartment
	case "普通住宅":
		receiver.House.Type = house.TypeResidence
	case "别墅":
		receiver.House.Type = house.TypeVilla
	}
}

func (receiver *SingleParser) setName() {
	receiver.House.Name = receiver.Single.Data.Zf.ComName
}

func (receiver *SingleParser) setImgUrls() error {
	uri := HostImage + PathImage
	//houseType := receiver.Single.Data.Zf.HouseType // ""/平层/复式/开间 当是中文时出错
	houseType := ""
	houseId := receiver.Single.Data.Zf.HouseId
	comId := receiver.Single.Data.Zf.ComId
	values := make(url.Values)
	values.Add("houseType", houseType)
	values.Add("houseId", strconv.Itoa(houseId))
	values.Add("comId", strconv.Itoa(comId))
	body := values.Encode()
	reader := strings.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, uri, reader)

	if err != nil {
		return err
	}

	// aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66/stewardnew/house/queryHouseImage1661610412251comid36072houseid7407559housetype1
	uuid := "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66"
	timestamp := "1661438403974"
	str := uuid + PathImage + timestamp + "comid" + strconv.Itoa(comId) + "houseid" + strconv.Itoa(houseId) + "housetype" + houseType
	clientSign := sign(str)

	req.Header.Add("ssid", "0000000063560a9f2070e2f300000000")
	req.Header.Add("androidid", "8874965091b33125")
	req.Header.Add("longitude", "114.05288999999999")
	req.Header.Add("uuid", uuid)
	req.Header.Add("mac", "08:00:27:6B:9E:8C")
	req.Header.Add("timestamp", timestamp)
	req.Header.Add("clientSign", clientSign)
	req.Header.Add("oaid", "0000000063560a9f2070e2f300000000")
	req.Header.Add("network", "WIFI")
	req.Header.Add("clientId", uuid)
	req.Header.Add("cit", "001187")
	req.Header.Add("sid", "311fb7cca1cee8fa0f28dc1282f171e1")
	req.Header.Add("phoneOS", "android")
	req.Header.Add("imei", "862641055496861")
	req.Header.Add("version", "8.1.8")
	req.Header.Add("d", "0")
	req.Header.Add("latitude", "22.54551666666667")
	req.Header.Add("phoneModel", "MuMu")
	req.Header.Add("aid", "APP001")
	req.Header.Add("channel", "online_32")
	req.Header.Add("imsi", "")
	req.Header.Add("carries", "0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "steward.leyoujia.com")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("User-Agent", "okhttp/3.9.1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var imgUrl ImgUrl
	err = json.Unmarshal(bs, &imgUrl)
	if err != nil {
		return err
	}

	if !imgUrl.Success {
		return errors.New("get imgUrl")
	}

	for _, item := range imgUrl.Data.HouseImageList {
		receiver.House.ImgUrls = append(receiver.House.ImgUrls, item.ImagePath)
	}

	if imgUrl.Data.HouseVideo.VideoUrl != "" {
		receiver.House.VideoUrls = append(receiver.House.VideoUrls, imgUrl.Data.HouseVideo.VideoUrl)
	}

	return nil
}

func (receiver *SingleParser) setArea() {
	receiver.House.Area = receiver.Single.Data.Zf.IndoorArea
}

func (receiver *SingleParser) setPrice() {
	receiver.House.Price = receiver.Single.Data.Zf.RentPrice
}
func (receiver *SingleParser) setFloor() {
	receiver.House.Floor = receiver.Single.Data.Zf.Layer
}

func (receiver *SingleParser) setElevator() {
	zf := receiver.Single.Data.Zf
	receiver.House.Elevator = false
	if strings.Contains(zf.BasicPackage, "电梯") || strings.Contains(zf.Tags, "电梯") {
		receiver.House.Elevator = true
	}
}

func (receiver *SingleParser) setLocation() {
	zf := receiver.Single.Data.Zf
	receiver.House.Location.City = zf.CityName
	receiver.House.Location.Region = zf.AreaName
	receiver.House.Location.Extra = zf.ComAddress
}

func (receiver *SingleParser) setBuildDate() {
	zf := receiver.Single.Data.Zf
	receiver.House.BuildTime = time.UnixMilli(zf.CompletionDate)
}

func (receiver *SingleParser) setFurniture() {
	zf := receiver.Single.Data.Zf
	receiver.House.Furniture = strings.Split(zf.Furniture, "@")
}

func (receiver *SingleParser) setFacility() {
	zf := receiver.Single.Data.Zf
	receiver.House.Facility = strings.Split(zf.BasicPackage, "@")
}

func (receiver *SingleParser) setTraffic() {
	Metros := receiver.Single.Data.Community.MetrosNearby
	for _, Metro := range Metros {
		traffic := house.Traffic{}
		switch Metro.Type {
		case 1:
			traffic.Type = house.TrafficTypeSubway
		}
		traffic.Line = Metro.LineName
		traffic.Station = Metro.Name
		traffic.Distance = int(Metro.Distance)
		receiver.House.Traffic = append(receiver.House.Traffic, traffic)
	}
}

func (receiver *SingleParser) setComposition() {
	zf := receiver.Single.Data.Zf
	receiver.House.Composition.Room = zf.Room
	receiver.House.Composition.Parlor = zf.Parlor
	receiver.House.Composition.Toilet = zf.Toilet
}
