package main

import (
	"fmt"
)

func minimumNumberOfOperations(numbers []int) int {

	operationsNumber := 0

	for {
		setForSizeComparison := make(map[int]struct{})
		for _, num := range numbers {
			setForSizeComparison[num] = struct{}{}
		}

		if len(setForSizeComparison) == len(numbers) || len(numbers) == 0 {
			break
		}

		elementsToRemove := 3
		if len(numbers) < 3 {
			elementsToRemove = len(numbers)
		}
		numbers = numbers[elementsToRemove:]
		operationsNumber++
	}

	return operationsNumber

}

func main() {
	myVector := []int{2, 2, 4, 7, 1, 4, 8, 0, 0, 5, 6, 8, 9, 3, 8, 8, 8, 1, 2, 7}
	result := minimumNumberOfOperations(myVector)
	fmt.Printf("Minimum numbers of operations to make elements in array distinct: %d\n", result)
}
