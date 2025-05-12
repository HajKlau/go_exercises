package main

import (
	"testing"
)

func TestReverseString_EmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_SingleChar(t *testing.T) {
	input := "a"
	expected := "a"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_TwoDifferentChars(t *testing.T) {
	input := "ab"
	expected := "ba"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_TwoSameChars(t *testing.T) {
	input := "aa"
	expected := "aa"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_NormalString(t *testing.T) {
	input := "animal"
	expected := "lamina"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_LongString(t *testing.T) {
	input := "animalisorangeandblackandhavedifferentname"
	expected := "emantnereffidevahdnakcalbdnaegnarosilamina"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_WithWhiteSpace(t *testing.T) {
	input := "ora nge"
	expected := "egn aro"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_WhiteSpaceAndSpecialChars(t *testing.T) {
	input := "or#a ng^e"
	expected := "e^gn a#ro"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_SpecialChars(t *testing.T) {
	input := "or#ang^e"
	expected := "e^gna#ro"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_DigitsAndLetters(t *testing.T) {
	input := "1a2b"
	expected := "b2a1"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_DigitsOnly(t *testing.T) {
	input := "1234"
	expected := "4321"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
