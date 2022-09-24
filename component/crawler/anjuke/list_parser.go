package anjuke

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const ListUrl = "https://apirent.anjuke.com/zufang/wechat/rent/api_get_list?platform=ajkplugin&appname=wx&action=getList&page=1&sidDict=&page_size=10&openid=ocS7q0JWrakjtDislkipP_WP-OeU&cityId=12&city_id=12&city=&filterParams=%7B%7D&lego_appname=ajkWechat&lego_appid=wx099e0647f9a4717d&keyword=&cateId=&pageType=zufang_index_list&inFrom=index-weinituijian-b&outFrom="
const BigInt = 999999

type ListParser struct {
	*crawler.ListParser
}

func NewListParser(param *crawler.ListParam) *ListParser {
	parser := crawler.NewListParser(param)
	return &ListParser{
		ListParser: parser,
	}
}

func (receiver *ListParser) Parse() ([]house.House, *crawler.ListInfo, error) {
	list, err := receiver.fetch()
	if err != nil {
		return nil, nil, err
	}

	houseIdList := list.Data.SidDict.HouseidList
	houses := make([]house.House, 0, len(houseIdList))
	for _, houseId := range houseIdList {
		tHouse := house.House{
			UId:      house.SourceAnjuke + "-" + houseId,
			SourceId: houseId,
			Source:   house.SourceAnjuke,
		}
		houses = append(houses, tHouse)
	}

	info := crawler.ListInfo{
		IsLastPage: list.Data.LastPage,
		TotalPage:  BigInt,
	}

	return houses, &info, nil
}

func (receiver *ListParser) fetch() (*List, error) {
	req, err := http.NewRequest(http.MethodGet, ListUrl, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Set("page", strconv.Itoa(receiver.Param.Page))
	req.URL.RawQuery = query.Encode()

	req.Header.Add("Cookie", "id58=CrIfoWMfMD59zgV6LZxVAg==")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var list List
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}

	var sidDict SidDict
	err = json.Unmarshal([]byte(list.Data.SidDictStr), &sidDict)
	if err != nil {
		return nil, err
	}

	list.Data.SidDict = sidDict

	return &list, nil
}
