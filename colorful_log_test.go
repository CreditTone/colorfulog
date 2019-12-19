package log

import (
	"testing"
)

func TestColorful(t *testing.T) {
	Info("普通日志，颜色为白色!")
	Warn("警告日志，颜色为黄色!")
	Error("错误日志，颜色为红色!")
}
