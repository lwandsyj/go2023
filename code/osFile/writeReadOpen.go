package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件路径
	filePath := "./upload/a.txt"
	// 以只读方式打开文件，然后尝试写入文件
	info, err := os.Open(filePath)
	// 检查是否有错误，如果有终止执行
	if err != nil {
		fmt.Println(err)
		return
	}
	// 要写入文件的txt
	txt := "hello go!"
	// FIleInfo.Write 写入，使用[]byte 类型
	n, wErr := info.Write([]byte(txt))
	// 写入错误
	if wErr != nil {
		fmt.Printf("write error:%v\n", wErr)
		return
	}
	// 写入的字节数
	fmt.Println(n)
}
