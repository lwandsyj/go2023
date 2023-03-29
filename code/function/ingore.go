package main

import "fmt"

func get() (int, int) {
	return 2, 3
}

func main() {

	// 使用_ 忽略返回值的第二个数
	a, _ := get()

	fmt.Println(a)
}
