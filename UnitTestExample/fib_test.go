package main

import "testing"


func TestFib(t *testing.T) {
	// Test the fib function
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range testCases {
		if output := fib(test.n); output != test.expected {
			t.Errorf("Test failed: %v inputted, %v expected, %v received", test.n, test.expected, output)
		}
	}
}