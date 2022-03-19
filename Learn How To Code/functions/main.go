package main

import "fmt"

func main() {
	ii := []int{2, 3, 4, 5, 6, 7}
	foo(ii...)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }
