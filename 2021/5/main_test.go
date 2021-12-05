package main

import (
	"fmt"
	"testing"
)

func TestParsePoints(t *testing.T) {
	tt := []struct {
		input         string
		expectedValue []point
	}{
		{
			input:         "0,9 -> 5,9",
			expectedValue: []point{{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9}, {x: 3, y: 9}, {x: 4, y: 9}, {x: 5, y: 9}},
		},
		{
			input:         "9,7 -> 7,9",
			expectedValue: []point{{x: 7, y: 9}, {x: 8, y: 8}, {x: 9, y: 7}},
		},
		{
			input:         "1,1 -> 3,3",
			expectedValue: []point{{x: 1, y: 1}, {x: 2, y: 2}, {x: 3, y: 3}},
		},
		{
			input:         "1,1 -> 3,3",
			expectedValue: []point{{x: 1, y: 1}, {x: 2, y: 2}, {x: 3, y: 3}},
		},

		{
			input: "0,0 -> 8,8",
			expectedValue: []point{
				{x: 0, y: 0}, {x: 1, y: 1}, {x: 2, y: 2}, {x: 3, y: 3}, {x: 4, y: 4}, {x: 5, y: 5}, {x: 6, y: 6}, {x: 7, y: 7}, {x: 8, y: 8},
			},
		},
		{
			input: "5,5 -> 8,2",
			expectedValue: []point{
				{x: 5, y: 5}, {x: 6, y: 4}, {x: 7, y: 3}, {x: 8, y: 2},
			},
		},
		{
			input: "6,4 -> 2,0",
			expectedValue: []point{
				{x: 2, y: 0}, {x: 3, y: 1}, {x: 4, y: 2}, {x: 5, y: 3}, {x: 6, y: 4},
			},
		},
		{
			input: "8,0 -> 0,8",
			expectedValue: []point{
				{x: 0, y: 8}, {x: 1, y: 7}, {x: 2, y: 6}, {x: 3, y: 5}, {x: 4, y: 4}, {x: 5, y: 3}, {x: 6, y: 2}, {x: 7, y: 1}, {x: 8, y: 0},
			},
		},
	}

	for _, tc := range tt {
		t.Run("TestParsePoints", func(t *testing.T) {
			actual := parsePoints(tc.input, true)
			fmt.Println(actual)
			for index, value := range actual {
				point := value
				if point.x != tc.expectedValue[index].x && point.y != tc.expectedValue[index].y {
					t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
				}
			}

		})
	}
}

func TestPart1(t *testing.T) {
	tt := []struct {
		input         []Line
		expectedValue int
	}{

		{
			input: []Line{
				{
					points: []point{
						{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9}, {x: 3, y: 9}, {x: 4, y: 9}, {x: 5, y: 9},
					},
				},
				{
					points: []point{
						{x: 9, y: 4}, {x: 8, y: 4}, {x: 7, y: 4}, {x: 6, y: 4}, {x: 5, y: 4}, {x: 4, y: 4}, {x: 3, y: 4},
					},
				},
				{
					points: []point{
						{x: 2, y: 2}, {x: 2, y: 1},
					},
				},
				{
					points: []point{
						{x: 7, y: 0}, {x: 7, y: 1}, {x: 7, y: 2}, {x: 7, y: 3}, {x: 7, y: 4},
					},
				},
				{
					points: []point{
						{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9},
					},
				},
				{
					points: []point{
						{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9},
					},
				},
				{
					points: []point{
						{x: 3, y: 4}, {x: 2, y: 4}, {x: 1, y: 4},
					},
				},
			},
			expectedValue: 5,
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

func TestPart2(t *testing.T) {
	tt := []struct {
		input         []Line
		expectedValue int
	}{

		{
			input: []Line{
				{
					points: []point{
						{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9}, {x: 3, y: 9}, {x: 4, y: 9}, {x: 5, y: 9},
					},
				},
				{
					points: []point{
						{x: 9, y: 4}, {x: 8, y: 4}, {x: 7, y: 4}, {x: 6, y: 4}, {x: 5, y: 4}, {x: 4, y: 4}, {x: 3, y: 4},
					},
				},
				{
					points: []point{
						{x: 2, y: 2}, {x: 2, y: 1},
					},
				},
				{
					points: []point{
						{x: 7, y: 0}, {x: 7, y: 1}, {x: 7, y: 2}, {x: 7, y: 3}, {x: 7, y: 4},
					},
				},
				{
					points: []point{
						{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9},
					},
				},
				{
					points: []point{
						{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9},
					},
				},
				{
					points: []point{
						{x: 3, y: 4}, {x: 2, y: 4}, {x: 1, y: 4},
					},
				},

				{
					points: []point{
						{x: 0, y: 0}, {x: 1, y: 1}, {x: 2, y: 2}, {x: 3, y: 3}, {x: 4, y: 4}, {x: 5, y: 5}, {x: 6, y: 6}, {x: 7, y: 7}, {x: 8, y: 8},
					},
				},
				{
					points: []point{
						{x: 5, y: 5}, {x: 6, y: 4}, {x: 7, y: 3}, {x: 8, y: 2},
					},
				},
				{
					points: []point{
						{x: 6, y: 4}, {x: 5, y: 3}, {x: 4, y: 2}, {x: 3, y: 1}, {x: 2, y: 0},
					},
				},
				{
					points: []point{
						{x: 8, y: 0}, {x: 7, y: 1}, {x: 6, y: 2}, {x: 5, y: 3}, {x: 4, y: 4}, {x: 3, y: 5}, {x: 2, y: 6}, {x: 1, y: 7}, {x: 0, y: 8},
					},
				},
			},
			expectedValue: 12,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart1", func(t *testing.T) {
			actual := Part2(tc.input)
			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}
		})
	}
}
