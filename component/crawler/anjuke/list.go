package anjuke

type List struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		IsToast  bool `json:"isToast"`
		Activity struct {
			Title string `json:"title"`
		} `json:"activity"`
		Configs struct {
			Banner       []interface{} `json:"banner"`
			HouseSubject []interface{} `json:"houseSubject"`
		} `json:"configs"`
		Tkd struct {
			Title       string `json:"title"`
			Keywords    string `json:"keywords"`
			Description string `json:"description"`
		} `json:"tkd"`
		Gongyu struct {
			IsShow bool `json:"isShow"`
		} `json:"gongyu"`
		GoldOwnerCity       bool `json:"goldOwnerCity"`
		ShengxinzuOwnerCity bool `json:"shengxinzuOwnerCity"`
		BaozhangOwnerCity   bool `json:"baozhangOwnerCity"`
		LastPage            bool `json:"lastPage"`
		List                []struct {
			Infoid          string `json:"infoid"`
			CateName        string `json:"cateName"`
			CateID          string `json:"cateID"`
			Title           string `json:"title"`
			Price           string `json:"price"`
			Img             string `json:"img"`
			Unit            string `json:"unit"`
			Coll            string `json:"coll"`
			SubwayDistance  string `json:"subwayDistance"`
			FangwuLiangDian []struct {
				Text string `json:"text"`
			} `json:"fangwuLiangDian"`
			Area            string `json:"area"`
			Toward          string `json:"toward"`
			Shangquan       string `json:"shangquan"`
			IsAnxuan        string `json:"isAnxuan"`
			Location        bool   `json:"location"`
			SourceType      string `json:"source_type"`
			Isauction       string `json:"isauction"`
			AuctionShowText string `json:"auctionShowText"`
			LegoAuction     string `json:"legoAuction"`
			Tid             string `json:"tid"`
			IsLiving        bool   `json:"isLiving"`
			IsVR            bool   `json:"isVR"`
			DistanceDesc    string `json:"distanceDesc"`
			IsGoldOwner     bool   `json:"isGoldOwner"`
			Shengxinzu      string `json:"shengxinzu"`
			Recommend       string `json:"recommend"`
			Soj             struct {
				Infoid              string `json:"infoid"`
				Houseid             string `json:"houseid"`
				GTID                string `json:"GTID"`
				IsBiz               bool   `json:"is_biz"`
				IsDown              string `json:"is_down"`
				Slot                string `json:"slot"`
				Sid                 string `json:"sid"`
				AdType              string `json:"ad_type"`
				IsBusiness          string `json:"is_business"`
				Gpos                string `json:"gpos"`
				Pos                 string `json:"pos"`
				Shengxinzu          string `json:"shengxinzu"`
				Shidiheyanzhuangtai string `json:"shidiheyanzhuangtai"`
				Qingheyanzhuangtai  string `json:"qingheyanzhuangtai"`
				Recomshowlog        string `json:"recomshowlog"`
				Tid                 string `json:"tid"`
				Anxuan              string `json:"anxuan"`
				Qiyeanxuan          string `json:"qiyeanxuan"`
				IsVr                string `json:"isVr"`
				InfoType            string `json:"info_type"`
			} `json:"soj"`
		} `json:"list"`
		SidDictStr string `json:"sidDict"`
		SidDict    SidDict
	} `json:"data"`
}

type SidDict struct {
	AbTest       string   `json:"ab_test"`
	Cate2        string   `json:"cate2"`
	Cate1        string   `json:"cate1"`
	GTID         string   `json:"GTID"`
	SessionId    string   `json:"sessionId"`
	Sid          string   `json:"sid"`
	Cityid158    int      `json:"cityid1_58"`
	Nameoflist   string   `json:"nameoflist"`
	Pagesource   string   `json:"pagesource"`
	Recomshowlog string   `json:"recomshowlog"`
	HouseidList  []string `json:"houseid_list"`
	Page         string   `json:"page"`
	ListKeywords struct {
		Search struct {
			Keyword string `json:"keyword"`
		} `json:"search"`
		Filter []interface{} `json:"filter"`
	} `json:"list_keywords"`
}
