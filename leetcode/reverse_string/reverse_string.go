package main

import (
	"fmt"
)

func reverseString(s string) string {

	var CharInString []rune

	for _, c := range s {

		CharInString = append(CharInString, c)
	}

	var newReverseString string

	for i := len(CharInString) - 1; i >= 0; i-- {
		newReverseString += string(CharInString[i])
	}

	return newReverseString

}

func main() {
	result := reverseString("animal")
	fmt.Println("Reverse string:", result)
}
