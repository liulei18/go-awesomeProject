package main

import "fmt"

type alive interface {
	walk()
}
type People struct {
	name string
	tele string
}

func (p People) walk() {
	fmt.Println(p.name)
}

type Man struct {
	People
}

func main() {

	m := Man{People{name: "people1"}}
	m.walk()
	m.People.walk()
	fmt.Println("Hello World")
}
