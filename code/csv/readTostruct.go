package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	id   int
	name string
	age  int
	sex  int // 1 男 0 nv
	num  string
}

func main() {
	// 文件路径
	filePath := "./files/user.csv"
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf(" Open file error:%v ", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 创建csv reader
	reader := csv.NewReader(file)
	// 读取全部内容
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("read csv file error:%v\n", err)
		return
	}
	// 第一行数据为字段行，只要结果行数据
	// 使用slice 切片
	dataRows := rows[1:]
	// struct 切片
	var users []User = make([]User, 0)
	//fmt.Println(users)
	for _, row := range dataRows {
		// csv读取回来的数据全部为string 类型，struct 中设计到int 类型
		id, _ := strconv.Atoi(row[0])
		age, _ := strconv.Atoi(row[2])
		sex, _ := strconv.Atoi(row[3])
		users = append(users, User{
			id:   id,
			name: row[1],
			age:  age,
			sex:  sex,
			num:  row[4],
		})
	}
	fmt.Printf("users=%+v\n", users)
}
