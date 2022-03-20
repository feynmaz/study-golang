package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5}
	s := sum(ii...)
	fmt.Println(s)

	e := even(sum, ii...) 
	fmt.Println(e)

}

func sum(si ...int) int {
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}

func even (f func(ei ...int) int, vi ...int) int {
	yi := make([]int, 0, len(vi))

	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}