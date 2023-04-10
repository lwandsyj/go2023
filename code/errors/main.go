package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建error 对象
	err := errors.New("自定义error")
	// 返回错误
	fmt.Println(err.Error())
}
