package main

import "fmt"

func main() {
	s := "abcdef"

	b := []byte(s)
	b[3] = ','
	c := []rune(s)
	c[3] = '|'

	fmt.Printf("s=%s\n", s)
	fmt.Printf("b=%s\n", b)
	fmt.Printf("c=%s\n", string(c)) // rune 字符类型

}
