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

	// 字符串连接 +

	e := "hello"

	fmt.Printf("e=%s\n", e+" world")

	// 切片

	f := "hello china"
	// 切片 f[i,j]
	// i 不存在，表示从下标0开始，到j 结束
	g := f[:3] // 返回hel
	fmt.Printf("g=%s\n", g)
	// j 不存在, 表示从i 开始，到字符串结束
	h := f[2:] // llo china
	fmt.Printf("h=%s\n", h)
	// i,j 都存在，从i 到j, 但是不包括j
	i := f[2:7] // llo c
	fmt.Printf("i=%s\n", i)

}
