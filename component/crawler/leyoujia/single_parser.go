package leyoujia

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
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
}

func NewSingleParser(param *crawler.SingleParam) *SingleParser {
	singleParser := crawler.NewSingleParser(param)
	return &SingleParser{
		SingleParser: singleParser,
	}
}

func (receiver *SingleParser) fetch() (*Single, error) {
	request, err := receiver.newRequest()
	if err != nil {
		return nil, err
	}
	bs, err := receiver.Do(http.DefaultClient, request)
	if err != nil {
		return nil, err
	}
	var single Single
	err = json.Unmarshal(bs, &single)
	if err != nil {
		return nil, err
	}

	if !single.Success {
		return nil, crawler.InitErr
	}

	return &single, nil
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
	single, err := receiver.fetch()
	if err != nil {
		return nil, err
	}

	tHouse := &house.House{}
	tHouse.UId = house.SourceLeyoujia + "-" + strconv.Itoa(single.Data.Zf.HouseId)
	tHouse.SourceId = strconv.Itoa(single.Data.Zf.HouseId)
	tHouse.Source = house.SourceLeyoujia
	tHouse.Type, err = receiver.getType(single)
	if err != nil {
		return nil, err
	}
	tHouse.Name = single.Data.Zf.ComName
	tHouse.Description = receiver.getDescription(single)
	tHouse.ImgUrls, tHouse.VideoUrls, err = receiver.getImgUrls(single)
	if err != nil {
		return nil, err
	}
	tHouse.Area = single.Data.Zf.IndoorArea
	price, err := receiver.getPrice(single)
	if err != nil {
		return nil, err
	}
	tHouse.Price = *price
	tHouse.Floor = single.Data.Zf.Layer
	tHouse.Location = *receiver.getLocation(single)
	// rentType todo 需对比发现
	tHouse.BuildTime = receiver.getBuildTime(single)
	tHouse.Facilities = receiver.getFacilities(single)
	tHouse.Traffic, err = receiver.getTraffic(single)
	if err != nil {
		return nil, err
	}
	tHouse.Composition = *receiver.getComposition(single)

	return tHouse, nil
}

func (receiver *SingleParser) getType(single *Single) (house.Type, error) {
	switch single.Data.Zf.PropertyType {
	case "公寓":
		return house.TypeApartment, nil
	case "普通住宅":
		return house.TypeResidence, nil
	case "别墅":
		return house.TypeVilla, nil
	default:
		return "", fmt.Errorf("乐有家 获取 房屋类型错误，原生数据：%s", single.Data.Zf.PropertyType)
	}
}

func (receiver *SingleParser) getImgUrls(single *Single) ([]string, []string, error) {
	uri := HostImage + PathImage
	//houseType := receiver.Single.Data.Zf.HouseType // ""/平层/复式/开间 当是中文时出错
	houseType := ""
	houseId := single.Data.Zf.HouseId
	comId := single.Data.Zf.ComId
	values := make(url.Values)
	values.Add("houseType", houseType)
	values.Add("houseId", strconv.Itoa(houseId))
	values.Add("comId", strconv.Itoa(comId))
	body := values.Encode()
	reader := strings.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, uri, reader)

	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var imgUrl ImgUrl
	err = json.Unmarshal(bs, &imgUrl)
	if err != nil {
		return nil, nil, err
	}

	if !imgUrl.Success {
		return nil, nil, fmt.Errorf("乐有家 获取图片失败，原始信息: %+v", imgUrl)
	}

	var imgUrls []string
	var videoUrls []string
	for _, item := range imgUrl.Data.HouseImageList {
		imgUrls = append(imgUrls, item.ImagePath)
	}

	if imgUrl.Data.HouseVideo.VideoUrl != "" {
		videoUrls = append(videoUrls, imgUrl.Data.HouseVideo.VideoUrl)
	}

	return imgUrls, videoUrls, nil
}

func (receiver *SingleParser) getPrice(single *Single) (*house.Price, error) {
	mFeeText := single.Data.Community.ManagerFee
	reg := regexp.MustCompile(`^([\d.]+)`)
	mFeeText = reg.FindString(mFeeText)
	mFee, err := strconv.ParseFloat(mFeeText, 64)
	if err != nil {
		return nil, err
	}
	return &house.Price{
		Rent:               single.Data.Zf.RentPrice,
		ManagementPerMeter: mFee,
	}, nil
}

func (receiver *SingleParser) getLocation(single *Single) *house.Location {
	zf := single.Data.Zf
	return &house.Location{
		City:   zf.CityName,
		Region: zf.AreaName,
		Extra:  zf.ComAddress,
	}
}

func (receiver *SingleParser) getBuildTime(single *Single) time.Time {
	zf := single.Data.Zf
	return time.UnixMilli(zf.CompletionDate)
}

func (receiver *SingleParser) getFacilities(single *Single) []string {
	zf := single.Data.Zf
	fur := strings.Split(zf.FurnitureAndPackage, "@")
	var facilities []string
	for _, v := range fur {
		if v == "" || v == "无" {
			continue
		}
		facilities = append(facilities, v)
	}
	return facilities
}

func (receiver *SingleParser) getTraffic(single *Single) ([]house.Traffic, error) {
	var traffics []house.Traffic
	Metros := single.Data.Community.MetrosNearby
	for _, Metro := range Metros {
		traffic := house.Traffic{}
		switch Metro.Type {
		case 1:
			traffic.Type = house.TrafficTypeSubway
		default:
			return nil, fmt.Errorf("获取交通类型失败，原始数据：%d", Metro.Type)
		}
		//traffic.Line = Metro.LineName
		traffic.Station = Metro.Name
		traffic.Distance = int(Metro.Distance)
		traffics = append(traffics, traffic)
	}
	return traffics, nil
}

func (receiver *SingleParser) getComposition(single *Single) *house.Composition {
	zf := single.Data.Zf
	return &house.Composition{
		Room:    zf.Room,
		Parlor:  zf.Parlor,
		Toilet:  zf.Toilet,
		Kitchen: zf.Kitchen,
	}
}

func (receiver *SingleParser) getDescription(single *Single) string {
	description := ""
	for _, highlight := range single.Data.HouseHighlights {
		description += highlight.Content
	}
	return description
}
