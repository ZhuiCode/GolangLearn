package main

import (
	"fmt"
)

type Sayer interface {
	say()
}
type Mover interface {
	move()
	gooo()
}

type dog struct {
	name string
}

type cat struct {
	name string
}

func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}
func (d dog) gooo() {
	fmt.Printf("%s go go go\n", d.name)
}

func (d cat) say() {
	fmt.Printf("%s会叫喵喵喵\n", d.name)
}

func (d cat) move() {
	fmt.Printf("%s会动\n", d.name)
}
func (d cat) gooo() {
	fmt.Printf("%s go go go\n", d.name)
}
func main() {
	var x Sayer
	var a = dog{name: "旺财"}
	var b = cat{name: "Kitty"}

	x = a
	x.say()

	x = b
	x.say()

	/*
	   var s int
	   var z interface{}
	   z = s

	   _, ok := z.(float32)

	   fmt.Println("aaaa", ok)
	*/
}
