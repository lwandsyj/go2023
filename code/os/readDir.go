package main

import (
	"fmt"
	"os"
)

func main() {
	dirPath := "./upload"

	// 读取目录内容
	infos, err := os.ReadDir(dirPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	for index, info := range infos {
		fmt.Printf("index=%d,info=%+v\n", index, info)
		fmt.Printf("name=%s\n", info.Name())
		fmt.Printf("isDir=%t\n", info.IsDir())
		fmt.Printf("type=%+v\n", info.Type())
		//fmt.Printf("type=%+v\n",info.Info())
		fmt.Println("....")
	}
}
