package main

import "fmt"

func main() {
	
	x := 5

	func(x int){
		fmt.Printf("Anonymous func run, printing %v\n", x)
	}(x)

	
}