package logger

import (
	"testing"
)

func TestExample(t *testing.T) {
	logger := NewLogger()
	err := logger.Init(&Options{Category: "house"})
	if err != nil {
		t.Fatal(err)
	}
	GetLogger().Info("da wdu ih asi dnwdda wdu ih asi dnwdda wdu ih asi dnwdda wdu ih asi dnwdda wdu ih asi dnwd")
	GetLogger().Error("测试的消息测试的消息测试的消息测试的消息测试的消息测试的消息")
	err = logger.Sync()
	if err != nil {
		t.Fatal(err)
	}
}
