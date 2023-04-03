package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a string = `
	{
		"name":"张三",
		"age":18,
 		"sex":true,
		"birthday":"1995-01-01"
	}`
	// 解析成map
	var b map[string]any = make(map[string]any, 6)
	// 解析json
	err := json.Unmarshal([]byte(a), &b)
	if err != nil {
		fmt.Printf("error=%v\n", err)
		return
	}
	fmt.Printf("b=%+v\n", b)
	// 遍历map
	for key, value := range b {
		fmt.Printf("key=%s,value=%v\n", key, value)
	}
}
