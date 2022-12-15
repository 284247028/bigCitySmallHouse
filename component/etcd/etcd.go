package etcd

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	etcdClientV3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var _etcd *Etcd

func init() {
	var environment string
	flag.StringVar(&environment, "environment", EnvDevelopment, "指定运行环境，开发环境-development，测试环境-test，生产环境-product")
	flag.Parse()
	etcdClient, err := etcdClientV3.New(etcdClientV3.Config{
		Endpoints:   []string{"43.138.174.42:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatalf("创建etcd client失败：%s", err.Error())
	}
	_etcd = &Etcd{
		environment: environment,
		Client:      etcdClient,
	}
	initConfig(environment)
}

func initConfig(environment string) {
	switch environment {
	case EnvDevelopment:
		gin.SetMode(gin.DebugMode)
	case EnvTest:
		gin.SetMode(gin.TestMode)
	case EnvProduct:
		gin.SetMode(gin.ReleaseMode)
	}
}

type Etcd struct {
	environment string
	*etcdClientV3.Client
}

func GetEtcd() *Etcd {
	return _etcd
}

func (receiver *Etcd) RegisterService(serviceName, Host string) {
	go receiver.registerService(serviceName, Host)
}

func (receiver *Etcd) registerService(serviceName, Host string) {
	leaseGrantResponse, err := receiver.Grant(context.TODO(), 3)
	if err != nil {
		log.Fatalln(err)
	}
	opOption := etcdClientV3.WithLease(leaseGrantResponse.ID)
	_, err = receiver.Put(context.TODO(), receiver.GetServiceKey(serviceName), Host, opOption)
	if err != nil {
		log.Fatalln(err)
	}
	leaseKeepAliveResponseChan, err := receiver.KeepAlive(context.TODO(), leaseGrantResponse.ID)
	if err != nil {
		log.Fatalln(err)
	}
	count := 0
	for leaseKeepResp := range leaseKeepAliveResponseChan {
		count++
		log.Printf("服务: %s, 地址: %s 运行正常 key: %s leaseId: %d\n", serviceName, Host, receiver.GetServiceKey(serviceName), leaseKeepResp.ID)
	}
	go receiver.registerService(serviceName, Host)
	log.Printf("服务: %s, 地址: %s 关闭", serviceName, Host)
}

func (receiver *Etcd) GetServiceHost(serviceName string) (string, error) {
	getResponse, err := receiver.Get(context.TODO(), receiver.GetServiceKey(serviceName))
	if err != nil {
		return "", err
	}
	if len(getResponse.Kvs) <= 0 {
		return "", fmt.Errorf("获取服务 %s 的host为空", serviceName)
	}
	return string(getResponse.Kvs[0].Value), nil
}

func (receiver *Etcd) GetServiceKey(serviceName string) string {
	return ServiceDir + receiver.environment + "/" + serviceName
}
