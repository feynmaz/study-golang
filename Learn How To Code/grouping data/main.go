package main

import "fmt"

func main() {
	// maps are fast for lookups by key

	m := map[string]int{
		"James": 32,
		"Miss": 27,
	}

	key := "Miss"
	if _, ok := m[key]; ok {
		delete(m, key)
		fmt.Printf("\"%v\" deleted", key)
	}


}