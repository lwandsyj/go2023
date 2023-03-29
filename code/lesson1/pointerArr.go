package main

import "fmt"

func main() {
	var p [3]*int
	a := 2
	b := a + 1
	c := b + 1
	p[0] = &a
	p[1] = &b
	p[2] = &c
	for index, value := range p {
		fmt.Printf("key=%d;value=%p,*value=%d\n", index, value, *value)
	}

}
