package main

import (
	"fmt"
	"os"
)

// 删除非空目录
func RemoveNotEmptyDir() {
	dirPath1 := "./upload"
	// 使用os.Remove 删除文件或者空目录
	err := os.Remove(dirPath1)

	if err != nil {
		fmt.Println(err) // Directory not empty
		return
	}
}

// 删除空目录
func RemoveEmptyDir() {
	dirPath1 := "./upload1"
	// 使用os.Remove 删除文件或者空目录
	err := os.Remove(dirPath1)

	if err != nil {
		fmt.Println(err) //
		return
	}
	fmt.Println("删除成功")
}

// 删除文件
func RemoveFile() {
	filePath := "./upload/hello.txt"
	// 使用os.Remove 删除文件或者空目录
	err := os.Remove(filePath)

	if err != nil {
		fmt.Println(err) //
		return
	}
	fmt.Println("删除成功")

}
func main() {
	RemoveFile()

}
