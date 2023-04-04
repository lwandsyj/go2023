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
func (r *People) Read() string {
	return fmt.Sprintf("name is %s,age is %d\n", r.name, r.age)
}

type Student struct {
	// 接口匿名字段
	Reader
	num int
}

// 重写read 方法
func (r *Student) Read() string {
	return fmt.Sprintf("%s,num is %d\n", r.Reader.Read(), r.num)
}

func main() {
	s := Student{&People{name: "lisi", age: 18}, 1234}
	fmt.Println(s.Read())
}
