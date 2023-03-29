package main

import "fmt"

// *type 指针类型，传递是地址
func addNum(a, b *int) {
	// 要使用原始值，需要先解指针*a
	c := *a + *b
	fmt.Println(c)
}
func main() {
	a := 2
	b := 3
	addNum(&a, &b)
}
