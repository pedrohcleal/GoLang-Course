package main

import "fmt"

func main() {
	var nome string = "João"
	var idade int = 30
	variavel := 1.5 // tipo inferido
	fmt.Println(nome, idade)
	fmt.Println(variavel)
	var (
		variavel3 string
		variavel4 int
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Maria", 25
	fmt.Println(variavel5, variavel6)

	const const1 string = "Constante1"

	fmt.Println(const1)
	//não é possível alterar o valor de uma constante
}
