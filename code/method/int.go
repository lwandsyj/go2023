package main

import "fmt"

// 别名不受限制，因此要扩展不在同一个包中的，可以使用别名

type Int int

// 方法要求接收者必须在同一个包里面
func (recv Int) add(x int) int {
	// recv 是类型的实例变量，
	// 虽然Int 是int 的别名，但是两者是不同的类型
	return int(recv) + x
}

func main() {
	var a Int = 2
	// 接收者相当于类型的实例
	// 通过类型的实例调用方法，类似于其他语言中的类的实例方法
	c := a.add(3)
	fmt.Printf("c=%d\n", c)
}
