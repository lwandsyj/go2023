package main

import "fmt"

func test1(a *[]int) {
	fmt.Printf("test a 的地址:%p\n", a) // a 为指针地址类型，0x14000196018
	if len(*a) == 0 {
		// 重新初始化指针地址，分配地址
		a = new([]int) // 指针地址指向的变量值由0x0改为新的地址0x140000b8000
		fmt.Printf("test new a 存的值的地址:%p\n", *a)
		fmt.Printf("test new a 的地址:%p\n", a) // 变量的地址改变0x14000196030
	}
	*a = []int{2, 3, 4, 5}                 // 赋值重新分配值，0x140001a2000
	fmt.Printf("test rtn a 值的地址:%p\n", *a) // 赋值重新分配值，0x140001a2000
	// a 指针的0x14000196030, 因为指针使用new 重新分配了地址
	fmt.Printf("test rtn a 的地址:%p\n", a)
}

func main() {
	var a []int
	fmt.Printf("zero a 中值的地址:%p\n", a) //0x0 因为没有初始化为[]
	fmt.Printf("zero a 的地址:%p\n", &a)  // a 变量的地址，0x14000196018
	test1(&a)
	fmt.Printf("zero a 中值的地址:%p\n", a) //  赋值重新分配值，0x0
	fmt.Printf("a=%v\n", a)
	fmt.Printf("zero a 中值的地址:%p\n", &a) //  赋值重新分配值，0x14000196018

	/*
		b := []int{1, 2, 3}
		fmt.Printf("zero b 的地址:%p\n", a)
		test1(&b)
		fmt.Printf("b=%v\n", b) */
}
