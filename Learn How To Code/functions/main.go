package main

import "fmt"

func main() {
	foo()
	bar("James")

}

// func (r receiver) identifier(perameters) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println(s)
}
