package main

import "fmt"

func main() {
	a := 0
	b := 5
	// 在go 语言中循环只有 for语句
	// for 可以省略初始化子语句和后置子语句，这种模拟while
	for a < b {
		a++
		fmt.Printf("a =%d\n", a)
	}
}
