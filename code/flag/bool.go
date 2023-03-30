package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.Parse() 收集命令行参数，必须在使用命令行参数执行调用
	//flag.Parse()
	// 命令行参数名称为flag
	// 默认值为false, 假如没有传参数，就使用默认值
	// 标记该参数的用途
	b := flag.Bool("flag", false, "is use Pintf")

	// flag.Parse() 收集命令行参数，必须在使用命令行参数执行调用
	//flag.Parse()

	// 返回的是指针类型，因此要使用原始值，需要先解指针
	if *b {
		fmt.Printf("this arg's value is %t\n", *b)
	} else {
		fmt.Println(b)
	}
}
