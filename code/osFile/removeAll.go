package main

import (
	"fmt"
	"os"
)

// 删除不是空的目录
func RemoveAll() {
	dirPath := "./upload"
	// 删除不是空的目录
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	fmt.Println("remove success")
}

func main() {
	RemoveAll()
}
