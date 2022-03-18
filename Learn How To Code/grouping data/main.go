package main

import "fmt"

func main() {
	// x := type{values}  - composite literal
	x := []int{4,5,7,8,42}
	fmt.Println(x)

	fmt.Println(x[1:3])
}