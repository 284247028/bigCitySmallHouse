package tests

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestSign(t *testing.T) {
	str := "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66/stewardnew/zf/queryZfList1660961249090citycode001187firstquery1includefrontimage1pageno1pagesize30starhousequery1tglocationkeyapp_zf_listuuid0000000063560a9f2070e2f300000000"
	res := md5.Sum([]byte(str))
	val := fmt.Sprintf("%x", res)
	bytes := []byte(val) // 这是32字节
	compare := res[:]    // 这是16字节
	res = md5.Sum(bytes)
	val = fmt.Sprintf("%x", res)
	log.Printf(val, compare)
}

const Domain = "https://steward.leyoujia.com"
const Path = "/stewardnew/zf/queryZfList"

func TestCrawler(t *testing.T) {
	// starHouseQuery=1&uuid=0000000063560a9f2070e2f300000000&pageSize=30&tgLocationKey=app_zf_list&includeFrontImage=1&pageNo=1&cityCode=001187&firstQuery=1
	reqUrl := Domain + Path

	val := make(url.Values, 0)
	val.Add("starHouseQuery", "1")
	val.Add("uuid", "0000000063560a9f2070e2f300000000")
	val.Add("pageSize", "30")
	val.Add("tgLocationKey", "app_zf_list")
	val.Add("includeFrontImage", "1")
	val.Add("pageNo", "1")
	val.Add("cityCode", "001187")
	val.Add("firstQuery", "1")
	query := val.Encode()

	body := bytes.NewReader([]byte(query))
	req, err := http.NewRequest(http.MethodPost, reqUrl, body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("ssid", "0000000063560a9f2070e2f300000000")
	req.Header.Add("androidid", "8874965091b33125")
	req.Header.Add("longitude", "114.05288999999999")
	req.Header.Add("uuid", "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66")
	req.Header.Add("mac", "08:00:27:6B:9E:8C")
	req.Header.Add("timestamp", "1661178039631")
	req.Header.Add("clientSign", "e8905b206d27c11aeee4650927935717")
	req.Header.Add("oaid", "0000000063560a9f2070e2f300000000")
	req.Header.Add("network", "WIFI")
	req.Header.Add("clientId", "aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66")
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
	req.Header.Add("carries", "0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "150")
	req.Header.Add("Host", "steward.leyoujia.com")
	req.Header.Add("Connection", "Keep-Alive")
	//req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("User-Agent", "okhttp/3.9.1")

	httpClient := http.DefaultClient

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(bs)

}

func leyoujiaSign(str string) string {
	res := md5.Sum([]byte(str))
	val := fmt.Sprintf("%x", res)
	bytes := []byte(val) // 这是32字节
	res = md5.Sum(bytes)
	val = fmt.Sprintf("%x", res)
	return val
}

// page 1
// citycode001187firstquery1includefrontimage1pageno1pagesize30starhousequery1tglocationkeyapp_zf_listuuid0000000063560a9f2070e2f300000000

// page 2
// citycode001187           includefrontimage1pageno2pagesize30starhousequery1tglocationkeyapp_zf_listuuid0000000063560a9f2070e2f300000000

/*
POST https://steward.leyoujia.com/stewardnew/zf/queryZfList HTTP/1.1
ssid: 0000000063560a9f2070e2f300000000
androidid: 8874965091b33125
longitude: 114.05288999999999
uuid: aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66
mac: 08:00:27:6B:9E:8C
timestamp: 1661178039631
clientSign: e8905b206d27c11aeee4650927935717
oaid: 0000000063560a9f2070e2f300000000
network: WIFI
clientId: aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66
cit: 001187
sid: f84488a0d54bf14eaf23152bea9a5859
phoneOS: android
imei: 862641055496861
version: 8.1.8
d: 0
latitude: 22.54551666666667
phoneModel: MuMu
aid: APP001
channel: online_32
imsi:
carries: 0
Content-Type: application/x-www-form-urlencoded
Content-Length: 150
Host: steward.leyoujia.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.9.1

starHouseQuery=1&uuid=0000000063560a9f2070e2f300000000&pageSize=30&tgLocationKey=app_zf_list&includeFrontImage=1&pageNo=1&cityCode=001187&firstQuery=1
*/

/*
POST https://steward.leyoujia.com/stewardnew/zf/queryZfList HTTP/1.1
ssid: 0000000063560a9f2070e2f300000000
androidid: 8874965091b33125
longitude: 114.05288999999999
uuid: aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66
mac: 08:00:27:6B:9E:8C
timestamp: 1661178135325
clientSign: 6da48cc7d0249590ed33d301992112fc
oaid: 0000000063560a9f2070e2f300000000
network: WIFI
clientId: aa2f03e4-798e-4b3d-8d8e-6d7d286c7f66
cit: 001187
sid: f84488a0d54bf14eaf23152bea9a5859
phoneOS: android
imei: 862641055496861
version: 8.1.8
d: 0
latitude: 22.54551666666667
phoneModel: MuMu
aid: APP001
channel: online_32
imsi:
carries: 0
Content-Type: application/x-www-form-urlencoded
Content-Length: 137
Host: steward.leyoujia.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.9.1

starHouseQuery=1&uuid=0000000063560a9f2070e2f300000000&pageSize=30&tgLocationKey=app_zf_list&includeFrontImage=1&pageNo=2&cityCode=001187
*/
