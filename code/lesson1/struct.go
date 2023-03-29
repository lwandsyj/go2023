package main

import "fmt"

// 使用type 和 struct 关键字定义struct
// struct 定义一组属性对象
type People struct {
	// 字段 type
	name string
	age  int
}

func main() {
	// 初始化struct,
	// 这种初始化字段是按照字段定义的顺序赋值的，
	var a People = People{"张三", 28}
	// 下面是错误的，类型对应不上
	//var c People = People{28, "张三"}
	fmt.Printf("a=%v\n", a)
	// 没有位置顺序
	b := People{
		name: "李四",
		age:  20, // 最后的, 逗号不能省略
	}
	fmt.Printf("b=%v\n", b)

	// 指针
	var c *People
	c = new(People)
	c.name = "王五"
	fmt.Printf("c=%v\n", c)
	fmt.Printf("c=%v\n", *c)
	fmt.Printf("c=%v\n", c.name)

	var d People
	fmt.Println(d)
	// 取地址实例化
	var e *People = &People{
		name: "zhangsan",
		age:  12,
	}
	fmt.Printf("e=%p\n", e)

	// 没有初始化任意字段，为零值
	var f People = People{}
	fmt.Printf("f=%v", f)
	// 只初始化name
	var g People = People{name: "张三"}
	fmt.Printf("g=%v", g)
	// 只初始化年龄
	var h People = People{age: 12}
	fmt.Printf("h=%v", h)
	// 全部初始化
	var i People = People{
		age:  12,
		name: "lisi"}
	fmt.Printf("i=%v", i)
}
