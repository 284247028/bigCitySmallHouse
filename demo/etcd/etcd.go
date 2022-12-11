package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"43.138.174.42:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer func(cli *clientv3.Client) {
		err = cli.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(cli)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//resp, err := cli.Put(ctx, "sample_key", "sample_value")
	//cancel()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(resp)
	respGet, err := cli.Get(ctx, "sample_key")
	cancel()
	if err != nil {
		log.Fatalln(err)
	}
	for _, kv := range respGet.Kvs {
		fmt.Printf("%s:%s\n", kv.Key, kv.Value)
	}
}
