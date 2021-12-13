package main

import (
	"testing"
)

func TestFoldY(t *testing.T) {
	tt := []struct {
		input         [][]int
		expectedValue [][]int
		foldLine      int
	}{
		{
			input: [][]int{
				{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			foldLine: 7,
			expectedValue: [][]int{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}

	for _, tc := range tt {
		t.Run("TestFold", func(t *testing.T) {
			actual := foldY(tc.input, tc.foldLine)
			for y, line := range tc.expectedValue {
				for x, value := range line {
					if value != actual[y][x] {
						t.Errorf("Got %d, expected %d", value, tc.expectedValue[y][x])
					}
				}
			}

		})
	}
}

func TestFoldX(t *testing.T) {
	tt := []struct {
		input         [][]int
		expectedValue [][]int
		foldLine      int
	}{
		{
			input: [][]int{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			foldLine: 5,
			expectedValue: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
	}

	for _, tc := range tt {
		t.Run("TestFoldX", func(t *testing.T) {
			actual := foldX(tc.input, tc.foldLine)

			for y, line := range tc.expectedValue {
				for x, value := range line {
					if value != actual[y][x] {
						t.Errorf("Got %d, expected %d", value, tc.expectedValue[y][x])
					}
				}
			}

		})
	}
}

func TestPart1(t *testing.T) {
	tt := []struct {
		input         [][]int
		foldLine      fold
		expectedValue int
	}{
		{
			input: [][]int{
				{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			foldLine:      fold{dimension: "y", line: 7},
			expectedValue: 17,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart1", func(t *testing.T) {
			actual := part1(tc.input, tc.foldLine)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}

// func TestPart2(t *testing.T) {
// 	tt := []struct {
// 		input         []string
// 		expectedValue int
// 	}{
// 		{
// 			input: []string{
// 				"[({(<(())[]>[[{[]{<()<>>",
// 				"[(()[<>])]({[<{<<[]>>(",
// 				"{([(<{}[<>[]}>{[]{[(<()>",
// 				"(((({<>}<{<{<>}{[]{[]{}",
// 				"[[<[([]))<([[{}[[()]]]",
// 				"[{[{({}]{}}([{[{{{}}([]",
// 				"{<[[]]>}<{[{[{[]{()[[[]",
// 				"[<(<(<(<{}))><([]([]()",
// 				"<{([([[(<>()){}]>(<<{{",
// 				"<{([{{}}[<[[[<>{}]]]>[]]",
// 			},
// 			expectedValue: 288957,
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run("TestPart2", func(t *testing.T) {
// 			actual := part2(tc.input)

// 			if actual != tc.expectedValue {
// 				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
// 			}

// 		})
// 	}
// }

// func TestFilterOutInvalidLines(t *testing.T) {
// 	tt := []struct {
// 		input         []string
// 		expectedValue []string
// 	}{
// 		{
// 			input: []string{
// 				"[({(<(())[]>[[{[]{<()<>>",
// 				"[(()[<>])]({[<{<<[]>>(",
// 				"{([(<{}[<>[]}>{[]{[(<()>",
// 				"(((({<>}<{<{<>}{[]{[]{}",
// 				"[[<[([]))<([[{}[[()]]]",
// 				"[{[{({}]{}}([{[{{{}}([]",
// 				"{<[[]]>}<{[{[{[]{()[[[]",
// 				"[<(<(<(<{}))><([]([]()",
// 				"<{([([[(<>()){}]>(<<{{",
// 				"<{([{{}}[<[[[<>{}]]]>[]]",
// 			},
// 			expectedValue: []string{
// 				"[({(<(())[]>[[{[]{<()<>>",
// 				"[(()[<>])]({[<{<<[]>>(",
// 				"(((({<>}<{<{<>}{[]{[]{}",
// 				"{<[[]]>}<{[{[{[]{()[[[]",
// 				"<{([{{}}[<[[[<>{}]]]>[]]",
// 			},
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run("TestFilterOutInvalidLines", func(t *testing.T) {
// 			actual := filterOutInvalidLines(tc.input)

// 			for i, v := range actual {
// 				if v != tc.expectedValue[i] {
// 					t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
// 				}
// 			}

// 		})
// 	}
// }

// func TestFindMissingLineEnding(t *testing.T) {
// 	tt := []struct {
// 		input         string
// 		expectedValue string
// 	}{
// 		{
// 			input:         "[",
// 			expectedValue: "]",
// 		},
// 		{
// 			input:         "[{",
// 			expectedValue: "}]",
// 		},
// 		{
// 			input:         "[({(<(())[]>[[{[]{<()<>>",
// 			expectedValue: "}}]])})]",
// 		},
// 		{
// 			input:         "[(()[<>])]({[<{<<[]>>(",
// 			expectedValue: ")}>]})",
// 		},
// 		{
// 			input:         "(((({<>}<{<{<>}{[]{[]{}",
// 			expectedValue: "}}>}>))))",
// 		},
// 		{
// 			input:         "{<[[]]>}<{[{[{[]{()[[[]",
// 			expectedValue: "]]}}]}]}>",
// 		},
// 		{
// 			input:         "<{([{{}}[<[[[<>{}]]]>[]]",
// 			expectedValue: "])}>",
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run("TestFindMissingLineEnding", func(t *testing.T) {
// 			actual := findMissingLineEnding(tc.input)

// 			if actual != tc.expectedValue {
// 				t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
// 			}
// 		})
// 	}
// }

// func TestGetEndingScore(t *testing.T) {
// 	tt := []struct {
// 		input         string
// 		expectedValue int
// 	}{
// 		{
// 			input:         "}}]])})]",
// 			expectedValue: 288957,
// 		},
// 		{
// 			input:         ")}>]})",
// 			expectedValue: 5566,
// 		},
// 		{
// 			input:         "}}>}>))))",
// 			expectedValue: 1480781,
// 		},
// 		{
// 			input:         "]]}}]}]}>",
// 			expectedValue: 995444,
// 		},
// 		{
// 			input:         "])}>",
// 			expectedValue: 294,
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run("TestGetEndingScore", func(t *testing.T) {
// 			actual := getEndingScore(tc.input)

// 			if actual != tc.expectedValue {
// 				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
// 			}
// 		})
// 	}
// }
