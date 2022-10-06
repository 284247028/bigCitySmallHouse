package t_zap

import (
	"testing"
)

func TestExample(t *testing.T) {
	logger, err := GetLogger(&LoggerOption{Category: "house"})
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("da wdu ih asi dnwdda wdu ih asi dnwdda wdu ih asi dnwdda wdu ih asi dnwdda wdu ih asi dnwd")
	logger.Error("测试的消息测试的消息测试的消息测试的消息测试的消息测试的消息")
	err = logger.Sync()
	if err != nil {
		t.Fatal(err)
	}
}
