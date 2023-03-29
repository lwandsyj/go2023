package main

import "fmt"

func test(a []int) {
	fmt.Printf("test a=%p\n", a)
	// 如果是零值nil, 则初始化slice
	if a == nil {
		// make([]type,len,cap)
		// 当len = 0 时，不能使用a[0]
		a = make([]int, 1, 10)
		fmt.Printf("test new a=%p\n", a)
	}
	a[0] = 11
	a = []int{2, 3, 3}
	fmt.Printf("test a=%v\n", a)
}
func test1(a *[]int) {
	if a == nil {
		// 分配地址
		a = new([]int)
	}
	*a = []int{2, 3, 4, 5}
}

func main() {
	var a []int
	fmt.Printf("zero a=%v\n", a)
	b := []int{1, 2, 3}
	test(a)
	fmt.Printf("a=%v\n", a)
	test1(&a)
	fmt.Printf("a=%v\n", a)
	test(b)
	fmt.Printf("b=%v\n", b)
	test1(&b)
	fmt.Printf("b=%v\n", b)
}
