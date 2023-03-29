package main

import "fmt"

func add(args ...int) {
	// 在函数体内部，可变参数是一个slice
	count := len(args)

	if count == 0 {
		fmt.Println("没有参数")
		return
	}
	// 变量args
	for index, value := range args {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}
}

func main() {
	// 可变参数可以不传值
	add()
	println()
	// 可变参数
	add(1, 2, 3)
	println()
	add(9, 8, 7, 6, 5)
}
