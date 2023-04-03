package main

import (
	"fmt"
	"os"
	"path"
)

// 创建文件
func CreateFile(filePath string) {
	// 判断路径是否为空
	if filePath == "" {
		return
	}
	// 判断路径是否存在
	_, err := os.Stat(filePath)
	// 判断是否存在
	if !os.IsNotExist(err) {
		fmt.Println("this file is exist")
		return
	}
	// 获取文件后缀名
	ext := path.Ext(filePath)
	if len(ext) == 0 {
		fmt.Println("this is not a file")
		return
	}
	// 创建文件
	_, createErr := os.Create(filePath)
	if createErr != nil {
		fmt.Println(createErr)
		return
	}
	fmt.Println("create file success")
}
func main() {
	filePath := "./upload/a.txt"
	CreateFile(filePath)
	/* path2 := "./upload/hello.txt" // 已存在的文件路径
	CreateFile(path2)
	fielPath := "./upload1" // 不存在的文件路径
	CreateFile(fielPath) */
	/* // 目录路径
	path1 := "./upload"
	CreateFile(path1)
	*/

}
