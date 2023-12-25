package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func intsToStrig(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	values := []int{2, 3, 4, 5}
	fmt.Println(intsToStrig(values))
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println("aaa", y, "bbb", strconv.Itoa(x))
	fmt.Print(strconv.FormatInt(int64(x), 2))
}
