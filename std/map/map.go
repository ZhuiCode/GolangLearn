package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]string{
		"aaa": "aaa",
		"bbb": "bbb",
	}
	m2 := maps.Clone(m1)
	m2["bbb"] = "fffff"
	fmt.Println(m2)

}
