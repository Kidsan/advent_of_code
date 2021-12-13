package main

import (
	"testing"
)

func TestGetPaths(t *testing.T) {
	tt := []struct {
		input                 map[string][]string
		expectedValue         [][]string
		usedFreeSmallCaveMove bool
	}{
		{
			input: map[string][]string{
				"start": {"A", "b"},
				"A":     {"start", "b", "c", "end"},
				"b":     {"start", "d", "A", "end"},
				"c":     {"A"},
				"d":     {"b"},
				"end":   {"A", "b"},
			},
			usedFreeSmallCaveMove: true,
			expectedValue: [][]string{
				{"start", "A", "b", "A", "c", "end"},
				{"start", "A", "b", "A", "end"},
				{"start", "A", "b", "end"},
				{"start", "A", "c", "A", "b", "A", "end"},
				{"start", "A", "c", "A", "b", "end"},
				{"start", "A", "c", "A", "end"},
				{"start", "A", "end"},
				{"start", "b", "A", "c", "A", "end"},
				{"start", "b", "A", "end"},
				{"start", "b", "end"},
			},
		},
		{
			input: map[string][]string{
				"start": {"A", "b"},
				"A":     {"start", "b", "c", "end"},
				"b":     {"start", "d", "A", "end"},
				"c":     {"A"},
				"d":     {"b"},
				"end":   {"A", "b"},
			},
			usedFreeSmallCaveMove: false,
			expectedValue: [][]string{
				{"start", "A", "b", "A", "b", "A", "c", "A", "end"},
				{"start", "A", "b", "A", "b", "A", "end"},
				{"start", "A", "b", "A", "b", "end"},
				{"start", "A", "b", "A", "c", "A", "b", "A", "end"},
				{"start", "A", "b", "A", "c", "A", "b", "end"},
				{"start", "A", "b", "A", "c", "A", "c", "A", "end"},
				{"start", "A", "b", "A", "c", "A", "end"},
				{"start", "A", "b", "A", "end"},
				{"start", "A", "b", "d", "b", "A", "c", "A", "end"},
				{"start", "A", "b", "d", "b", "A", "end"},
				{"start", "A", "b", "d", "b", "end"},
				{"start", "A", "b", "end"},
				{"start", "A", "c", "A", "b", "A", "b", "A", "end"},
				{"start", "A", "c", "A", "b", "A", "b", "end"},
				{"start", "A", "c", "A", "b", "A", "c", "A", "end"},
				{"start", "A", "c", "A", "b", "A", "end"},
				{"start", "A", "c", "A", "b", "d", "b", "A", "end"},
				{"start", "A", "c", "A", "b", "d", "b", "end"},
				{"start", "A", "c", "A", "b", "end"},
				{"start", "A", "c", "A", "c", "A", "b", "A", "end"},
				{"start", "A", "c", "A", "c", "A", "b", "end"},
				{"start", "A", "c", "A", "c", "A", "end"},
				{"start", "A", "c", "A", "end"},
				{"start", "A", "end"},
				{"start", "b", "A", "b", "A", "c", "A", "end"},
				{"start", "b", "A", "b", "A", "end"},
				{"start", "b", "A", "b", "end"},
				{"start", "b", "A", "c", "A", "b", "A", "end"},
				{"start", "b", "A", "c", "A", "b", "end"},
				{"start", "b", "A", "c", "A", "c", "A", "end"},
				{"start", "b", "A", "c", "A", "end"},
				{"start", "b", "A", "end"},
				{"start", "b", "d", "b", "A", "c", "A", "end"},
				{"start", "b", "d", "b", "A", "end"},
				{"start", "b", "d", "b", "end"},
				{"start", "b", "end"},
			},
		},
	}

	for _, tc := range tt {
		t.Run("TestGetPaths", func(t *testing.T) {
			result := make([][]string, 0)
			path := []string{"start"}
			getPaths("start", path, tc.input, &result, tc.usedFreeSmallCaveMove)

			if len(result) != len(tc.expectedValue) {
				t.Errorf("Got %d paths, expected %d", len(result), len(tc.expectedValue))
			}

		})
	}
}
