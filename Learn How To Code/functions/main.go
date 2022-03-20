package main

import "fmt"

func main() {
	n := 4

	f := factorial(n)
	fmt.Println(f)

	f = 1
	for ; n > 0; n-- {
		f *= n
	} 
	fmt.Println(f)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}