package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// 文件路径
	fp := "./upload/hello.txt"
	// 读取文件信息
	info, err := os.Stat(fp)
	// 如果错误是路径指定的文件不存在
	if err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Println(err)
		return
	}
	fmt.Printf("info=%+v", info)

	fp1 := "./upload/hello1.txt"
	_, err1 := os.Stat(fp1)

	if os.IsNotExist(err1) {
		fmt.Printf("not exist,%v\n", err1)
	}

}
