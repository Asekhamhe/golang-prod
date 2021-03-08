package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum224([]byte("x"))
	fmt.Printf("%x", c1)
}
