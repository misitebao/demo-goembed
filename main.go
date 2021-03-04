package main

import (
	"embed"
	_ "embed"
	"fmt"
	"time"
)

//go:embed demo.txt
var b []byte

//go:embed hello.txt
var s string

// []byte 和 string 只支持单个文件

//go:embed template
var t embed.FS

func main() {

	fmt.Printf("字节数组表示为%#v\n", b)
	fmt.Printf("转换为字符串表示为%#v\n", string(b))

	fmt.Printf("字节数组表示为%#v\n", s)

	data, err := t.ReadFile("template/index.html")

	if err != nil {
		fmt.Println("读取文件错误")
	}

	fmt.Printf("读取到的字符串为%#v\n\n", string(data))

	data2, _ := t.ReadFile("template/about.html")

	fmt.Printf("读取到的字符串为%#v\n", string(data2))

	// 为了不让程序运行完就退出
	time.Sleep(time.Second * 5)

	// 单独一个程序  即时移动文件也不会影响，就说明静态文件已经打包进去了

	// 这个视频是使用笔记本录制的 麦克风有点不好用，所以暂时用音乐代替了，稍后会上传b站  用户名就是  米司特包

	// 另外 所有的代码和示例说明会同步更新到我的博客中  百度 搜索  “米司博客”第一个就是

	好了 本期视频就到这里了 后面我会优化视频录制和剪辑 这期就暂且这样了，后续有时间再重制，Bye

}
