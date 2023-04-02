package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Environ() 返回所有的环境变量
	var envs []string = os.Environ()

	// 遍历环境变量
	for _, value := range envs {
		fmt.Println(value)
	}

}
