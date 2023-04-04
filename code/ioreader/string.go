package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 要读取的字符串
	content := `
	hello,go!
	你好，中国go!
	`
	// 创建string Reader
	reader := strings.NewReader(content)
	// 每次读取3个字节
	var b []byte = make([]byte, 4)
	for {
		// 每次读取byte slice 设置的长度
		n, err := reader.Read(b)
		if err != nil {
			// err 为io.EOF 表示读取完毕
			// 不等于io.EOF 表示是其他的错误
			if err != io.EOF {
				fmt.Printf("err:=%v\n", err)
				return // 退出函数
			} /* else { // 当err==io.EOF 时表示读取完毕
				break
			} */
		}
		// n 返回读取到的字节数
		fmt.Printf("n=%d\n", n)
		// 每次读取的数据存放在b 中
		fmt.Printf("b=%s\n", string(b))
		// n==0 表示执行到最后
		if n == 0 {
			break
		}
	}
}
