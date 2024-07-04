package main

import (
	"embed"
	"fmt"
)

// var a []byte
//
//go:embed hello.txt hello2.txt
var f embed.FS

func main() {
	//fmt.Println(s)
	//fmt.Println(string(s))
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))

	data, _ = f.ReadFile("hello2.txt")
	fmt.Println(string(data))
}
