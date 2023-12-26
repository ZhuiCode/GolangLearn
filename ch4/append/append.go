package main

import "fmt"

func append(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	x := []int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		z := append(x, i)
		fmt.Printf("cap=%d\t%v\n", cap(z), z)
		x = z
	}
}
