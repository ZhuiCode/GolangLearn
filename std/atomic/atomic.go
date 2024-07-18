package main

import (
	"fmt"
	"sync/atomic"
)

type Person struct {
	name string
	age  int8
}

func main() {
	var a int32
	atomic.StoreInt32(&a, 1)
	fmt.Println(a)

	fmt.Println(atomic.LoadInt32(&a))
	oldValue := atomic.SwapInt32(&a, 333)
	fmt.Println(oldValue, a)

	atomic.AddInt32(&a, 222)
	fmt.Println(a)

	var b atomic.Int32
	b.Add(12)
	fmt.Println("BBBB", b.Load())

	var c atomic.Value
	p := Person{
		name: "aa",
		age:  18,
	}
	c.Store(p)
	fmt.Println(c.Load())
}
