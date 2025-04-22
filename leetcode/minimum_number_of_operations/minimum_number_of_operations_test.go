package main

import (
	"testing"
)

func TestMinimumNumberOfOperationsEmptySlice(t *testing.T) {
	result := minimumNumberOfOperations([]int{})
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMinimumNumberOfOperationsAllUniqueElements(t *testing.T) {
	result := minimumNumberOfOperations([]int{1, 2, 3})
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMinimumNumberOfOperationsAllSameElements(t *testing.T) {
	result := minimumNumberOfOperations([]int{3, 3, 3, 3, 3})
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMinimumNumberOfOperationsMixedElements(t *testing.T) {
	result := minimumNumberOfOperations([]int{2, 2, 4, 7, 1, 4, 8, 0, 0, 5, 6, 8, 9, 3, 8, 8, 8, 1, 2, 7})
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMinimumNumberOfOperationsShorterThanThree(t *testing.T) {
	result := minimumNumberOfOperations([]int{5, 5})
	expected := 1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMinimumNumberOfOperationsThreeRepeatsElements(t *testing.T) {
	result := minimumNumberOfOperations([]int{5, 5, 5})
	expected := 1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMinimumNumberOfOperationsAlternatingRepeats(t *testing.T) {
	result := minimumNumberOfOperations([]int{2, 4, 2, 4, 2, 4, 2, 4, 2, 4})
	expected := 3
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
