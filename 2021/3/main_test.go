package main

import "testing"

func TestPart1(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue int64
	}{

		{
			input:         []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			expectedValue: 198,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart1", func(t *testing.T) {
			actual := Part1(tc.input)
			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}
		})
	}
}

func TestGetOxygenGeneratorRating(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue string
	}{

		{
			input:         []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			expectedValue: "10111",
		},
	}

	for _, tc := range tt {
		t.Run("getOxygenGeneratorRating", func(t *testing.T) {
			actual := getOxygenGeneratorRating(tc.input, 0)
			if actual != tc.expectedValue {
				t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
			}
		})
	}
}

func TestGetCO2ScrubberRating(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue string
	}{

		{
			input:         []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			expectedValue: "01010",
		},
	}

	for _, tc := range tt {
		t.Run("getCO2ScrubberRating", func(t *testing.T) {
			actual := getCO2ScrubberRating(tc.input, 0)
			if actual != tc.expectedValue {
				t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue int64
	}{

		{
			input:         []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			expectedValue: 230,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart2", func(t *testing.T) {
			actual := Part2(tc.input)
			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}
		})
	}
}
