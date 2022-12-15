package main

import "bigCitySmallHouse/component/etcd"

func main() {
	etcd.GetEtcd().RegisterService("test_service_name", "test_host")
}
