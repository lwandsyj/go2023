package main

import "fmt"

func main() {
	var sMap = map[string]int{
		"tom": 18,
		"ben": 20,
	}
	// len 返回map 中的个数
	fmt.Printf("map key count=%d\n", len(sMap))

	// key 表示map 的key ,value 表示map 中key 对应的值
	for key, value := range sMap {
		fmt.Printf("key=%s,value=%d\n", key, value)
	}

	var sMap2 = make(map[string]int)
	// 使用key 作为下标读取和设置或更新value
	sMap2["tom"] = 3
}
