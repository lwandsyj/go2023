package main

import "fmt"

func main() {
	a := 10
	// if 可以有一个子语句，初始化局部变量，用; 分号分割
	if a = 20; a < 30 {
		fmt.Println(true)
	}
	fmt.Println(a)
	// 子语句定义和初始化变量，用;分号分割
	if b := 30; b < 40 {
		fmt.Println("b<40")
	}
	// 在if 子语句中声明的变量，不能在外面访问
	//fmt.Println(b)
	if c, d := 30, 10; c < d {

	}
}
