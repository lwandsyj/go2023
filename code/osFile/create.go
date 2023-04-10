package main

import (
	"fmt"
	"os"
)

// 使用os.Create 创建或者写入文件
// Create 当文件不存在会创建文件
func Create() {
	// 不存在的文件路径
	filePath := "./upload/c.txt"
	// 使用os.Create 创建不存在的文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create open file error:%v\n", err)
	}
	// 关闭文件
	defer file.Close()
	// 要写入文件的内容
	content := "要写入文件的内容"
	// 写入文件
	_, wErr := file.Write([]byte(content))
	if wErr != nil {
		fmt.Printf("write error:%v\n", wErr)
		return
	}
	fmt.Println("write success")
}

// 已存在的文件，比如上面函数创造的文件
// 会清除掉文件中已有的内容
func CreateExist() {
	// 以存在的文件路径
	filePath := "./upload/c.txt" // 通过上面的函数创建的文件，内容为：要写入文件的内容
	// 使用os.Create 创建不存在的文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create open file error:%v\n", err)
	}
	// 关闭文件
	defer file.Close()
	// 要写入文件的内容
	content := "hello go"
	// 写入文件
	_, wErr := file.Write([]byte(content))
	if wErr != nil {
		fmt.Printf("write error:%v\n", wErr)
		return
	}
	fmt.Println("write success")
}
func main() {
	CreateExist()
}
