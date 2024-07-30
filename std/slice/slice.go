package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{23, 454, 242, 45, 24, 234}
	//s2 := []int{23, 24, 45, 234}
	fmt.Println("IsSorted :", slices.IsSorted(s1))
	slices.Sort(s1)
	fmt.Println("Sorted s1 :", s1)
	n, found := slices.BinarySearch(s1, 45)
	if found {
		fmt.Println("found :", n, found)
	} else {
		fmt.Println("not found :", n, found)
	}

}
