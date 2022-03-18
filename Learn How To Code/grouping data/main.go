package main

import "fmt"

func main() {
	// https://youtu.be/10LW7NROfOQ
	x := make([]int, 10, 100)
	fmt.Println(x)

	x[5] = 5
	fmt.Println(x)
}