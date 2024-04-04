package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int8 = 10
	var number2 uint8 = 10 //tipo de dado
	var number3 int16 = -10
	//var number4 uint16 = -30000

	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)

	float32, float64 := 10.0, 10.0
	fmt.Println(float32)
	fmt.Println(float64)

	var text string = "Hello, world!"
	fmt.Println(text)

	var boolean bool = true
	fmt.Println(boolean)

	charb := "b"
	chara := "A"
	fmt.Println(chara)
	fmt.Println(charb)

	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array)

	var slice []int = []int{1, 2, 3}
	fmt.Println(slice)

	var mapa map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(mapa)

	var error error = fmt.Errorf("this is an error")
	fmt.Println(error)

	var erro2 error = errors.New("this is an error")
	fmt.Println(erro2)

}
