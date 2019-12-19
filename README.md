# colorfulog


		colorfulog是一个简单且易于扩展的日志框架，其本身提供的功能非常简单，但是很实用。程序员查找日志最重要两个信息，日志输出的文件和行号。另外如果有输出颜色区分那最好不过了。无奈市面上所谓的大名鼎鼎的日志框架logrus、seelog自以为封装的极好。但是任然get不到程序员最痛苦的点。别怪我喷你俩，喷就喷了。能把我咋地？日志快速定位问题才是重要的。搞那么复杂有屌用啊！


## 快速开始

	go get -u github.com/CreditTone/colorfulog

##使用日志输出
```golang
import(
	"github.com/CreditTone/colorfulog"
)

func main() {
	colorfulog.Info("普通日志，颜色为白色!")
	colorfulog.Warn("警告日志，颜色为黄色!")
	colorfulog.Error("错误日志，颜色为红色!")
}
```
##输出效果如下

