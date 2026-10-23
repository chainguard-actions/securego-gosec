package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("insecure"))
	fmt.Printf("%x\n", h.Sum(nil))
}
