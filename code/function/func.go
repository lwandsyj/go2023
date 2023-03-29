package main

import "fmt"

func addNum(a int, b int) {
	c := a + b
	fmt.Println(c)
}

// 相邻参数类型相同，可以使用简写
func addNum1(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	a := 2
	b := 3
	addNum1(a, b)
}
