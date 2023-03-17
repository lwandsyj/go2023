package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	var x, y int
	fmt.Println(x)
	fmt.Println(y)
	x = 2
	y = 3
	z := Add(x, y)
	fmt.Println(z)
}
