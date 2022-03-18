package main

import "fmt"

func main() {
	// maps are fast for lookups by key

	m := map[string]int{
		"James": 32,
		"Miss": 27,
	}

	m["Alex"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

}