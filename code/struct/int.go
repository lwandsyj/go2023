package main

import "fmt"

type Mystrcut struct {
	// 匿名类型，字段名为类型名称
	int
}

func main() {
	// 创建结构指针类型
	my := new(Mystrcut)
	fmt.Printf("%+v\n", *my)

	// 使用匿名类型
	my.int = 80
	fmt.Println(my.int)
}
