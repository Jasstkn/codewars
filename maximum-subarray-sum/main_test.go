package main

import "testing"

func TestEmptyMaximumSubarraySum(t *testing.T) {
	got := MaximumSubarraySum([]int{})
	if got != 0 {
		t.Errorf("MaximumSubarraySum([]int{}) = %d; want 0", got)
	}
}

func TestMaximumSubarraySum(t *testing.T) {
	got := MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	expected := 6
	if got != expected {
		t.Errorf("MaximumSubarraySum([]int{}) = %d; want %d", got, expected)
	}
}

func TestAllNegativeMaximumSubarraySum(t *testing.T) {
	got := MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})
	expected := 0
	if got != expected {
		t.Errorf("MaximumSubarraySum([]int{}) = %d; want %d", got, expected)
	}
}
