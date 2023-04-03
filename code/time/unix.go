package main

import (
	"fmt"
	"time"
)

// 时间类型转成秒数
func ConvertTimeToNum() {
	// 当前时间
	now := time.Now()
	// 转成自1970年以来的秒数
	uxin := now.Unix()
	fmt.Printf("unix=%d\n", uxin)

	// 转成自1970年以来的毫秒数
	millli := now.UnixMilli()
	fmt.Printf("millli=%d\n", millli)
}

// 自1970年以来的秒数，毫秒数转成日期类型
func ConvertNumToTime() {
	// 自1970年以来的毫秒数
	num := 1680505186034
	// 自1970年以来的秒数
	miao := num / 1000
	fmt.Printf("miao=%d\n", miao)
	// 毫秒数转成日期
	t1 := time.UnixMilli(int64(num))
	fmt.Printf("t1=%v\n", t1.Format(time.DateTime))
	// 秒数
	t2 := time.Unix(int64(miao), 0)
	fmt.Printf("t2=%v\n", t2.Format(time.DateTime))
}
func main() {
	ConvertNumToTime()
}
