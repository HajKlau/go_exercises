package main

import (
	"testing"
)

func TestFirstUniqueCharEmptyString(t *testing.T) {
	result := firstUniqueChar("")
	expected := -1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstUniqueCharAllSameChars(t *testing.T) {
	result := firstUniqueChar("aaa")
	expected := -1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstUniqueCharAllUniqueChars(t *testing.T) {
	result := firstUniqueChar("abcd")
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstUniqueCharMixed_Chars(t *testing.T) {
	result := firstUniqueChar("apple")
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstUniqueCharCaseSensitivity(t *testing.T) {
	result := firstUniqueChar("oOraNge")
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstUniqueCharWhiteSpaceAndSpecialChars(t *testing.T) {
	result := firstUniqueChar("ba a!b!na")
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstUniqueCharNumbersChars(t *testing.T) {
	result := firstUniqueChar("112213113456")
	expected := 9
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstUniqueCharUnicode(t *testing.T) {
	result := firstUniqueChar("peęęeopłpa")
	expected := 7
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
