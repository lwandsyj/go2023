package main

import (
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

func main() {
	StatDirInfo()
}
