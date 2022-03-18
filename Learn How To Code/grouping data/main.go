package main

import "fmt"

func main() {
	// x := type{values}  - composite literal
	x := []int{4,5,7,8,42}
	fmt.Println(x)

	x = append(x, 6, 9, 1, 0, 1002)
	fmt.Println(x)

	y := []int{12, 23, 34, 56, 67}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}