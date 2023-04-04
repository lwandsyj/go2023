package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 字符串reader
func StringReader() {
	// 创建strings.Reader
	reader := strings.NewReader("hello go!")
	// 读取reader 中的数据
	// err 不返回io.EOF
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("content=%s\n", string(content))
}

// 文件Reader
func FileReader() {
	// 文件路径
	filePath := "./hello.txt"
	// 打开文件，file 实现了io.Reader 接口
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 读取全部内容,io.ReadAll 接收实现了io.Reader 接口的实例
	content, rErr := io.ReadAll(file)
	if rErr != nil {
		fmt.Printf("readAll error:%v\n", rErr)
		return
	}
	fmt.Printf("content=%s\n", string(content))
}
func main() {
	FileReader()
}
