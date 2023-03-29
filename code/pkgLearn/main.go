package main

import (
	"fmt"

	"pkgLearn/pkgtest"
)

func main() {
	fmt.Println("hellocl")
	p := pkgtest.People{Name: "zhangsan"}
	p.Name = "李四"

	fmt.Println(p)

	fmt.Println(pkgtest.GetName())
}
