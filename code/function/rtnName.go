package main

import "fmt"

// 使用命名返回值
func addNum(x, y int) (total int) {
	// total 返回值变量也是局部变量
	total = x + y
	return // 最后return 可以省略返回值，会自动返回total的值
}
func main() {
	x := 2
	y := 3
	c := addNum(x, y)
	fmt.Println(c)
}
