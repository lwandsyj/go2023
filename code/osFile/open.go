package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	fp := "./upload/hello.txt"
	// Open 以只读的方式打开文件
	file, err := os.Open(fp)
	if err != nil {
		return
	}
	// 获取文件句柄
	fileId := file.Fd()
	fmt.Printf("fileId=%d\n", fileId)

	// 读取文件内容
}
