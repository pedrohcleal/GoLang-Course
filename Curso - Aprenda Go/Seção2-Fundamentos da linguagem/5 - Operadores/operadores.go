package main

import "fmt"

func main() {
	// op aritmeticos:
	a := 10
	b := 20
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("a / b=", a/b)
	fmt.Println("a % b=", a%b)

	// op relacionais:
	fmt.Println("a == b=", a == b)
	fmt.Println("a != b=", a != b)
	fmt.Println("a > b=", a > b)
	fmt.Println("a < b=", a < b)
	fmt.Println("a >= b=", a >= b)
	fmt.Println("a <= b=", a <= b)

	// op lógicos:
	verdade1 := true
	falso := false
	fmt.Println("verdade1 && falso=", verdade1 && falso)
	fmt.Println("verdade1 || falso=", verdade1 || falso)
	fmt.Println("!(verdade1)=", !verdade1)

	// op de atribuição:
	a = 10
	b = 20
	a += b
	fmt.Println("a += b=", a)

	// op unários:
	a = 10
	b = 20
	fmt.Println(-a) //resultado= 9
	fmt.Println(+a) //resultado= 10

	// op de incremento e decremento:
	a = 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

}
