package main

import "fmt"

// People 结构体
type People struct {
	name string
	age  int
}

// 定义结构体的方法
func (r *People) getName() string {
	return r.name
}

// 定义结构体的方法
func (r *People) setName(n string) {
	r.name = n
}

// 使用结构体匿名
type Student struct {
	People
	classNum string
}

func main() {
	// 创建学生实例
	s := Student{
		People: People{
			name: "张三",
			age:  18,
		},
		classNum: "03-02",
	}
	// 结构体可以直接调用匿名结构体中的方法或属性
	name := s.name
	fmt.Printf("直接调用属性name = %s\n", name)
	// 调用接头体中的方法，设置姓名
	s.setName("李四")
	// 获取get 方法
	n := s.getName()
	fmt.Printf("new name=%s\n", n)

}
