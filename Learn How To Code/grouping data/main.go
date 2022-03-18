package main

import "fmt"

func main() {
	// maps are fast for lookups by key

	m := map[string]int{
		"James": 32,
		"Miss M": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	// comma ok idiom
	if v, ok := m["random"]; ok {
		fmt.Println(v)
	}

	if v, ok := m["Miss M"]; ok {
		fmt.Println(v)
	}
}