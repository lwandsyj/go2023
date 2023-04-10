package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件路径
	filePath := "./upload/hello.txt"
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	// 使用defer
	// 文件关闭
	defer file.Close()
	// 查看文件状态
	state, err1 := os.Stat(filePath)
	if err1 != nil {
		fmt.Printf("state error:%v\n", err1)
		return
	}
	// 文件字节长度
	size := state.Size()
	// 创建byte slice，存取读取的文件数据
	content := make([]byte, size)
	// 读取文件内容
	_, rErr := file.Read(content)
	if rErr != nil {
		if rErr != io.EOF {
			fmt.Printf("read error %v\n", rErr)
		}
	}
	// 字节类型转成string
	fmt.Println(string(content))
}
