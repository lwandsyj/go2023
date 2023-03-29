package main

import "fmt"

type People struct {
	name string
	age  int
}

// 这里是通过结构的值作为接收者
func (recv People) changeName(name string) {
	recv.name = name
}

// 要想改变接收者的值，使用指针类型
func (recv *People) changeName1(name string) {
	recv.name = name
}
func main() {
	a := People{"张三", 20}
	a.changeName1("李四")
	fmt.Printf("change 以后的结构：%v\n", a)

	// 结构体指针类型同值类型调用方法一样，不需要*
	b := &a
	b.changeName1("王五")
	fmt.Printf("pointer change  以后的结构：%v\n", a)
	fmt.Printf("pointer change  以后的结构：%v\n", b)
}
