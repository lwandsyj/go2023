package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 要写入的文件数据
	content := [][]string{
		{"4", "赵六", "1", "17"},
		{"5", "生气", "0", "18"},
	}
	// 文件路径
	filePath := "./files/student.csv"
	// 打开文件追加
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	defer file.Close()
	// 创建csv writer
	writer := csv.NewWriter(file)
	for _, elem := range content {
		// 单行写入,写入被缓冲(buffered)，
		// 因此最终必须调用 Flush 以确保将记录写入底层 io.Writer。
		err = writer.Write(elem)
		if err != nil {
			fmt.Printf("writer error:%v", err)
			return
		}
	}
	//  flush, 不调用flush 不会真正写入数据
	writer.Flush()
	if writer.Error() != nil {
		fmt.Printf("flush error:%v\n", writer.Error())
		return
	}
	fmt.Println("write success")
}
