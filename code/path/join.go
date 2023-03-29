package main

import (
	"fmt"
	"path"
)

func main() {
	dir := "../vest"
	filename := "doc.md"

	// path.Join 合并多个路径
	filePath := path.Join(dir, filename)

	fmt.Println(filePath)
}
