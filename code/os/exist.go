package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// 存在的路径
	dirPath := "./upload"
	_, err := os.Stat(dirPath)
	// 如果err == nil ,表示给定的路径存在
	if os.IsNotExist(err) {
		fmt.Println(err)
	}
	// 不存在的路径
	dirPath1 := "./upload1"
	_, err1 := os.Stat(dirPath1)
	// err1 != nil 表示给定的路径不存在
	fmt.Println(err1)
	if os.IsNotExist(err1) {
		fmt.Println(err1)
	}
	// 使用os 中定义的错误常量
	if errors.Is(err1, os.ErrNotExist) {
		fmt.Println("error constants:")
		fmt.Println(err1)
	}
}
