package main

import (
	"fmt"
	"strconv"
)

func main() {
	var b bool = false
	var a bool = true
	c := 1.2
	// strconv.FormatBool 把bool 类型转换成string 类型
	fmt.Printf("a=%s\n", strconv.FormatBool(a))
	fmt.Printf("b=%s\n", strconv.FormatBool(b))
	// f 格式化形式，ddd.dddd
	fmt.Printf("c=%s\n", strconv.FormatFloat(c, 'f', 1, 64))
	// e 对应形式d.ddddE±dd
	fmt.Printf("c=%s\n", strconv.FormatFloat(c, 'e', 1, 64))

}
