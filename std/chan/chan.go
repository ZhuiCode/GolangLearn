package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int16
}

func recvChan(ch <-chan int) {

	for val := range ch {
		fmt.Println("recvChan", val)
	}
}
func sendChan(ch chan<- int) {
	for val := 0; val < 3; val++ {
		fmt.Println("sendChan", val)
		ch <- val
	}
}
func printPerson(ch <-chan Person) {

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
func main() {
	ch := make(chan int, 0)

	go sendChan(ch)
	go recvChan(ch)

	//ch <- 4
	//time.Sleep(time.Second * 10)
	/*
		st := make(chan Person, 10)
		st <- Person{"aaaa", 23}
		st <- Person{"vvvvv", 33}
		go printPerson(st)
	*/
	time.Sleep(time.Second * 10)

}
