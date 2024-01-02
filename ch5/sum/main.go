package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum(2, 3, 4, 5))
	values := []int{2, 3, 4, 5}
	fmt.Println(sum(values[2:4]...))
}
