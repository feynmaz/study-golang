package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct{
	firstname string
	lastname string
	contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("Jonah")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newName string) {
	p.firstname = newName
}
