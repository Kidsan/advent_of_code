package main

import (
	"testing"
)

func TestInsertions(t *testing.T) {
	tt := []struct {
		input         string
		expectedValue string
		rules         []rule
	}{
		{
			input: "NNCB",
			rules: []rule{
				{
					matches: "CH",
					inserts: "B",
				},
				{
					matches: "HH",
					inserts: "N",
				},
				{
					matches: "CB",
					inserts: "H",
				},
				{
					matches: "NH",
					inserts: "C",
				},
				{
					matches: "HB",
					inserts: "C",
				},
				{
					matches: "HC",
					inserts: "B",
				},
				{
					matches: "HN",
					inserts: "C",
				},
				{
					matches: "NN",
					inserts: "C",
				},
				{
					matches: "BH",
					inserts: "H",
				},
				{
					matches: "NC",
					inserts: "B",
				},
				{
					matches: "NB",
					inserts: "B",
				},
				{
					matches: "BN",
					inserts: "B",
				},
				{
					matches: "BB",
					inserts: "N",
				},
				{
					matches: "BC",
					inserts: "B",
				},
				{
					matches: "CC",
					inserts: "N",
				},
				{
					matches: "CN",
					inserts: "C",
				},
			},
			expectedValue: "NCNBCHB",
		},
	}

	for _, tc := range tt {
		t.Run("TestInsertions", func(t *testing.T) {
			inputList := InputString{}
			for i := len(tc.input) - 1; i >= 0; i-- {
				inputList.add(string(tc.input[i]))
			}
			actual := insertions(&inputList, tc.rules, 1)

			if actual != tc.expectedValue {
				t.Errorf("Got %s, expected %s", actual, tc.expectedValue)
			}

		})
	}
}

func TestPart1(t *testing.T) {
	tt := []struct {
		input         string
		rules         []rule
		expectedValue int
	}{
		{
			input: "NNCB",
			rules: []rule{
				{
					matches: "CH",
					inserts: "B",
				},
				{
					matches: "HH",
					inserts: "N",
				},
				{
					matches: "CB",
					inserts: "H",
				},
				{
					matches: "NH",
					inserts: "C",
				},
				{
					matches: "HB",
					inserts: "C",
				},
				{
					matches: "HC",
					inserts: "B",
				},
				{
					matches: "HN",
					inserts: "C",
				},
				{
					matches: "NN",
					inserts: "C",
				},
				{
					matches: "BH",
					inserts: "H",
				},
				{
					matches: "NC",
					inserts: "B",
				},
				{
					matches: "NB",
					inserts: "B",
				},
				{
					matches: "BN",
					inserts: "B",
				},
				{
					matches: "BB",
					inserts: "N",
				},
				{
					matches: "BC",
					inserts: "B",
				},
				{
					matches: "CC",
					inserts: "N",
				},
				{
					matches: "CN",
					inserts: "C",
				},
			},
			expectedValue: 1588,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart1", func(t *testing.T) {
			actual := part1(tc.input, tc.rules)
			if actual != tc.expectedValue {
				t.Errorf("Got %v, expected %v", actual, tc.expectedValue)
			}

		})
	}
}

// func TestFoldX(t *testing.T) {
// 	tt := []struct {
// 		input         [][]int
// 		expectedValue [][]int
// 		foldLine      int
// 	}{
// 		{
// 			input: [][]int{
// 				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0},
// 				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
// 				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
// 				{0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1},
// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 			},
// 			foldLine: 5,
// 			expectedValue: [][]int{
// 				{1, 1, 1, 1, 1},
// 				{1, 0, 0, 0, 1},
// 				{1, 0, 0, 0, 1},
// 				{1, 0, 0, 0, 1},
// 				{1, 1, 1, 1, 1},
// 				{0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0},
// 			},
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run("TestFoldX", func(t *testing.T) {
// 			actual := foldX(tc.input, tc.foldLine)

// 			for y, line := range tc.expectedValue {
// 				for x, value := range line {
// 					if value != actual[y][x] {
// 						t.Errorf("Got %d, expected %d", value, tc.expectedValue[y][x])
// 					}
// 				}
// 			}

// 		})
// 	}
// }

// func TestPart1(t *testing.T) {
// 	tt := []struct {
// 		input         [][]int
// 		foldLine      fold
// 		expectedValue int
// 	}{
// 		{
// 			input: [][]int{
// 				{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
// 				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 				{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1},
// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 				{0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0},
// 				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
// 				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 				{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
// 			},
// 			foldLine:      fold{dimension: "y", line: 7},
// 			expectedValue: 17,
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run("TestPart1", func(t *testing.T) {
// 			actual := part1(tc.input, tc.foldLine)

// 			if actual != tc.expectedValue {
// 				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
// 			}

// 		})
// 	}
// }

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
