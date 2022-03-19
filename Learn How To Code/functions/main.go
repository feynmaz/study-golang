package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	s2 := bar()
	fmt.Printf("%T\n", s2)
	fmt.Println(s2())
}

func foo() string {
	s := "Hello, world!"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}