package main

import "fmt"

// struct is data structure which allows us to compose togeether values of different types
type person struct {
	first string
	last string
}

func main() {
	// value of type person
	p1 := person{
		first: "James",
		last: "Bond",
	}

	fmt.Println(p1.last)
	fmt.Println(p1)

}