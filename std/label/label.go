package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
Label:
	for _, i := range list {
		if i == 2 {
			continue Label
		}
		fmt.Println(i)
	}
}
