package main

import "fmt"

func main() {

	var x int
	fmt.Printf("%T %v\n", x, x)
	x++

	{
		y := 42
		fmt.Println(y)
	}
}
