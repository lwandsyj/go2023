package main

import "fmt"

func main() {
	var x any = 1
	switch i := x.(type) {
	case nil:
		fmt.Printf("nil,x 的类型为%T\n", i)
	case int:
		fmt.Printf("int,x 的类型为%T\n", i)
	default:
		fmt.Println("unkown type")
	}
}
