package leyoujia

import (
	"bigCitySmallHouse/component/crawler"
	houseModel "bigCitySmallHouse/model/house"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const HostList = "https://steward.leyoujia.com"
const PathList = "/stewardnew/zf/queryZfList"

type ListParser struct {
	*crawler.ListParser
	List *List
}

func NewListParser(param *crawler.ListParam) *ListParser {
	listParser := crawler.NewListParser(param)
	return &ListParser{ListParser: listParser}
}

func (receiver ListParser) Parse() ([]houseModel.House, *crawler.ListInfo, error) {
	err := receiver.init()
	if err != nil {
		return nil, nil, err
	}

	items := receiver.List.Data.ZfList.Data
	houses := make([]houseModel.House, 0, len(items))
	for _, item := range items {
		tHouse := houseModel.House{}
		tHouse.SourceId = strconv.Itoa(item.HouseId)
		tHouse.Source = houseModel.SourceLeyoujia
		tHouse.UId = houseModel.SourceLeyoujia + "-" + strconv.Itoa(item.HouseId)
		//switch item.PropertyType {
		//case "公寓":
		//	house.Type = houseModel.TypeApartment
		//}
		//house.Name = item.ComAlias
		houses = append(houses, tHouse)
	}

	data := receiver.List.Data.ZfList
	info := &crawler.ListInfo{
		PageSize:   data.PageSize,
		TotalPage:  data.TotalPage,
		TotalCount: data.TotalRecord,
	}

	return houses, info, nil
}

func (receiver *ListParser) init() error {
	request, err := receiver.newRequest()
	if err != nil {
		return err
	}
	bs, err := receiver.Do(http.DefaultClient, request)
	if err != nil {
		return err
	}
	var list List
	err = json.Unmarshal(bs, &list)
	if err != nil {
		return err
	}

	if !list.Success {
		return fmt.Errorf("获取乐有家列表信息失败")
	}

	receiver.List = &list
	return nil
}

func (receiver ListParser) newRequest() (*http.Request, error) {
	uri := HostList + PathList
	body := receiver.buildBody()
	reader := bytes.NewBufferString(body)
	request, err := http.NewRequest(http.MethodPost, uri, reader)
	if err != nil {
		return nil, err
	}
	receiver.setHeader(request)
	return request, nil
}

func (receiver *ListParser) buildBody() string {
	page := receiver.ListParser.Param.Page
	values := make(url.Values)
	values.Add("starHouseQuery", "1")
	values.Add("uuid", "0000000063560a9f2070e2f300000000")
	values.Add("pageSize", "30")
	values.Add("tgLocationKey", "app_zf_list")
	values.Add("includeFrontImage", "1")
	values.Add("pageNo", strconv.Itoa(page))
	values.Add("cityCode", "001187")
	if page == 1 {
		values.Add("firstQuery", "1")
	}
	return values.Encode()
}

func (receiver *ListParser) setHeader(req *http.Request) {
	page := receiver.Param.Page
	clientId := "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66"
	timestamp := "1661178135325"
	cityCode := "citycode001187"
	firstQuery := "firstquery1"
	includeImg := "includefrontimage1"
	pageNo := "pageno" + strconv.Itoa(page)
	pageSize := "pagesize30"
	StartHouseQuery := "starhousequery1"
	tgKey := "tglocationkeyapp_zf_list"
	uuid := "uuid0000000063560a9f2070e2f300000000"

	str := clientId + PathList + timestamp + cityCode
	if page == 1 {
		str += firstQuery
	}
	str += includeImg + pageNo + pageSize + StartHouseQuery + tgKey + uuid

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
	req.Header.Add("clientId", clientId)
	req.Header.Add("cit", "001187")
	req.Header.Add("sid", "f84488a0d54bf14eaf23152bea9a5859")
	req.Header.Add("phoneOS", "android")
	req.Header.Add("imei", "862641055496861")
	req.Header.Add("version", "8.1.8")
	req.Header.Add("d", "0")
	req.Header.Add("latitude", "22.54551666666667")
	req.Header.Add("phoneModel", "MuMu")
	req.Header.Add("aid", "APP001")
	req.Header.Add("channel", "online_32")
	req.Header.Add("imsi", "")
	req.Header.Add("carries", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Content-Length", "137")
	req.Header.Add("Host", "steward.leyoujia.com")
	req.Header.Add("Connection", "Keep-Alive")
	//req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("User-Agent", "okhttp/3.9.1")
}

func sign(str string) string {
	res := md5.Sum([]byte(str))
	val := fmt.Sprintf("%x", res)
	bs := []byte(val) // 这是32字节
	res = md5.Sum(bs)
	val = fmt.Sprintf("%x", res)
	return val
}
