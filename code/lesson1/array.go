package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	for index, value := range a {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}
}
