package main

import "fmt"

func main() {
	b := map[string]any{
		"name": "zhangsan",
		"age":  18,
	}
	fmt.Printf("b 存的值的地址:%p\n", b)
	fmt.Printf("b 的地址：%p\n", &b)
}
