package main

import (
	"reflect"
	"testing"
)

func TestSortDescending_EmptySliceForEmptyInput(t *testing.T) {
	input := []int{}
	expected := []int{}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnSliceForSingleElement(t *testing.T) {
	input := []int{1}
	expected := []int{1}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForTwoElements(t *testing.T) {
	input := []int{1, 2}
	expected := []int{2, 1}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForMixed(t *testing.T) {
	input := []int{1, 2, 8, 4, 12, 15, 5}
	expected := []int{15, 12, 8, 5, 4, 2, 1}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnSliceForTwoSameElements(t *testing.T) {
	input := []int{5, 5}
	expected := []int{5, 5}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnSliceForSameElements(t *testing.T) {
	input := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	expected := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForMixedWithTheSameElements(t *testing.T) {
	input := []int{5, 5, 2, 8, 2, 8, 4, 12, 18, 22, 3}
	expected := []int{22, 18, 12, 8, 8, 5, 5, 4, 3, 2, 2}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForMixedWithZero(t *testing.T) {
	input := []int{5, 5, 2, 8, 2, 8, 0, 4, 12, 18, 22, 3}
	expected := []int{22, 18, 12, 8, 8, 5, 5, 4, 3, 2, 2, 0}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnSliceForZero(t *testing.T) {
	input := []int{0}
	expected := []int{0}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForZeroAndOther(t *testing.T) {
	input := []int{0, 5}
	expected := []int{5, 0}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForNegatives(t *testing.T) {
	input := []int{-2, -1, -10, -15, -3}
	expected := []int{-1, -2, -3, -10, -15}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForNegativeAndPositive(t *testing.T) {
	input := []int{8, 10, -2, -1, -10, 1, -15, -3}
	expected := []int{10, 8, 1, -1, -2, -3, -10, -15}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForLargeElements(t *testing.T) {
	input := []int{555, 269, 123456, 1234567}
	expected := []int{1234567, 123456, 555, 269}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnDescendingSliceForMaxAndMinInt(t *testing.T) {
	input := []int{int(^uint(0) >> 1), 2, -int(^uint(0)>>1) - 1}
	expected := []int{int(^uint(0) >> 1), 2, -int(^uint(0)>>1) - 1}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnSliceForSingleNegative(t *testing.T) {
	input := []int{-1}
	expected := []int{-1}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestSortDescending_ReturnSliceForDescendingMixed(t *testing.T) {
	input := []int{30, 25, 20, 15, 10, 5, 4, 3, 2, 1, 0}
	expected := []int{30, 25, 20, 15, 10, 5, 4, 3, 2, 1, 0}
	sortDescending(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}
