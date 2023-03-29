package main

import (
	"fmt"
	"path"
)

func main() {
	filePath := "../app/file.txt"
	// 返回文件名称,file.txt
	fmt.Println(path.Base(filePath))

	// 返回文件后缀名.txt
	fmt.Println(path.Ext(filePath))
}
