package leyoujia

type ImgUrl struct {
	Data struct {
		HouseImageList []struct {
			Id         *int   `json:"id"`
			ImagePath  string `json:"imagePath"`
			Type       int    `json:"type"`
			TypeString string `json:"typeString"`
		} `json:"houseImageList"`
		HouseVideo struct {
			Attributes struct {
			} `json:"attributes"`
			HouseId    int    `json:"houseId"`
			Id         int    `json:"id"`
			ImageUrl   string `json:"imageUrl"`
			Type       int    `json:"type"`
			UpdateDate int64  `json:"updateDate"`
			VideoUrl   string `json:"videoUrl"`
		} `json:"houseVideo"`
		HouseVr interface{} `json:"houseVr"`
	} `json:"data"`
	ErrorCode   string `json:"errorCode"`
	ErrorMsg    string `json:"errorMsg"`
	ErrorMsgDev string `json:"errorMsgDev"`
	Success     bool   `json:"success"`
	Tips        string `json:"tips"`
}
