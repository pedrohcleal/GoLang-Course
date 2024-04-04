package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1
	fmt.Println("var1:", var1)
	fmt.Println("var2:", var2)

	var1++
	fmt.Println("var1:", var1)
	fmt.Println("var2:", var2)

	// var2 não é afetado por var1++
	var2++
	fmt.Println("var1:", var1)
	fmt.Println("var2:", var2)

	// Ponteiro é referencia de memoria
	var var3 int
	var ponteiro1 *int //guarda

	var3 = 100
	ponteiro1 = &var3
	fmt.Println("var3:", var3)
	fmt.Println("ponteiro1:", ponteiro1)        //
	fmt.Println("ponteiro1 valor:", *ponteiro1) //*ponteiro1 desreferencia

}
