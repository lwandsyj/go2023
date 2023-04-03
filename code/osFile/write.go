package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件路径，文件是一个空文件，没有数据
	filePath := "./upload/a.txt"
	// os.O_CREATE: 如果文件路径不存在，则创建文件
	// os.O_APPEND: 再文件最后面追加写入的数据
	// os.O_RDWR: 读取和写入
	info, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0555)
	// 如果有错误
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	// 要写入的数据
	txt := "hello go!"
	n, wErr := info.Write([]byte(txt))
	if wErr != nil {
		fmt.Printf("write file error:%v\n", wErr)
		return
	}
	fmt.Printf("write byte = %d\n", n)
	// 关闭文件
	info.Close()
}
