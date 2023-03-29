package main

import (
	"path"
)

// 定义接口
type IPeople interface {
	// 获取名称
	getName() string
	// 设置新的名称，没有返回值
	setName(string)
	// 设置新的年龄
	// newAge 新的年龄
	setAge(newAge int)
	// 获取年龄
	getAge() int
}

func main() {
	path.Join()
}
