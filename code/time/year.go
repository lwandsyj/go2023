package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间，返回Time 类型
	t := time.Now()
	// 获取当前年份
	fmt.Println(t.Year())
	// 获取当前月份
	fmt.Println(t.Month()) // 默认调用的是Month 的string() 方法
	fmt.Printf("month convert to int= %d\n", int(t.Month()))
	// 获取日期
	fmt.Printf("day is %d\n", t.Day())
	// 获取小时
	fmt.Printf("hour is %d\n", t.Hour())
	// 获取分钟
	fmt.Printf("minute is %d\n", t.Minute())
	// 获取秒数
	fmt.Printf("second is %d\n", t.Second())

}
