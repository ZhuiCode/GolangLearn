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

	//byte.Buffer
	buf := bytes.NewBuffer(b)
	fmt.Printf("####[%d]\n", buf.Len())
	fmt.Printf("####[%s]\n", b)
	//flag := bytes.Contains(b, []byte("a"))
	s := []byte("hello world")
	fmt.Println(bytes.HasPrefix(s, []byte("el")))
	d := bytes.NewBuffer(s)
	fmt.Println(d.Len())

	//byte.Reader --只读操作
	r := bytes.NewReader(s) //将s的内容拷贝到r中
	f := make([]byte, 5)    //申请一块空间f
	r.Read(f)               //将r中的空间读入到f
	fmt.Println(string(f))
}
