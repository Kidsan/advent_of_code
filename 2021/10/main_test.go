package main

import (
	"testing"
)

func TestFindFirstError(t *testing.T) {
	tt := []struct {
		input         string
		expectedValue string
	}{
		{
			input:         "{([(<{}[<>[]}>{[]{[(<()>",
			expectedValue: "}",
		},
		{
			input:         "[[<[([]))<([[{}[[()]]]",
			expectedValue: ")",
		},
		{
			input:         "[{[{({}]{}}([{[{{{}}([]",
			expectedValue: "]",
		},
		{
			input:         "[<(<(<(<{}))><([]([]()",
			expectedValue: ")",
		},
		{
			input:         "<{([([[(<>()){}]>(<<{{",
			expectedValue: ">",
		},
		{
			input:         "[({(<(())[]>[[{[]{<()<>>",
			expectedValue: "",
		},
	}

	for _, tc := range tt {
		t.Run("TestPart1", func(t *testing.T) {
			actual := findFirstError(tc.input)

			if actual != tc.expectedValue {
				t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
			}

		})
	}
}

func TestPart1(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue int
	}{
		{
			input: []string{
				"[({(<(())[]>[[{[]{<()<>>",
				"[(()[<>])]({[<{<<[]>>(",
				"{([(<{}[<>[]}>{[]{[(<()>",
				"(((({<>}<{<{<>}{[]{[]{}",
				"[[<[([]))<([[{}[[()]]]",
				"[{[{({}]{}}([{[{{{}}([]",
				"{<[[]]>}<{[{[{[]{()[[[]",
				"[<(<(<(<{}))><([]([]()",
				"<{([([[(<>()){}]>(<<{{",
				"<{([{{}}[<[[[<>{}]]]>[]]",
			},
			expectedValue: 26397,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart2", func(t *testing.T) {
			actual := part1(tc.input)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}
