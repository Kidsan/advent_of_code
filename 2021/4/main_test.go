package main

import (
	"testing"
)

func TestMarkNumber(t *testing.T) {

	t.Run("markNumber", func(t *testing.T) {
		var board = *NewBoard("00,11,22,33,44\n55,66,77,88,99\n4,5,6,7,8\n84,23,46,16,8\n6,700,95,75,06")
		board.markNumber(4)
		if board[0][2].marked != false {
			t.Errorf("Didn't mark")
		}
	})

}

func TestPart1(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue int
	}{

		{
			input:         []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1", "22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n", " 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n", "14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7\n"},
			expectedValue: 4512,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart1", func(t *testing.T) {
			actual := temp(tc.input)
			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue int
	}{

		{
			input:         []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1", "22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n", " 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n", "14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7\n"},
			expectedValue: 1924,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart2", func(t *testing.T) {
			actual := temp(tc.input)
			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}
		})
	}
}
