package leyoujia

type Single struct {
	Data struct {
		BranchCertificates interface{} `json:"branchCertificates"`
		Agent              interface{} `json:"agent"`
		DkCountSevenDay    int         `json:"dkCountSevenDay"`
		TopicId            interface{} `json:"topicId"`
		Recommend          interface{} `json:"recommend"`
		HouseHighlights    []struct {
			Content string `json:"content"`
			Title   string `json:"title"`
		} `json:"houseHighlights"`
		Zf struct {
			AreaCode   string `json:"areaCode"`
			AreaName   string `json:"areaName"`
			Attributes struct {
			} `json:"attributes"`
			BasicPackage         string      `json:"basicPackage"`
			BuildingArea         float64     `json:"buildingArea"`
			CityCode             string      `json:"cityCode"`
			CityName             string      `json:"cityName"`
			Collected            interface{} `json:"collected"`
			ComAddress           string      `json:"comAddress"`
			ComAlias             string      `json:"comAlias"`
			ComId                int         `json:"comId"`
			ComName              string      `json:"comName"`
			CompletionDate       int64       `json:"completionDate"`
			CompletionDateString string      `json:"completionDateString"`
			DeployDate           int64       `json:"deployDate"`
			DgId                 int         `json:"dgId"`
			DpCount              int         `json:"dpCount"`
			Fitment              string      `json:"fitment"`
			FloorHeightString    string      `json:"floorHeightString"`
			Furniture            string      `json:"furniture"`
			FurnitureAndPackage  string      `json:"furnitureAndPackage"`
			FxWorkerId           string      `json:"fxWorkerId"`
			HasAlbum             bool        `json:"hasAlbum"`
			HasElevator          int         `json:"hasElevator"`
			HasKey               bool        `json:"hasKey"`
			HouseAge             interface{} `json:"houseAge"`
			HouseId              int         `json:"houseId"`
			HouseNumber          string      `json:"houseNumber"`
			HouseType            string      `json:"houseType"`
			ImageUrl             string      `json:"imageUrl"`
			IndoorArea           float64     `json:"indoorArea"`
			IsNearSubway         interface{} `json:"isNearSubway"`
			IsOutsideShow        int         `json:"isOutsideShow"`
			JointRent            int         `json:"jointRent"`
			Kitchen              int         `json:"kitchen"`
			Latitude             float64     `json:"latitude"`
			Layer                int         `json:"layer"`
			LayerHeight          int         `json:"layerHeight"`
			LayoutId             int         `json:"layoutId"`
			Longitude            float64     `json:"longitude"`
			LpWorkerId           string      `json:"lpWorkerId"`
			LupDate              int64       `json:"lupDate"`
			NearbyTraffic        string      `json:"nearbyTraffic"`
			Orientation          string      `json:"orientation"`
			Parlor               int         `json:"parlor"`
			Payment              string      `json:"payment"`
			PaymentWay           string      `json:"paymentWay"`
			PlaceCode            string      `json:"placeCode"`
			PlaceName            string      `json:"placeName"`
			Portal               int         `json:"portal"`
			PropertyType         string      `json:"propertyType"`
			RailIds              []int       `json:"railIds"`
			RecommendReasons     string      `json:"recommendReasons"`
			RecommendScore       interface{} `json:"recommendScore"`
			RecommendType        interface{} `json:"recommendType"`
			Remark               string      `json:"remark"`
			RentMark             float64     `json:"rentMark"`
			RentPrice            float64     `json:"rentPrice"`
			RentRankScore        int         `json:"rentRankScore"`
			RentUnitPrice        interface{} `json:"rentUnitPrice"`
			RightYear            string      `json:"rightYear"`
			Room                 int         `json:"room"`
			SaleStatus           int         `json:"saleStatus"`
			SameComMaxPrice      float64     `json:"sameComMaxPrice"`
			SameComMinPrice      float64     `json:"sameComMinPrice"`
			Sign                 interface{} `json:"sign"`
			SkWorkerId           string      `json:"skWorkerId"`
			StarHouse            bool        `json:"starHouse"`
			State                int         `json:"state"`
			StationIds           []int       `json:"stationIds"`
			SurfaceLayer         int         `json:"surfaceLayer"`
			SxHsl                bool        `json:"sxHsl"`
			SxWorkerId           string      `json:"sxWorkerId"`
			Tags                 string      `json:"tags"`
			TgHouse              bool        `json:"tgHouse"`
			TgId                 interface{} `json:"tgId"`
			TgLocationId         interface{} `json:"tgLocationId"`
			TgType               interface{} `json:"tgType"`
			TgWorkerId           string      `json:"tgWorkerId"`
			Title                string      `json:"title"`
			ToDayReserved        bool        `json:"toDayReserved"`
			Toilet               int         `json:"toilet"`
			UniqueKey            string      `json:"uniqueKey"`
			VideoHouse           bool        `json:"videoHouse"`
			VrHouse              bool        `json:"vrHouse"`
			YsWorkerId           string      `json:"ysWorkerId"`
			ZrWorkerId           string      `json:"zrWorkerId"`
		} `json:"zf"`
		Community struct {
			Address  string `json:"address"`
			AerialVR struct {
				ComId      int         `json:"comId"`
				CreateDate int64       `json:"createDate"`
				Id         int         `json:"id"`
				ImageUrl   string      `json:"imageUrl"`
				Status     int         `json:"status"`
				Type       interface{} `json:"type"`
				UpdateDate int64       `json:"updateDate"`
				VrUrl      string      `json:"vrUrl"`
			} `json:"aerialVR"`
			AerialVRUrl string `json:"aerialVRUrl"`
			AirHouse    bool   `json:"airHouse"`
			Albums      []struct {
				SyPath     string `json:"syPath"`
				Type       int    `json:"type"`
				TypeString string `json:"typeString"`
			} `json:"albums"`
			Alias           string      `json:"alias"`
			Alias2          string      `json:"alias2"`
			Alias3          string      `json:"alias3"`
			AliasSpell      string      `json:"aliasSpell"`
			AreaCode        string      `json:"areaCode"`
			AreaName        string      `json:"areaName"`
			AttentionCount  interface{} `json:"attentionCount"`
			AttentionString string      `json:"attentionString"`
			Attributes      struct {
			} `json:"attributes"`
			AvgPrice               float64       `json:"avgPrice"`
			BuildingArea           float64       `json:"buildingArea"`
			BuildingStructure      string        `json:"buildingStructure"`
			BusesNearBy            []interface{} `json:"busesNearBy"`
			CityCode               string        `json:"cityCode"`
			CityName               string        `json:"cityName"`
			CitySpell              string        `json:"citySpell"`
			CjCount                interface{}   `json:"cjCount"`
			Collected              bool          `json:"collected"`
			CompletionDate         int64         `json:"completionDate"`
			CompletionDateString   string        `json:"completionDateString"`
			Cqnx                   string        `json:"cqnx"`
			CubageRate             float64       `json:"cubageRate"`
			Developer              string        `json:"developer"`
			DgCount                interface{}   `json:"dgCount"`
			DjReferenceWord        string        `json:"djReferenceWord"`
			Environment            string        `json:"environment"`
			GrantDate              int64         `json:"grantDate"`
			GreenRate              float64       `json:"greenRate"`
			HasVillage             int           `json:"hasVillage"`
			HeatingType            string        `json:"heatingType"`
			HouseAge               int           `json:"houseAge"`
			HydroPowerGas          string        `json:"hydroPowerGas"`
			Id                     int           `json:"id"`
			ImageUrl               string        `json:"imageUrl"`
			InnerFacility          string        `json:"innerFacility"`
			Introduction           string        `json:"introduction"`
			IsHotCommunity         int           `json:"isHotCommunity"`
			Latitude               float64       `json:"latitude"`
			Longitude              float64       `json:"longitude"`
			ManagerCompany         string        `json:"managerCompany"`
			ManagerFee             string        `json:"managerFee"`
			ManagerTelephone       string        `json:"managerTelephone"`
			MetroDistanceShowValue string        `json:"metroDistanceShowValue"`
			MetrosNearby           []struct {
				ConcatLineName string  `json:"concatLineName"`
				Distance       float64 `json:"distance"`
				LineName       string  `json:"lineName"`
				Name           string  `json:"name"`
				RailId         string  `json:"railId"`
				StationId      string  `json:"stationId"`
				Type           int     `json:"type"`
			} `json:"metrosNearby"`
			Name                    string      `json:"name"`
			NameSpell               string      `json:"nameSpell"`
			OriginalAvgPrice        float64     `json:"originalAvgPrice"`
			ParkingCount            int         `json:"parkingCount"`
			PlaceCode               string      `json:"placeCode"`
			PlaceName               string      `json:"placeName"`
			PlanBuildingNumber      int         `json:"planBuildingNumber"`
			PropertyType            string      `json:"propertyType"`
			RankingShowValue        string      `json:"rankingShowValue"`
			ReferenceUnitPrice      interface{} `json:"referenceUnitPrice"`
			Remark                  string      `json:"remark"`
			RentCount               int         `json:"rentCount"`
			RightType               string      `json:"rightType"`
			SaleCount               int         `json:"saleCount"`
			SchoolDistanceShowValue string      `json:"schoolDistanceShowValue"`
			SellCount               int         `json:"sellCount"`
			ShowOffer               int         `json:"showOffer"`
			SoldCount               interface{} `json:"soldCount"`
			StreetName              string      `json:"streetName"`
			TotalHouse              int         `json:"totalHouse"`
			TrafficVos              []struct {
				ConcatLineName string  `json:"concatLineName"`
				Distance       float64 `json:"distance"`
				LineName       string  `json:"lineName"`
				Name           string  `json:"name"`
				RailId         string  `json:"railId"`
				StationId      string  `json:"stationId"`
				Type           int     `json:"type"`
			} `json:"trafficVos"`
		} `json:"community"`
		Agents []struct {
			AllTags []struct {
				Id          int    `json:"id"`
				InsertTime  int64  `json:"insertTime"`
				Status      int    `json:"status"`
				TagCategory int    `json:"tagCategory"`
				TagName     string `json:"tagName"`
				TagOrder    int    `json:"tagOrder"`
				TagType     int    `json:"tagType"`
				UpdateTime  int64  `json:"updateTime"`
				WorkerId    string `json:"workerId"`
			} `json:"allTags"`
			BestAgent          bool        `json:"bestAgent"`
			CityCode           string      `json:"cityCode"`
			CommentInfo        string      `json:"commentInfo"`
			CommentTitle       string      `json:"commentTitle"`
			DeptId             interface{} `json:"deptId"`
			DeptName           string      `json:"deptName"`
			DutyId             int         `json:"dutyId"`
			DutyName           string      `json:"dutyName"`
			DutyNameCN         string      `json:"dutyNameCN"`
			DutyNameFull       string      `json:"dutyNameFull"`
			EntryYear          float64     `json:"entryYear"`
			ExtNum             string      `json:"extNum"`
			HasQualification   bool        `json:"hasQualification"`
			HeadTag            string      `json:"headTag"`
			HistoryTotalRecord int         `json:"historyTotalRecord"`
			ImRate             interface{} `json:"imRate"`
			IsDz               int         `json:"isDz"`
			IsUniqueAgent      interface{} `json:"isUniqueAgent"`
			IsXqExpert         interface{} `json:"isXqExpert"`
			KnowAverage        interface{} `json:"knowAverage"`
			LeadReputably      *float64    `json:"leadReputably"`
			MainExtNum         string      `json:"mainExtNum"`
			MainNum            string      `json:"mainNum"`
			Mobile             string      `json:"mobile"`
			Name               string      `json:"name"`
			PjCount            int         `json:"pjCount"`
			Portrait           string      `json:"portrait"`
			ProfressAverage    interface{} `json:"profressAverage"`
			SchoolExpert       bool        `json:"schoolExpert"`
			Score              float64     `json:"score"`
			SeeCount           int         `json:"seeCount"`
			SeeTime            interface{} `json:"seeTime"`
			ServerCustomers    interface{} `json:"serverCustomers"`
			ServiceArea        string      `json:"serviceArea"`
			ServiceAverage     interface{} `json:"serviceAverage"`
			ServicePersonCount int         `json:"servicePersonCount"`
			ShowType           interface{} `json:"showType"`
			ShowTypeId         interface{} `json:"showTypeId"`
			State              int         `json:"state"`
			Status             int         `json:"status"`
			StorePlace         string      `json:"storePlace"`
			Tag                string      `json:"tag"`
			Tags               string      `json:"tags"`
			TagsNew            []string    `json:"tagsNew"`
			TeamId             interface{} `json:"teamId"`
			TeamName           string      `json:"teamName"`
			TgAgent            bool        `json:"tgAgent"`
			TgId               interface{} `json:"tgId"`
			TgType             interface{} `json:"tgType"`
			ThirtyDaySee       *int        `json:"thirtyDaySee"`
			ThirtyDaySeeNew    int         `json:"thirtyDaySeeNew"`
			WorkerId           string      `json:"workerId"`
			WorkerNo           string      `json:"workerNo"`
			WorkerState        int         `json:"workerState"`
			WorkerStatus       int         `json:"workerStatus"`
			WorkerYear         int         `json:"workerYear"`
			WorkerYearShow     string      `json:"workerYearShow"`
		} `json:"agents"`
	} `json:"data"`
	ErrorCode   string `json:"errorCode"`
	ErrorMsg    string `json:"errorMsg"`
	ErrorMsgDev string `json:"errorMsgDev"`
	Success     bool   `json:"success"`
	Tips        string `json:"tips"`
}