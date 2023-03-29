package main

import "fmt"

type People struct {
	name string
	age  int
}

func test() {
	/* var a int = 2
	var b bool = false
	var c rune = 'a'
	var s string = "hello"
	var d byte = '2'
	var e [3]int = [3]int{1, 2, 3}
	var f People = People{
		name: "张三",
		age:  12,
	}*/
	var g []int
	/* fmt.Printf("a=%p\n", a)
	fmt.Printf("b=%p\n", b)
	fmt.Printf("c=%p\n", c)
	fmt.Printf("s=%p\n", s)
	fmt.Printf("d=%p\n", d)
	fmt.Printf("e=%p\n", e)
	fmt.Printf("f=%p\n", f) */
	fmt.Printf("f=%p,g==nil:%t\n", g, g == nil)
	g = []int{1, 2, 3}
	fmt.Printf("f=%p,g==nil:%t\n", g, g == nil)
}

func test1(x int) {
	x = 20
}
func test2(x *int) {
	*x = 30
}
func test3(x string) {
	x = "hello world"
}
func test4(x *string) {
	*x = "string pointer"
}
func main() {
	x := 10
	// 传递int 副本，改变副本的值，并不会影响原有的变量的值
	test1(x)
	fmt.Printf("test1 不改变原有的值，x=%d\n", x)
	// 传递指针，指针是地址，改变地址中的值，会影响原有的变量的值
	test2(&x)
	fmt.Printf("test2 指针改变原有的值，x=%d\n", x)

	s := "hello"
	test3(s)
	fmt.Printf("test3不改变原有的值，s=%s\n", s)
	test4(&s)
	fmt.Printf("test4 指针改变原有的值，s=%s\n", s)

	f := People{
		name: "lisi",
		age:  18,
	}
	test5(f)
	fmt.Printf("test5不改变原有的值，s=%v\n", f)
	test6(&f)
	fmt.Printf("test6指针改变原有的值，s=%v\n", f)
}

func test5(p People) {
	p.name = "update 张三"
}
func test6(p *People) {
	*&p.name = "update pointer zhangsan"
}
