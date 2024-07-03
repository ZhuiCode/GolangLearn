package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("abc")
	fmt.Printf("####[%s]\n", b)
	c := bytes.Clone(b)
	fmt.Printf("####[%s]\n", c)
	c[0] = 'd'
	fmt.Printf("####[%s]\n", b)
	fmt.Printf("####[%s]\n", c)
	buf := bytes.NewBuffer(b)
	fmt.Printf("####[%d]\n", buf.Len())
	fmt.Printf("####[%s]\n", b)
}
