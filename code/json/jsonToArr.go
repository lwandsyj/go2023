package main

import (
	"encoding/json"
	"fmt"
)

// json 字符串解析成数组
func ConvertToArr() {
	a := "[1,2,3,4]"
	var b []int = make([]int, 2)
	// json 反序列化, 第二个参数传递指针
	err := json.Unmarshal([]byte(a), &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	// 遍历slice
	for index, value := range b {
		fmt.Printf("index:=%d,value=%d\n", index, value)
	}
}

func main() {
	ConvertToArr()
}
