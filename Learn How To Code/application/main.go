package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	s := `[{"First":"John","Last":"Smith","Age":32},{"First":"Alice","Last":"A","Age":27}]`
	bs := []byte(s)

	people := []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(people)
}
