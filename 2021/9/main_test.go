package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue int
	}{
		{
			input:         []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"},
			expectedValue: 15,
		},
	}

	for _, tc := range tt {
		var numbers [][]int
		for lineIndex, line := range tc.input {
			parts := strings.Split(line, "")
			numbers = append(numbers, make([]int, 0))

			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				numbers[lineIndex] = append(numbers[lineIndex], num)
			}
		}
		t.Run("TestPart1", func(t *testing.T) {
			actual := part1(numbers)

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
			input:         []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"},
			expectedValue: 1134,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart2", func(t *testing.T) {
			var numbers [][]int
			for lineIndex, line := range tc.input {
				parts := strings.Split(line, "")
				numbers = append(numbers, make([]int, 0))

				for _, part := range parts {
					num, _ := strconv.Atoi(part)
					numbers[lineIndex] = append(numbers[lineIndex], num)
				}
			}
			actual := part2(numbers)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}

func TestGetLowestPoints(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue [][]int
	}{
		{
			input:         []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"},
			expectedValue: [][]int{{1, 0, 1}, {9, 0, 0}, {2, 2, 5}, {6, 4, 5}},
		},
	}
	for _, tc := range tt {
		var numbers [][]int
		for lineIndex, line := range tc.input {
			parts := strings.Split(line, "")
			numbers = append(numbers, make([]int, 0))

			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				numbers[lineIndex] = append(numbers[lineIndex], num)
			}
		}

		t.Run("TestGetLowestPoints", func(t *testing.T) {
			if got := getLowestPoints(numbers); !reflect.DeepEqual(got, tc.expectedValue) {
				t.Errorf("getLowestPoints() = %v, want %v", got, tc.expectedValue)
			}
		})
	}
}
