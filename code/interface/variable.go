package main

import (
	"fmt"
	"math"
)

// 定义接口
type Abser interface {
	abs() float64
}

// 方法定义要求接收者类型和方法必须在一个package 里面
// 但是可以使用别名定义类型，不受这以限制，因为别名虽然和类型
// 是同一种类型方法存储数据，但是他们是不同的类型
type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		// MyFloat 和 float 不是同一类型
		// 需要转换
		return -float64(f)
	}
	return float64(f)
}

// 定义接口提
type Vertex struct {
	X, Y float64
}

// 接收者为指针类型
func (v *Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	// 定义接口类型的变量
	var a Abser
	var f MyFloat = -2.3
	// 因为MyFloat 实现了接口，因此可以赋值给a
	a = f
	s := Vertex{10, 10}
	a = &s // 因为接收变量为指针类型，所以认为指针类型实现了接口，因此s 可以赋值给接口类型变量
	//a=s // 错误，因为接收者是指针类型，被认为值变量没有实现接口
	fmt.Println(a)

}
