package main

import (
	"fmt"
	"path"
)

func main() {
	filePath := "../vest/doc.md"

	fmt.Println(path.Dir(filePath))
}
