package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
}

func (p Person) Walk() {
	fmt.Println(p.name + "andando")
}

func main() {
	//fmt.Println("Hello, world!")
	var pessoa Person
	pessoa.name = "lucas"
	pessoa.age = 25
	//pessoa.Walk()
	go cont(6)
	go cont(6)
	go cont(6)
}

func cont(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
