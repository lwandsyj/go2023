package main

import "fmt"

func main() {
	// 指针类型的指针
	a := 2
	// b 为指针类型
	b := &a
	// c 为指针类型的指针
	c := &b
	// a 的值
	fmt.Printf("a=%d\n", a)
	// b 为a 的指针
	fmt.Printf("b=%p\n", b)
	// *b 为b 指针对应的值
	fmt.Printf("*b=%d\n", *b)
	// c 为指向b指针的指针
	fmt.Printf("c=%p\n", c)
	// *c 为b 的指针
	fmt.Printf("*c=%p,b=%p\n", *c, b)
	// **c 为c 指向的原始值
	fmt.Printf("**c=%d\n", **c)
}
