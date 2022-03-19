package main

import "fmt"

func main() {
	foo()
	foo(1, 2, 3)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }
