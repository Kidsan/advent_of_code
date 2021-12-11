package main

import (
	"testing"
)

func TestTick(t *testing.T) {
	tt := []struct {
		input         [][]int
		expectedValue int
	}{
		{
			input: [][]int{
				{1, 1, 1, 1, 1},
				{1, 9, 9, 9, 1},
				{1, 9, 1, 9, 1},
				{1, 9, 9, 9, 1},
				{1, 1, 1, 1, 1},
			},
			expectedValue: 9,
		},
	}

	for _, tc := range tt {
		t.Run("TestTick", func(t *testing.T) {
			actual := tick(tc.input)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}
