package main

import "fmt"

func main() {
	//var colors map[string]string

	colors := make(map[int]string)
	//colors := map[string]string{
	//	"red": "#FF0000",
	//	"green": "#00FF00",
	//	"blue": "#0000FF",
	//}

	//colors["red"] = "#FF0000"
	colors[10] = "#GGG111"

	delete(colors, 10)

	fmt.Println(colors)
}

