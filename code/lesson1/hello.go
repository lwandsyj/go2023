package main

import (
	"fmt"
)

/* var name ="hello"
var age uint
var sex int =1 */
// 如果没有初始值，则必须定义类型
//var classes;

/* var (
	name string
	age  uint
	sex  int
) */

/* var name, age, sex = "", 1, 12
var a, b, c int */

func main() {
	/* a, b := 2, 3
	a, c := b, a */
	a := 3
	// p 是指向a 的指针，是a 的地址，他的类型是*int
	p := &a

	fmt.Printf("hellog go a=%d,b=%p\n", a, p)
	// *p 是变量a 的值，是变量a 的别名，可以读取和更新变量a
	*p = 5
	fmt.Printf("hellog go a=%d,b=%p\n", a, p)

	// 指针的类型为*type 类型，零值为nil
	var z *int

	if z == nil {
		fmt.Printf("指针的零值为nil,z=%p\n", z)
	}
	// 指针类型赋值,*z 表示取当前地址的值，而当前地址为nil 地址，尚未分配内存，不能给nil 地址赋值
	// 给z 指针分配新的指针
	z = &a
	if z != nil {
		fmt.Printf("指针的零值为nil,z=%p,*z=%d\n", z, *z)
	}

	var y *string
	// 使用new 给指针分配一块新的内存，不在是nil
	y = new(string)
	*y = "hello"
	fmt.Printf("指针y=%p,*y=%s", y, *y)
}
