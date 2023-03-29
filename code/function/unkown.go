package main

import "fmt"

func main() {
	// 匿名函数赋值给变量
	a := func(x, y int) int { return x + y }
	// 执行函数
	c := a(2, 3)

	fmt.Println(c)
	fmt.Printf("a=%p\n", a)
	fmt.Printf("&a=%p\n", &a)

	b := "hello"
	// 定义并同时调用函数
	func() {
		b = "hello world"
	}()
	fmt.Println(b)

	// 带有参数的
	func(data string) {
		fmt.Printf("data = %s\n", b)
	}(b)
}
