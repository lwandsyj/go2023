package main

import (
	"fmt"
	"os"
)

// 使用RemoveAll 删除带有文件的目录
func RemoveAll() {
	dirPath := "./upload1"
	// 删除
	if err := os.RemoveAll(dirPath); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("删除成功")
	}
}

func main() {
	RemoveAll()
}
