package main

import "fmt"

type person struct {
	first string
	last  string
	age   uint8
}

type agent struct {
	person
	ltk bool
}

type human interface {
	// any type that has speak method is also of type human
	speak()
}

func bar(h human) {
	switch h.(type) {
		case agent:
			fmt.Println("person", h.(agent).first)
	}
}

func (a agent) speak() {
	fmt.Println(a.last)
	fmt.Println(a.first, a.last)
}

func main() {
	a := agent{
		person: person{
			"James",
			"Bond",
			35,
		},
		ltk: true,
	}

	fmt.Println(a)
	a.speak()

	bar(a)
}