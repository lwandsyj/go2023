package main

import (
	"encoding/json"
	"fmt"
)

// 将数组转成json 数组字符串
func ConvertArr() {
	// int 类型数组
	var a [3]int = [...]int{1, 2, 3}
	// 转成json 字符串
	s, err := json.Marshal(a) // 返回byte[] 数组

	if err != nil {
		fmt.Printf("object to json string error:%s\n", s)
		return
	}
	// 输出json 字符串
	fmt.Println(string(s))
}

// 将slice 转成json 数组字符串
func ConvertSlice() {
	var a []string = []string{"hello", "go", "world"}
	// 转成json 字符串
	s, err := json.Marshal(a) // 返回byte[] 数组

	if err != nil {
		fmt.Printf("object to json string error:%s\n", s)
		return
	}
	// 输出json 字符串
	fmt.Println(string(s))
}

func main() {
	ConvertSlice()
}
