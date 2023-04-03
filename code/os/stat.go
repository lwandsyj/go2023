package main

import (
	"errors"
	"fmt"
	"os"
)

// 查看文件信息
func StatFileInfo() {
	filepath := "./hello.txt"
	// 查看文件信息
	info, err := os.Stat(filepath)
	if err == nil {
		//fmt.Println(info)
		//v, _ := json.Marshal(info)
		fmt.Printf("fileInof:%v\n", info.IsDir())
		return
	}
	fmt.Println("error")
	fmt.Println(err)
}

// 查看目录信息
func StatDirInfo() {
	dirPath := "./upload"
	// 查看目录信息
	if info, err := os.Stat(dirPath); err == nil {
		fmt.Printf("dir Info:%+v\n", info)
	} else {
		fmt.Println(err)
	}

}

// 查看目录信息，如果路径不存在，info 返回nil
func StatNotExistPath() {
	dirPath := "./upload1"
	// 查看路径信息，当路径在项目中不存在使，info 返回nil
	info, err := os.Stat(dirPath)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(info)
	}
}

func main() {
	StatNotExistPath()
}
