package main

import (
	"bigCitySmallHouse/component/crawler/model/push"
	"bigCitySmallHouse/component/house_center/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func init() {
	opts := mongodb.NewOptions()
	opts.Uri = "mongodb://admin:admin@43.138.174.42:27017/"
	err := mongodb.NewDB().ConnectMongodb(opts)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.Println("开始接收房源入库")
	conn, err := amqp.Dial("amqp://root:123456@43.138.174.42:5672/")
	if err != nil {
		log.Fatalln(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	msgs, err := ch.Consume(
		"crawler_house", // queue
		"house_center",  // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Fatalln(err)
	}

	coll := collections.NewCollectionHouseCenter(nil)

	for d := range msgs {
		var tPush push.Push
		err = json.Unmarshal(d.Body, &tPush)
		if err != nil {
			log.Fatalln(err)
		}

		tHouse := house.House{}
		tHouse.House = tPush.House
		tHouse.UpdateAt = time.Now()
		tHouse.Shelve = tPush.Status == push.StatusPushValid

		result, err := coll.MCollection().InsertOne(context.TODO(), tHouse)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(result)
	}
	log.Println("正常结束")
}
