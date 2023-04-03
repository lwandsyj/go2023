package main

import (
	"bufio"
	"fmt"
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
	// 创建scanner
	scanner := bufio.NewScanner(file)

	// 逐行读取文本中的内容，当到结尾时，返回false
	for scanner.Scan() {
		// scanner.Text 返回扫描到的文本
		fmt.Println(scanner.Text())
	}

	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Printf("scanErr:%v\n", scanErr)
	}

	// 关闭文件
	file.Close()
}
