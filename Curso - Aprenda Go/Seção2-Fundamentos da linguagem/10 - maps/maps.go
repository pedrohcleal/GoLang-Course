package main

import "fmt"

func main() {
	fmt.Println("map")

	user := map[string]string{
		"name":       "Pedro ",
		"sobrenoome": " Leal"}

	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["sobrenome"])

}
