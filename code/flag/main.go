package main

import (
	"flag"
	"fmt"
)

func main() {
	var args []string = flag.Args()
	// 收集参数
	flag.Parse()

	// flag.Args() 收集没有命名的参数
	// flag.Args() 只有在flag.Parse() 之后才能获取
	var argsAfter []string = flag.Args()

	fmt.Printf("Parse before :%v\n", args)
	fmt.Printf("Parse after :%v\n", argsAfter)
}
