package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	testCases := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 9},
		[]int{1, 1, 2},
		[]int{5, 1, 6},
	}

	for _, testCase := range testCases {
		result := sum(testCase[0], testCase[1])
		if result != testCase[2] {
			t.Error("Incorrect sum data", testCase[0], "+", testCase[1], "must be", testCase[2])
		}
	}
}
