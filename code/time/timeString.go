package main

import (
	"fmt"
	"time"
)

// 时间和string 互转
func main() {
	// 当前时间
	t := time.Now()
	// 时间类型转出string
	c := t.Format(time.DateTime)
	fmt.Println(c)
	// 只保留日期
	d := t.Format(time.DateOnly)
	fmt.Println(d)
	// 只保留时间
	h := t.Format(time.TimeOnly)
	fmt.Println(h)

	s := "2023-03-15 22:34:56"
	// 字符串转时间
	e, err := time.Parse(time.DateTime, s)
	// err 为nil ，表示转换成功
	if err == nil {
		fmt.Println(e)
	}
}
