package main

import "fmt"

// 定义函数返回值
func addNum(a, b int) int {
	return a + b
}

/* func addNum1(a, b int) (int, int) {
	a += 1
	b += 2
	return a, b
} */

func addNum1(a, b int) int {
	a += 1
	b += 2
	return a, b
}
func main() {
	a := 2
	b := 3
	c := addNum(a, b)
	fmt.Println(c)
	d, f := addNum1(a, b)
	fmt.Println(d)
	fmt.Println(f)
}
