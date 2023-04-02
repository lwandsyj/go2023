package main

import (
	"fmt"
	"os"
)

// 获取进程的id
func main() {
	// 获取当前进程的id
	fmt.Println(os.Getpid())
	// 当前进程的父进程的id
	fmt.Println(os.Getppid())
}
