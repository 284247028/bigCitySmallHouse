package config

import (
	"flag"
	"fmt"
	"path"
	"runtime"
	"sync"
)

var _once sync.Once
var _conf *Config

const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvProduct     = "product"
)

type Config struct {
	Env string // 当前所处的环境
}

func NewConfig() *Config {
	return &Config{}
}

func GetConfig() (*Config, error) {
	var err error = nil
	_once.Do(func() {
		_conf = NewConfig()
		_conf.Init()
		err = _conf.Check()
	})
	return _conf, err
}

func (receiver *Config) Init() {
	flag.StringVar(&receiver.Env, "env", EnvDevelopment, "部署环境")
	flag.Parse()
}

func (receiver *Config) Check() error {
	switch receiver.Env {
	case EnvDevelopment, EnvTest, EnvProduct:
		return nil
	default:
		return fmt.Errorf("错误的环境：%s", receiver.Env)
	}
}

func CurrencyPath() (string, error) {
	_, tPath, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("获取运行时路径失败")
	}
	tPath = path.Dir(tPath)
	return tPath, nil
}

func RootPath() (string, error) {
	currencyPath, err := CurrencyPath()
	if err != nil {
		return "", err
	}
	currencyPath += "/../"
	return currencyPath, nil
}
