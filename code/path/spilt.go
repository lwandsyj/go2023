package main

import (
	"fmt"
	"path"
)

func main() {
	filepath := "../vest/doc.md"

	dir, fileName := path.Split(filepath)

	// 输出路径目录部分
	fmt.Printf("dir=%s\n", dir)

	// 输出文件部分
	fmt.Printf("file=%s\n", fileName)
}
