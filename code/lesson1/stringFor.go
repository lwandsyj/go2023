package main

import "fmt"

func main() {
	s := "abcdefg"

	/* length := len(s)
	for i := 0; i < length; i++ {
		fmt.Printf("index=%d,s[i]=%c\n", i, s[i])
	} */

	for index, value := range s {
		fmt.Printf("index=%d,value=%c\n", index, value)
	}
}
