package main

import "fmt"

func main() {

	f := func(x int){
		fmt.Printf("Anonymous func run, printing %v\n", x)
	}

	f(4)


}