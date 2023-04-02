package main

import (
	"fmt"
	"os"
)

func main() {
	// 目录路径
	dirPath := "a/b/c"

	err := os.MkdirAll(dirPath, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("创建目录和子目录成功")
}
