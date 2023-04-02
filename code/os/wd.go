package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前目录
	dir, err := os.Getwd()
	if err == nil {
		fmt.Println(dir)
	}
}
