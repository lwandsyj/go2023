package main

import (
	"fmt"
	"strconv"
)

// 定义结构体
type People struct {
	name string
	age  int
}

func (recv People) toString() string {
	// recv 为接收者对象，即调用该方法的对象，本例中为p := People{"张三", 18}
	return recv.name + "'s age is " + strconv.Itoa(recv.age)
}

func main() {
	p := People{"张三", 18}
	// 通过p 的实例调用定义的方法
	c := p.toString()
	fmt.Println(c)
}
