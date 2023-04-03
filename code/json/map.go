package main

import (
	"encoding/json"
	"fmt"
)

// 将map 转成json 中对象格式{}
func ConvertMap() {
	var a map[string]any = map[string]any{
		"name": "张三",
		"age":  18,
	}
	// 将map 转成json 字符串
	b, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
		return
	}
	// 输出字符串
	fmt.Println(string(b))
}

func main() {
	ConvertMap()
}
