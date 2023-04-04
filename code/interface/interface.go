package main

import "fmt"

// 定义接口
type Reader interface {
	Read() string
}

// 在go 语言中只要实现了接口中定义的所有方法就可称为实现了接口
type People struct {
	name string
	age  int
}

// 实现接口中的方法
func (r People) Read() string {
	return fmt.Sprintf("name is %s,age is %d\n", r.name, r.age)
}

// 测试接口，参数变量为接口类型
func handleTestInterface(r Reader) {
	// 执行
	s := r.Read()
	fmt.Println(s)
}

func main() {
	// 实例化结构体
	p := People{
		name: "站三",
		age:  18,
	}
	handleTestInterface(p)
	handleTestInterface(&p)
}
