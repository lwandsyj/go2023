package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件路径
	var filePath string = "./upload/a.txt"
	// 读取文件内容
	var content []byte

	// os.ReadFile 一次性读取文件所有内容
	content, err := os.ReadFile(filePath)

	if err != nil {
		// err == io.EOF 表示读取完毕
		if err != io.EOF {
			fmt.Printf("readFile error:%v\n", err)
			return
		}
	}
	fmt.Printf("content is:%s\n", string(content))
}
