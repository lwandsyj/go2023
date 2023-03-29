package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello world"
	b := "aaaabbbbcccc"
	// strings.Compare 比较两个字符串的大小
	fmt.Printf("strings.Compare(a,b)==%d\n", strings.Compare(a, b))
	// strings.Contains(str,substr string) 判断s 是否包含子字符串substr
	fmt.Printf("strings.Contains(a,\"hello\")==%t\n", strings.Contains(a, "hello"))
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	// 检测是否包含某个字符，比如 a 是否在 hello 中
	fmt.Printf("a is in hello ==%t\n", strings.ContainsRune("hello", 'a'))
	fmt.Printf("a is in hello ==%t\n", strings.ContainsRune("aave", 97))
	// Count 返回子字符串 在字符串中出现的次数
	fmt.Printf("counts=%d\n", strings.Count("hello", "l"))

	// HasPrefix 是否已prefix 开始
	fmt.Printf("startWith =%t\n", strings.HasPrefix("hello", "he"))
	fmt.Printf("startWith =%t\n", strings.HasPrefix("hello", "e"))

	// HasSuffix 是否已suffix 结尾
	fmt.Printf("endsWith = %t\n", strings.HasSuffix("hello", "lo"))
	fmt.Printf("endsWith= %t\n", strings.HasSuffix("hello", "he"))

}
