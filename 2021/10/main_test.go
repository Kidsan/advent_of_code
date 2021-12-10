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
		t.Run("TestPart1", func(t *testing.T) {
			actual := part1(tc.input)

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
			expectedValue: 288957,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart2", func(t *testing.T) {
			actual := part2(tc.input)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}

func TestFilterOutInvalidLines(t *testing.T) {
	tt := []struct {
		input         []string
		expectedValue []string
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
			expectedValue: []string{
				"[({(<(())[]>[[{[]{<()<>>",
				"[(()[<>])]({[<{<<[]>>(",
				"(((({<>}<{<{<>}{[]{[]{}",
				"{<[[]]>}<{[{[{[]{()[[[]",
				"<{([{{}}[<[[[<>{}]]]>[]]",
			},
		},
	}

	for _, tc := range tt {
		t.Run("TestFilterOutInvalidLines", func(t *testing.T) {
			actual := filterOutInvalidLines(tc.input)

			for i, v := range actual {
				if v != tc.expectedValue[i] {
					t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
				}
			}

		})
	}
}

func TestFindMissingLineEnding(t *testing.T) {
	tt := []struct {
		input         string
		expectedValue string
	}{
		{
			input:         "[",
			expectedValue: "]",
		},
		{
			input:         "[{",
			expectedValue: "}]",
		},
		{
			input:         "[({(<(())[]>[[{[]{<()<>>",
			expectedValue: "}}]])})]",
		},
		{
			input:         "[(()[<>])]({[<{<<[]>>(",
			expectedValue: ")}>]})",
		},
		{
			input:         "(((({<>}<{<{<>}{[]{[]{}",
			expectedValue: "}}>}>))))",
		},
		{
			input:         "{<[[]]>}<{[{[{[]{()[[[]",
			expectedValue: "]]}}]}]}>",
		},
		{
			input:         "<{([{{}}[<[[[<>{}]]]>[]]",
			expectedValue: "])}>",
		},
	}

	for _, tc := range tt {
		t.Run("TestFindMissingLineEnding", func(t *testing.T) {
			actual := findMissingLineEnding(tc.input)

			if actual != tc.expectedValue {
				t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
			}
		})
	}
}

func TestGetEndingScore(t *testing.T) {
	tt := []struct {
		input         string
		expectedValue int
	}{
		{
			input:         "}}]])})]",
			expectedValue: 288957,
		},
		{
			input:         ")}>]})",
			expectedValue: 5566,
		},
		{
			input:         "}}>}>))))",
			expectedValue: 1480781,
		},
		{
			input:         "]]}}]}]}>",
			expectedValue: 995444,
		},
		{
			input:         "])}>",
			expectedValue: 294,
		},
	}

	for _, tc := range tt {
		t.Run("TestGetEndingScore", func(t *testing.T) {
			actual := getEndingScore(tc.input)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}
		})
	}
}
