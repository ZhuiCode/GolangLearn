package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
func main() {
	const GoUsage = `aaa
	add
	cc`
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a.go"))
	fmt.Println(basename("abc"))
	fmt.Println(GoUsage)
}
