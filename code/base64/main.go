package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var s string = "hello go"
	// 编码成base
	bs64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Printf("base64=%s\n", bs64)

	// 解码
	f, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解码以后值为：%s\n", string(f))
}
