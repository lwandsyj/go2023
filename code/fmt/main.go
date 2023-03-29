package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {
	p := People{"张三", 20}

	fmt.Printf("p=%v\n", p)

	fmt.Printf("+p=%+v\n", p)

	fmt.Printf("#p=%#v\n", p)
}
