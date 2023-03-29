package main

import "fmt"

func Switch(marks int, grade string) {
	switch marks {
	case 90:
		grade = "A"
		// 不用显示使用break, 在go语言中默认中断
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的分为marks=%d,成绩grade=%s\n", marks, grade)
}

func main() {
	grade := "B"
	marks := 90

	switch {
	case marks >= 90:
		grade = "A"
	case marks >= 80:
		grade = "B"
	case marks >= 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的分为marks=%d,成绩grade=%s\n", marks, grade)

}
