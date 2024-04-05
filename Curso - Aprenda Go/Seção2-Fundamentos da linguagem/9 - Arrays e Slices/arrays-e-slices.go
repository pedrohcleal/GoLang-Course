package main

import "fmt"

// arrays e slices - introdução
func main() {
	fmt.Println("Arrays e slices - introdução")

	var array1 [5]int
	fmt.Println("array1:", array1)
	array1[0] = 1
	fmt.Println("array1:", array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2:", array2)

	array3 := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("array3:", array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("array4:", array4)

	// slices - introdução
	fmt.Println("\nSlices - introdução")

	slice0 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice0:", slice0)

	slice1 := array2[1:3]
	fmt.Println("slice1:", slice1)

	//arrays internos

	slice3 := make([]int, 5, 10)
	fmt.Println("slice3:", slice3)
	fmt.Println("len(slice3):", len(slice3))
	fmt.Println("cap(slice3):", cap(slice3))

	slice4 := make([]int, 5)
	fmt.Println("slice4:", slice4)
	fmt.Println("len(slice4):", len(slice4))
	fmt.Println("cap(slice4):", cap(slice4))

}
