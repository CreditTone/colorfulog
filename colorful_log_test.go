package log

import (
	"testing"
)

func TestColorful(t *testing.T) {
	//如果不想输出到指定文件，注释此函数即可
	SetOutputfilename("tmp.log")
	Info("普通日志，颜色为白色!")
	Warn("警告日志，颜色为黄色!")
	Error("错误日志，颜色为红色!")
}
