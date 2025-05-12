package main

import (
	"fmt"
	"sort"
)

func sortDescending(nums []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
}

func main() {
	mySlice := []int{5, 1, 6, 2, 4}
	sortDescending(mySlice)

	for _, num := range mySlice {
		fmt.Printf("%d ", num)
	}

	fmt.Println()
}
