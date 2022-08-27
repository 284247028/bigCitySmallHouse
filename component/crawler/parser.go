package crawler

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Parser struct {
	//Req        *http.Request
	//HttpClient *http.Client
}

func NewParser() *Parser {
	return &Parser{
		//Req:        &http.Request{},
		//HttpClient: http.DefaultClient,
	}
}

func (receiver *Parser) Do(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return ioutil.ReadAll(resp.Body)
}

//func (receiver *Parser) Parse() {
//
//}
//
//func (receiver *Parser) SetUId(UId string) {
//	receiver.UId = UId
//}
//
//func (receiver *Parser) SetId(Id string) {
//	receiver.Id = Id
//}
//
//func (receiver *Parser) SetSource(Source Source) {
//	receiver.Source = Source
//}
//
//func (receiver *Parser) SetName(Name string) {
//	receiver.Name = Name
//}
//
//func (receiver *Parser) SetImgUrls(ImgUrls []string) {
//	receiver.ImgUrls = ImgUrls
//}
//
//func (receiver *Parser) SetHouseType(HouseType string) {
//	receiver.HouseType = HouseType
//}
//
//func (receiver *Parser) SetArea(Area float64) {
//	receiver.Area = Area
//}
//
//func (receiver *Parser) SetPrice(price float64) {
//	receiver.price = price
//}
//
//func (receiver *Parser) SetFloor(Floor int) {
//	receiver.Floor = Floor
//}
//
//func (receiver *Parser) SetElevator(Elevator bool) {
//	receiver.Elevator = Elevator
//}
//
//func (receiver *Parser) SetLocation(Location Location) {
//	receiver.Location = Location
//}
//
//func (receiver *Parser) SetBuildDate(BuildDate time.Time) {
//	receiver.BuildDate = BuildDate
//}
//
//func (receiver *Parser) SetFurniture(Furniture []string) {
//	receiver.Furniture = Furniture
//}
//
//func (receiver *Parser) SetSubwayStation(SubwayStation string) {
//	receiver.SubwayStation = SubwayStation
//}
//
//func (receiver *Parser) SetSubwayDistance(SubwayDistance int) {
//	receiver.SubwayDistance = SubwayDistance
//}
//
//func (receiver *Parser) SetBusStation(BusStation string) {
//	receiver.BusStation = BusStation
//}
//
//func (receiver *Parser) SetBusDistance(BusDistance int) {
//	receiver.BusDistance = BusDistance
//}
//
//func (receiver *Parser) SetParlor(Parlor int) {
//	receiver.Parlor = Parlor
//}
//
//func (receiver *Parser) SetRoom(Room int) {
//	receiver.Room = Room
//}
