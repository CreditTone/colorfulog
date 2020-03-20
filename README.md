# colorfulog


colorfulog是一个简单且易于扩展的go语言日志类库，其本身提供的功能非常简单，但是很实用。程序员查找日志最重要两个信息，日志输出的文件和行号。那些乱七八糟的日志框架，又配置这又配置那，搞的自己很高大上的样子。日志这个东西简单的东西非要搞那么复杂。要花那么多时间去学个锦上添花，并没有什么卵用的东西，有时间我还有其他更重要的研究工作呢。不爽，自己又造了个轮子

## 快速开始

	go get -u github.com/CreditTone/colorfulog

## 使用
```golang
import(
	colorfulog "github.com/CreditTone/colorfulog"
)

func main() {
    //如果不想输出到指定文件，注释此函数即可
	SetOutputfilename("tmp.log")
	colorfulog.Info("普通日志，颜色为白色!")
	colorfulog.Warn("警告日志，颜色为黄色!")
	colorfulog.Error("错误日志，颜色为红色!")
}
```

## 输出效果如下
![测试效果图](https://raw.githubusercontent.com/CreditTone/staticfiles/master/1576747983685.jpg "测试效果图")
