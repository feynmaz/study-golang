package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

// func (r receiver) identifier(parameters) (return(s)) { ... }
