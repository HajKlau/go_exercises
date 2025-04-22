package main

import (
	"reflect"
	"testing"
)

func TestCountChar_EmptyString(t *testing.T) {
	result := countChar("")
	if len(result) != 0 {
		t.Errorf("Expected empty map, got %v", result)
	}
}

func TestCountChar_AllSameChars(t *testing.T) {
	result := countChar("ooo")
	expected := map[rune]int{'o': 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountChar_AllUniqueChars(t *testing.T) {
	result := countChar("abcd")
	expected := map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountChar_MixedChars(t *testing.T) {
	result := countChar("accommodation")
	expected := map[rune]int{
		'a': 2,
		'c': 2,
		'o': 3,
		'm': 2,
		'd': 1,
		't': 1,
		'i': 1,
		'n': 1,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountChar_CaseSensitivity(t *testing.T) {
	result := countChar("AaBb")
	expected := map[rune]int{'A': 1, 'a': 1, 'B': 1, 'b': 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountChar_WhiteSpaceAndSpecialChars(t *testing.T) {
	result := countChar("a a!b@")
	expected := map[rune]int{
		'a': 2,
		' ': 1,
		'!': 1,
		'b': 1,
		'@': 1,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountChar_NumbersChars(t *testing.T) {
	result := countChar("12331")
	expected := map[rune]int{'1': 2, '2': 1, '3': 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountChar_Unicode(t *testing.T) {
	result := countChar("żółw🐰")
	expected := map[rune]int{'ż': 1, 'ó': 1, 'ł': 1, 'w': 1, '🐰': 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
