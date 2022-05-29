package main

import "testing"

// test function
func TestIsPalindrome(t *testing.T) {
	actual := IsPalindrome("Abba")
	expected := true
	if actual != expected {
		t.Errorf("Expected %t is not same as"+
			" actual %t", expected, actual)
	}
}
