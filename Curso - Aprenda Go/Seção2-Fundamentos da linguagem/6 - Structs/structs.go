package main

import "fmt"

type user struct {
	name   string
	email  string
	age    int
	adress adress // struct aninhada
}

type adress struct {
	street string
	city   string
}

// structs.go

// main.go

func main() {
	fmt.Println("arquivo structs.go")

	var u user
	fmt.Println(u) // { 0}
	// atribuindo valores aos campos da struct)
	u.name = "John Doe"
	u.email = "john.doe@example.com"
	u.age = 30
	u.adress = adress{
		street: "123 Main Street",
		city:   "Anytown",
	}

	fmt.Println(u)

	user2 := user{
		name:  "Jane Doe",
		email: "jane.doe@example.com",
		age:   25,
		adress: adress{
			street: "456 Elm Street",
			city:   "Anytown",
		},
	}

	user2.age = 26

	fmt.Println(user2)

}
