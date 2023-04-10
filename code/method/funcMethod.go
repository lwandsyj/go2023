package main

import "fmt"

// 定义结构体
type People struct {
	name string
	age  int
}

// 方法
func (r *People) getName() string {
	return r.name
}

// 函数
func getName(p People) string {
	return p.name
}

// 函数接收指针类型
func getName1(p *People) string {
	return p.name
}

func main() {
	p := People{
		name: "站三",
		age:  18,
	}
	// 值变量类型调用方法
	fmt.Println(p.getName())
	pointer := &p
	// 指针类型调用方法
	fmt.Println(pointer.getName())
	// 函数
	getName(p)
	getName1(pointer)
	getName1(p)
}
