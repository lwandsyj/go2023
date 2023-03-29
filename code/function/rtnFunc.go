package main

import "fmt"

// 返回函数类型
func test() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	// 函数赋值给变量
	a := test()
	// 执行返回的函数
	c := a(2, 3)
	// 输出返回值
	fmt.Println(c)
}
