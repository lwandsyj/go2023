package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// csv 文件路径
	filePath := "./files/user.csv"
	// 要读取文件数据，首先要打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	defer file.Close()
	// 创建csv reader
	// file 实现了io.Reader 和 io.Write 接口
	reader := csv.NewReader(file)
	// 读取全部记录
	// := 只要左边有一个是新变量即可
	contents, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("reader csv error:%v\n", err)
		return
	}
	fmt.Println(contents)
}
