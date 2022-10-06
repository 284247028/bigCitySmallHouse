package logger

import (
	"bigCitySmallHouse/config"
	"bigCitySmallHouse/util/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var _logger *Logger

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{}
}

func GetLogger() *Logger {
	return _logger
}

type Options struct {
	Category string
}

func (receiver *Logger) Init(opts *Options) error {
	conf, err := config.GetConfig()
	if err != nil {
		return err
	}

	var logger *zap.Logger
	var encoderConf zapcore.EncoderConfig
	var encoder zapcore.Encoder
	//var writeSyncer zapcore.WriteSyncer

	switch conf.Env {
	case config.EnvDevelopment:
		encoderConf = zap.NewDevelopmentEncoderConfig()

	case config.EnvTest:
		encoderConf = zap.NewDevelopmentEncoderConfig()

	case config.EnvProduct:
		encoderConf = zap.NewProductionEncoderConfig()

	}

	tPath, err := GetLogsPath()
	if err != nil {
		return err
	}

	today := time.Now().Format("2006-01-02")

	tPath += opts.Category + "/" + today

	errLogName := opts.Category + "-" + today + "-error.log"
	//infoLogName := opts.Category + "-" + today + "-info.log"

	if !file.Exists(tPath) {
		err = os.MkdirAll(tPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	errPath := tPath + "/" + errLogName
	//infoPath := tPath + "/" + infoLogName

	var (
		errF *os.File
		//infoF *os.File
	)

	errF, err = os.OpenFile(errPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	//infoF, err = os.OpenFile(infoPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	//if err != nil {
	//	return err
	//}

	encoder = zapcore.NewJSONEncoder(encoderConf)

	core1 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
	//core2 := zapcore.NewCore(encoder, zapcore.AddSync(infoF), zap.InfoLevel)
	core := zapcore.NewTee(core1)
	logger = zap.New(core)

	receiver.SugaredLogger = logger.Sugar()

	_logger = receiver

	return nil
}

func GetLogsPath() (string, error) {
	tPath, err := config.CurrencyPath()
	if err != nil {
		return "", err
	}
	logsPath := tPath + "/../logs/"
	return logsPath, nil
}
