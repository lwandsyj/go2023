package main

import (
	"fmt"
	"os"
)

// 删除文件
func RemoveFile() {
	// 不存在的文件路径
	filePath := "./upload/d.txt"
	// 删除文件
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("remove not exist file error:%v\n", err)
	}
	// 删除已存在的文件路径
	filePath1 := "./upload/c.txt"
	// 删除文件
	rErr := os.Remove(filePath1)
	if rErr != nil {
		fmt.Printf("remove exist file error:%v\n", rErr)
		return
	}
	fmt.Println("remove success")
}

// 删除目录
func RemoveDir() {
	// 不存在的目录路径
	dirPath := "./upload1"
	// 删除不存在的目录
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Printf("remove not exist dir error:%v\n", err)
	}
	// 存在的目录，但是不是空目录
	dirPath1 := "./upload"
	err1 := os.Remove(dirPath1)
	if err1 != nil {
		fmt.Printf("remove not empty dir error:%v\n", err1)
	}
	// 存在的空目录
	dirPath2 := "./ooo"
	err2 := os.Remove(dirPath2)
	if err2 != nil {
		fmt.Printf("remove  empty dir error:%v\n", err2)
		return
	}
	fmt.Println("remove success")
}
func main() {
	RemoveDir()
}
