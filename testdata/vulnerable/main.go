package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename) // G304
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println(f)
}
