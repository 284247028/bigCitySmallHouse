package leyoujia

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model"
	"bytes"
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const Domain = "https://steward.leyoujia.com"
const Path = "/stewardnew/zf/queryZfList"

type Crawler struct {
	*crawler.HttpCrawler
}

func NewCrawler(page int) *Crawler {
	httpCrawler := crawler.NewHttpCrawler()
	qCrawler := &Crawler{HttpCrawler: httpCrawler}
	qCrawler.Page = page
	return qCrawler
}

func (receiver *Crawler) Init() error {
	err := receiver.setReq()
	if err != nil {
		return err
	}
	receiver.setHeader()
	receiver.setHttpClient()
	return nil
}

func (receiver *Crawler) setReq() error {
	host := Domain + Path
	bodyStr := receiver.buildBody()
	reader := bytes.NewReader([]byte(bodyStr))
	req, err := http.NewRequest(http.MethodPost, host, reader)
	if err != nil {
		return err
	}
	receiver.Req = req
	return nil
}

func (receiver *Crawler) setHttpClient() {
	receiver.HttpClient = http.DefaultClient
}

func (receiver *Crawler) setHeader() {
	clientId := "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66"
	timestamp := "1661178135325"
	cityCode := "citycode001187"
	firstQuery := "firstquery1"
	includeImg := "includefrontimage1"
	pageNo := "pageno" + strconv.Itoa(receiver.Page)
	pageSize := "pagesize30"
	StartHouseQuery := "starhousequery1"
	tgKey := "tglocationkeyapp_zf_list"
	uuid := "uuid0000000063560a9f2070e2f300000000"

	str := clientId + Path + timestamp + cityCode
	if receiver.Page == 1 {
		str += firstQuery
	}
	str += includeImg + pageNo + pageSize + StartHouseQuery + tgKey + uuid

	clientSign := sign(str)

	receiver.Req.Header.Add("ssid", "0000000063560a9f2070e2f300000000")
	receiver.Req.Header.Add("androidid", "8874965091b33125")
	receiver.Req.Header.Add("longitude", "114.05288999999999")
	receiver.Req.Header.Add("uuid", uuid)
	receiver.Req.Header.Add("mac", "08:00:27:6B:9E:8C")
	receiver.Req.Header.Add("timestamp", timestamp)
	receiver.Req.Header.Add("clientSign", clientSign)
	receiver.Req.Header.Add("oaid", "0000000063560a9f2070e2f300000000")
	receiver.Req.Header.Add("network", "WIFI")
	receiver.Req.Header.Add("clientId", clientId)
	receiver.Req.Header.Add("cit", "001187")
	receiver.Req.Header.Add("sid", "f84488a0d54bf14eaf23152bea9a5859")
	receiver.Req.Header.Add("phoneOS", "android")
	receiver.Req.Header.Add("imei", "862641055496861")
	receiver.Req.Header.Add("version", "8.1.8")
	receiver.Req.Header.Add("d", "0")
	receiver.Req.Header.Add("latitude", "22.54551666666667")
	receiver.Req.Header.Add("phoneModel", "MuMu")
	receiver.Req.Header.Add("aid", "APP001")
	receiver.Req.Header.Add("channel", "online_32")
	receiver.Req.Header.Add("imsi", "")
	receiver.Req.Header.Add("carries", "")
	receiver.Req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//receiver.Req.Header.Add("Content-Length", "137")
	receiver.Req.Header.Add("Host", "steward.leyoujia.com")
	receiver.Req.Header.Add("Connection", "Keep-Alive")
	//receiver.Req.Header.Add("Accept-Encoding", "gzip")
	receiver.Req.Header.Add("User-Agent", "okhttp/3.9.1")
}

func (receiver *Crawler) buildBody() string {
	values := make(url.Values)
	values.Add("starHouseQuery", "1")
	values.Add("uuid", "0000000063560a9f2070e2f300000000")
	values.Add("pageSize", "30")
	values.Add("tgLocationKey", "app_zf_list")
	values.Add("includeFrontImage", "1")
	values.Add("pageNo", strconv.Itoa(receiver.Page))
	values.Add("cityCode", "001187")
	if receiver.Page == 1 {
		values.Add("firstQuery", "1")
	}
	return values.Encode()
}

func (receiver *Crawler) Parse() ([]model.House, error) {
	//TODO implement me
	panic("implement me")
}

func sign(str string) string {
	res := md5.Sum([]byte(str))
	val := fmt.Sprintf("%x", res)
	bs := []byte(val) // 这是32字节
	res = md5.Sum(bs)
	val = fmt.Sprintf("%x", res)
	return val
}
