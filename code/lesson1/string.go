package main

import "fmt"

func main() {
	// 字符rune 类型使用单引号括起来
	a := 'c'

	fmt.Printf("a=%c \n", a)

	// 字符串使用string 类型，使用双引号括起来
	var b string
	if b == "" {
		fmt.Printf("zero b=%s \n", b)
		fmt.Printf("zero b 是否等于空%t\n", b == "")
	}
	b = "hello"
	fmt.Printf("new b=%s \n", b)

	// 字符串下标获取字符
	d := b[0]

	fmt.Printf("d=%c \n", d)
}
