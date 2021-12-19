package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		"Gandhi",
		"Be the change",
	}

	jesus := sage{
		"Jesus",
		"Love all",
	}

	f := car{
		"Ford",
		"F150",
		2,
	}

	c := car{
		"Toyota",
		"Corolla",
		4,
	}

	sages := []sage{buddha, gandhi, jesus}
	cars := []car{f, c}

	data := items{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
