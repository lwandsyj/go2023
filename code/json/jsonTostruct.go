package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	CardNum    int    `json:"-" ` // - 始终忽略
	Name       string `json:"name"`
	Age        int    `json:"age,omitempty"` // 0 或者空时，false 等忽略
	ClassesNum string `json:"classNum"`
}

func main() {
	var a string = `
	  {
		"CardNum":12345,
		"name":"张三",
		"age":18,
		"classNum":"03-03"
	  }
	`
	// 解析成struct
	p := &Student{}
	err := json.Unmarshal([]byte(a), p)
	if err != nil {
		fmt.Printf("error = %v\n", err)
		return
	}
	fmt.Printf("struct = %+v\n", p)
}
