package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1"
	b := 3
	// 数字转字符串
	a = strconv.Itoa(b)
	fmt.Printf("a=%s\n", a)
	// 字符串转换为数字，返回(int,error)
	c, _ := strconv.Atoi(a)
	fmt.Printf("c=%d\n", c)
}
