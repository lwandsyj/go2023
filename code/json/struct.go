package main

import (
	"encoding/json"
	"fmt"
)

// people struct
type People struct {
	name string
	age  int
}

func ConvertStruct() {
	p := People{
		name: "张三",
		age:  18,
	}
	// struct 转成json {}
	o, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(o))
}

type Student struct {
	CardNum    int    `json:"-" ` // - 始终忽略
	Name       string `json:"name"`
	Age        int    `json:"age,omitempty"` // 0 或者空时，false 等忽略
	ClassesNum string `json:"classNum"`
}

func ConvertStudent() {
	s := Student{
		ClassesNum: "03-01",
		Name:       "黎明",
		Age:        0,
		CardNum:    1234,
	}
	o, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(o))
}

func main() {
	ConvertStudent()
}
