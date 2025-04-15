package main

import (
	"fmt"
)

func countChar(s string) map[rune]int {

	counter := make(map[rune]int)

	for _, c := range s {
		counter[c]++
	}

	return counter
}

func printCounterChar(counter map[rune]int) {
	for k, v := range counter {
		fmt.Printf("'%c' : %d, ", k, v)
	}
	fmt.Println()

}

func main() {
	myString := "aabb"
	result := countChar(myString)
	fmt.Printf("Counted word characters: '%s'\n", myString)
	printCounterChar(result)
}
