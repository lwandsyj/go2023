package main

import (
	"flag"
	"fmt"
)

func main() {
	var b bool
	// 使用Flag.BoolVar 传入变量指针的方式接受参数
	flag.BoolVar(&b, "flag", false, "usage")

	// 收集参数
	flag.Parse()
	// 不用再解指针
	fmt.Println(b)
}
