package main

import (
	"fmt"
)

const (
	a = 1
	b
	c = 2
	d
)

const (
	e = iota
	f
	g
	h = 4
	m
	i = iota
	j
)

func main() {
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
	fmt.Printf("c=%d\n", c)
	fmt.Printf("d=%d\n", d)
	fmt.Printf("e=%d\n", e)
	fmt.Printf("f=%d\n", f)
	fmt.Printf("g=%d\n", g)
	fmt.Printf("h=%d\n", h)
	fmt.Printf("m=%d\n", m)
	fmt.Printf("i=%d\n", i)
	fmt.Printf("j=%d\n", j)
}
