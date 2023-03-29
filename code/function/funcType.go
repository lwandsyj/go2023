package main

import "fmt"

// 函数也是一种类型，使用type 定义函数的别名，可以使用简单语句代币函数签名
// 定义函数类型别名，使函数可以重复使用
// 函数签名
type AddNum func(int, int) int

/** 传入函数当做参数*/
func test(f AddNum) {
	a := 2
	b := 3
	c := f(a, b)
	fmt.Println(c)
}

// 定义f 为函数类型，这种方式不能重用，因此常用type 定义别名，简单并且复用
func test1(f func(int, int) int) {
	a := 2
	b := 3
	c := f(a, b)
	fmt.Println(c)
}

func addNum(x, y int) int {
	return x + y
}

func main() {
	//test(addNum)
	// 定义a 为AddNum 函数类型，只为addNum 具体函数
	var a AddNum = addNum
	test1(a)
}
