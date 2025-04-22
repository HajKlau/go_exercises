package main

import (
	"fmt"
)

func firstUniqueChar(s string) int {

	charCounter := make(map[rune]int)

	for _, c := range s {
		charCounter[c]++
	}

	for i, c := range s {
		if charCounter[c] == 1 {
			return i
		}
	}

	return -1
}

func main() {

	newString := "accomodation"
	result := firstUniqueChar(newString)

	if result != -1 {
		fmt.Printf("First unique character in string '%s' has index: [%d]\n", newString, result)
	} else {
		fmt.Printf("No unique character found in string '%s'\n", newString)
	}
}
