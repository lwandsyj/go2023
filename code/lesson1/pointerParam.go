package main

import "fmt"

/*
* 交换x和 y的值
* x 和 y 存放的数变量指针，即变量的存储位置
 */
func swap(x, y *int) {
	var tmp int
	// *x 获取指针位置的原始值
	tmp = *x
	// 将y 的原始值赋值给x 的原始值
	*x = *y
	// 将x 的原始值赋值费y 的原始值
	*y = tmp
}

func main() {
	x := 2
	y := 3
	fmt.Printf("交换前x=%d,y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("交换后x=%d,y=%d\n", x, y)
}
