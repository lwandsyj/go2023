package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

// 使用复制(io.Copy 方法)的方式
func copy(url string) {
	// 获取远程数据，返回 *http.Response, 实现了reader 接口
	r, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()
	// 返回*os.File，file 实现了读和写接口
	file, _ := os.Create("copy.data")
	defer file.Close()
	writer := bufio.NewWriter(file)
	// 把reader 拷贝到writer 中
	io.Copy(writer, r.Body)
	writer.Flush()
}

func main() {

}
