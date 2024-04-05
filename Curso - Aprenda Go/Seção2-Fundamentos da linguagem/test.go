package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
	fmt.Println(ToCamelCase("The_Stealth-Warrior"))

}

func ToCamelCase(s string) string {

	for _, letra := range s {
		fmt.Println((letra))

	}

	words := strings.Split(s, "_")
	words2 := strings.Split(s, "-")

	fmt.Println(words, words2)

	return ""
}
