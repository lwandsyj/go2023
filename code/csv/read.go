package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// csv 文件路径
	filePath := "./files/user.csv"
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 创建csv Reader
	reader := csv.NewReader(file)
	for {
		rows, err := reader.Read()
		if err != nil {
			// 如果err==io.EOF 表示读取完毕
			if err == io.EOF {
				break
			}
			fmt.Printf("csv read error:%v\n", err)
			return
		}
		fmt.Println(rows)
	}
}
