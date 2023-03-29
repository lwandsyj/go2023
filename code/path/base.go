package main

import (
	"fmt"
	"path"
)

// path.Base(path string) 返回路径的最后一个元素
// 如果路径为空，则 Base 返回“.”
// 如果路径完全由斜杠组成，则 Base 返回“/”。

func main() {
	// 返回最后/ 后面的内容，比如path.go
	fmt.Println(path.Base("../path/path.go"))

	// 在提取最后一个元素之前删除尾部斜线
	fmt.Println(path.Base("../path/index.html/"))

	// 如果路径完全由斜杠组成，则 Base 返回“/”。
	fmt.Println(path.Base(("/")))

	// 如果路径为空，则 Base 返回“.”
	fmt.Println(path.Base(""))
}
