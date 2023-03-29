package main

import "fmt"

// slice的指针
func test(a *[]int) {
	// 获取指针对应的原始值
	fmt.Printf("*a=%v\n", *a)
	// 获取其中的元素，
	// 这是slice指针，而不是指针类型slice，
	fmt.Printf("*a[0]=%d\n", (*a)[0])
}

// slice 类型指针，每一个元素都是一个指针
func test1(a []*int) {
	fmt.Printf("a[0]的原始值：%d\n", *a[0])
}

func main() {
	/* a := []int{1, 2, 3}
	test(&a) */
	a := 1
	b := 2
	c := []*int{&a, &b}
	test1(c)
}
