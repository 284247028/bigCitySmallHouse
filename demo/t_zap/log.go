package t_zap

import (
	"bigCitySmallHouse/config"
	"bigCitySmallHouse/util/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

var _once = sync.Once{}
var _logger *Logger

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{}
}

func GetLogger(opts *LoggerOption) (*Logger, error) {
	var err error
	_once.Do(func() {
		_logger = NewLogger()
		err = _logger.Init(opts)
	})
	return _logger, err
}

type LoggerOption struct {
	Category string
}

func (receiver *Logger) Init(opts *LoggerOption) error {
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

	tPath += opts.Category

	today := time.Now().Format("2006-01-02")

	errLogName := opts.Category + "-" + today + "-error.log"

	if !file.Exists(tPath) {
		err = os.MkdirAll(tPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	filePath := tPath + "/" + errLogName

	var f *os.File
	if !file.Exists(filePath) {
		f, err = os.Create(filePath)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
	}

	encoder = zapcore.NewJSONEncoder(encoderConf)

	core1 := zapcore.NewCore(encoder, zapcore.AddSync(f), zap.ErrorLevel)
	core := zapcore.NewTee(core1)
	logger = zap.New(core)

	if err != nil {
		return err
	}

	receiver.SugaredLogger = logger.Sugar()

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
