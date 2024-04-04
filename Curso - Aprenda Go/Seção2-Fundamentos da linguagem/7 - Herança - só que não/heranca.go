// O que se assemelha a herança em go:

package main

import "fmt"

type user struct {
	name string
	age  int
}

type student struct {
	user   // herança
	grade  int
	school string
}

// heranca.go
func main() {
	fmt.Println("Hello, World!")

	p1 := user{
		name: "João",s
		age:  20,
	}

	e1 := student{
		user:   p1,
		grade:  10,
		school: "USP",
	}

	fmt.Println(p1)
	fmt.Println(e1)
	fmt.Println(e1.name)

}
