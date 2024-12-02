package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  []int
	Mapp map[string]int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Smith",
		Bio:  []int{1, 2, 3, 4},
		Mapp: map[string]int{"one": 1, "two": 2, "three": 3},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
