package main

import "fmt"

func main() {
	fmt.Println(soma(2, 3))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Olá, mundo!")
	fmt.Println(result)

	resultadoSoma, resultadosubtracao := calculosmath(10, 5)
	fmt.Println(resultadoSoma, resultadosubtracao)
}

func soma(a int8, b int8) int8 {
	return a + b
}

// calculosmath retorna a soma e a subtração de dois números.
func calculosmath(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
