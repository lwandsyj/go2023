package main

import "fmt"

func test1(a *[]int) {
	fmt.Printf("test a 的地址:%p\n", a)       // a 为指针地址类型，0x14000114030
	*a = []int{2, 3, 4, 5}                 // 赋值重新分配值，0x14000130000
	fmt.Printf("test rtn a 值的地址:%p\n", *a) // 赋值重新分配值，0x14000130000
	// a 指针的0x14000130000
	fmt.Printf("test rtn a 的地址:%p\n", a)
}

func test2(a *[]int) {
	fmt.Printf("test a 的地址:%p\n", a)
	// 不在重新赋值，只改变其中某个元素的值
	//*a = []int{2, 3, 4, 5}
	(*a)[0] = 111
	fmt.Printf("test rtn a 值的地址:%p\n", *a)
	// a 指针的0x14000130000
	fmt.Printf("test rtn a 的地址:%p\n", a)
}

func main() {
	/* var a []int
	fmt.Printf("zero a 中值的地址:%p\n", a) //0x0 因为没有初始化为[]
	fmt.Printf("zero a 的地址:%p\n", &a)  // a 变量的地址，0x14000114030
	test1(&a)
	fmt.Printf("zero a 中值的地址:%p\n", a)  //  赋值重新分配值，0x0
	fmt.Printf("a=%v\n", a)             //a=[2 3 4 5]
	fmt.Printf("zero a 中值的地址:%p\n", &a) //  赋值重新分配值，0x14000114030 */

	b := []int{1, 2, 3}
	fmt.Printf("zero b 值的地址:%p\n", b)
	fmt.Printf("zero b 的地址:%p\n", &b)
	test2(&b)
	fmt.Printf("b=%v\n", b)
}
