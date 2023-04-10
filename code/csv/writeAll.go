package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// 写入一个不存在的文件
func WriteNotExist() {
	// 要写入的内容
	contents := [][]string{
		{"id", "name", "sex", "age"},
		{"1", "张三", "1", "18"},
		{"2", "李四", "1", "19"},
	}
	// 要写入的文件
	filePath := "./files/student.csv"
	// 以写的方式打开文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	defer file.Close()
	// 创建写csv.Writer
	// file 既实例了io.Reader, 亦实现了io.Writer 接口
	writer := csv.NewWriter(file)
	// 写入文件
	wErr := writer.WriteAll(contents)
	if wErr != nil {
		fmt.Printf("write data to csv error:%v\n", wErr)
		return
	}
	fmt.Println("write success")
}

// 已追加的形式写入数据
func AppendCsv() {
	filePath := "./files/student.csv"
	// 打开文件,要有读写的权限和追加的权限
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	// 关闭
	defer file.Close()
	content := [][]string{
		{"3", "王娇", "0", "18"},
	}
	// 创建Writer
	writer := csv.NewWriter(file)
	wErr := writer.WriteAll(content)
	if wErr != nil {
		fmt.Printf("write data to csv error:%v\n", wErr)
		return
	}
	fmt.Println("write success")
}

func main() {
	WriteNotExist()
}
