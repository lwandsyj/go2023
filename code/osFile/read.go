package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件路径
	filePath := "./upload/a.txt"
	// 打开文件
	fileInfo, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	// 使用[]byte 读取数据，并储存在变量中
	a := make([]byte, 5)
	rtn := make([]byte, 5)
	for {
		n, err := fileInfo.Read(a)

		if err != nil {
			// 当err==io.EOF 时表示读取完毕
			if err != io.EOF {
				fmt.Println(err)
				return // 退出执行当前函数，for 后面的打印语句不会执行
			}
		}
		if n == 0 {
			fmt.Printf("a=%s\n", string(a)) // 虽然n==0 ,但是a 中仍有数据
			break                           // 结束for 循环
		} else {
			rtn = append(rtn, a...)
		}
	}
	fmt.Println(string(rtn))
	// 关闭
	fileInfo.Close()
}
