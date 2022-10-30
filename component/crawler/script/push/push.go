package main

import (
	"bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/component/crawler/model/push"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
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
	packs, err := load()
	if err != nil {
		log.Fatalln(err)
	}

	err = pushPacks(packs)
	if err != nil {
		return
	}

	err = setStatus()
	if err != nil {
		log.Fatalln(err)
	}
}

func load() ([]house.Pack, error) {
	log.Println("从数据库读取需要推送的房源...")
	tCollection := collections.NewCollectionPack(nil)
	filter := bson.D{
		{"status", bson.D{
			{"$in", bson.A{house.PackStatusSingle, house.PackStatusList}},
		}},
	}

	cursor, err := tCollection.MCollection().Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var packs []house.Pack
	err = cursor.All(context.TODO(), &packs)
	if err != nil {
		return nil, err
	}
	return packs, nil
}

func pushPacks(packs []house.Pack) error {
	log.Println("开始推送")
	conn, err := amqp.Dial("amqp://root:123456@43.138.174.42:5672/")
	if err != nil {
		return err
	}
	defer func(conn *amqp.Connection) {
		_ = conn.Close()
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func(ch *amqp.Channel) {
		_ = ch.Close()
	}(ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, pack := range packs {
		tPush, err := push.Pack2push(&pack)
		if err != nil {
			return err
		}
		pushJson, err := json.Marshal(tPush)
		if err != nil {
			return err
		}
		err = ch.PublishWithContext(ctx,
			"", // exchange
			//q.Name, // routing key
			"crawler_house",
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        pushJson,
			})
		if err != nil {
			return err
		}

		log.Printf("推送成功，id:%s, status:%s\n", tPush.Id, tPush.Status)
	}

	return nil
}

func setStatus() error {
	log.Println("设置状态")
	tCollection := collections.NewCollectionPack(nil)
	filter := bson.D{
		{"status", house.PackStatusList},
	}

	update := bson.D{
		{"$set", bson.D{
			{"status", house.PackStatusFail},
		}},
	}

	result, err := tCollection.MCollection().UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	log.Println("设置失败的结果：", result)

	filter = bson.D{
		{"status", house.PackStatusSingle},
	}

	update = bson.D{
		{"$set", bson.D{
			{"status", house.PackStatusSuccess},
		}},
	}

	result, err = tCollection.MCollection().UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	log.Println("设置成功的结果：", result)
	return nil
}
