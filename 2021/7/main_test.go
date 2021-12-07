package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		input         []int
		expectedValue int
	}{
		{
			input:         []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, // []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expectedValue: 37,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart1", func(t *testing.T) {
			actual := part1(tc.input, 16)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		input         []int
		expectedValue int
	}{
		{
			input:         []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, // []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expectedValue: 168,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart2", func(t *testing.T) {
			actual := part2(tc.input, 16)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}
