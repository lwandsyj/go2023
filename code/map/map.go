package main

import "fmt"

func main() {
	a := map[string]string{
		"name": "张三",
		"age":  "18",
	}
	// 如果存在，则返回对应值
	name := a["name"]
	fmt.Println(name) // 张三
	// 返回值和ok, ok 为true ，表示存在，false 表示不存在
	sex, ok := a["sex"] // sex 不存在，ok 返回false
	if ok {
		fmt.Printf("sex:%s\n", sex)
	} else {
		fmt.Println("not exist") // 执行此语句
	}
	age, ageOk := a["age"] // age 存在，ageOk 返回true
	if ageOk {
		fmt.Printf("age exist,age=%s\n", age) // 执行此语句
	} else {
		fmt.Println("age not exist")
	}
	// 不存在，返回值类型的零值，比如此实例中值为 string类型，零值为“”
	if x := a["sex"]; x == "" { // sex 不存在，返回默认值类型的零值
		fmt.Println("sex not exist")
	} else {
		fmt.Printf("x=%s\n", x)
	}
}
