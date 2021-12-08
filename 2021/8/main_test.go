package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		input         []entry
		expectedValue int
	}{
		{
			input: []entry{
				{
					signal: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"},
					output: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
				},
				{
					signal: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"},
					output: []string{"fcgedb", "cgb", "dgebacf", "gc"},
				},
				{
					signal: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"},
					output: []string{"cg", "cg", "fdcagb", "cbg"},
				},
				{
					signal: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"},
					output: []string{"efabcd", "cedba", "gadfec", "cb"},
				},
				{
					signal: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fc", "ge", "gcbea", "fcaegb", "dgceab", "fcbdga"},
					output: []string{"gecf", "egdcabf", "bgf", "bfgea"},
				},
				{
					signal: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gc", "fdb", "baec", "bfadeg", "bafgc", "acf"},
					output: []string{"gebdcfa", "ecba", "ca", "fadegcb"},
				},
				{
					signal: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ec", "dfab", "fbedc", "dacgb", "gdcebf", "gf"},
					output: []string{"cefg", "dcbef", "fcge", "gbcadfe"},
				},
				{
					signal: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"},
					output: []string{"ed", "bcgafe", "cdgba", "cbgef"},
				},
				{
					signal: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gb", "defca", "cg", "fgcdab", "egfdb", "bfceg"},
					output: []string{"gbdfcae", "bgc", "cg", "cgb"},
				},
				{
					signal: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "ab", "cdeg", "gaef", "cafbge", "fdbac", "fegbdc"},
					output: []string{"fgae", "cfgab", "fg", "bagce"},
				},
			},
			expectedValue: 26,
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

func TestParseSignals(t *testing.T) {
	tt := []struct {
		expectedValue map[string]string
		input         []string
	}{
		{
			input:         []string{"fbcad", "acedgfb", "cdfbe", "gcdfa", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			expectedValue: map[string]string{"ab": "1", "abcdef": "9", "abcdefg": "8", "abcdeg": "0", "abcdf": "3", "abd": "7", "abef": "4", "acdfg": "2", "bcdef": "5", "bcdefg": "6"},
		},
	}

	for _, tc := range tt {
		t.Run("TestParseSignals", func(t *testing.T) {
			actual := parseSignals(tc.input)

			for index, num := range tc.expectedValue {
				if num != actual[index] {
					t.Errorf("Got %s, expected %s", actual, tc.expectedValue)

				}
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		input         []entry
		expectedValue int
	}{
		{
			input: []entry{
				{
					signal: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
					output: []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
				},
			},
			expectedValue: 5353,
		},
		{
			input: []entry{
				{
					signal: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, output: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
				},
				{
					signal: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, output: []string{"fcgedb", "cgb", "dgebacf", "gc"},
				},
				{
					signal: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"}, output: []string{"cg", "cg", "fdcagb", "cbg"},
				},
				{
					signal: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"}, output: []string{"efabcd", "cedba", "gadfec", "cb"},
				},
				{
					signal: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"}, output: []string{"gecf", "egdcabf", "bgf", "bfgea"},
				},
				{
					signal: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"}, output: []string{"gebdcfa", "ecba", "ca", "fadegcb"},
				},
				{
					signal: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"}, output: []string{"cefg", "dcbef", "fcge", "gbcadfe"},
				},
				{
					signal: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, output: []string{"ed", "bcgafe", "cdgba", "cbgef"},
				},
				{
					signal: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"}, output: []string{"gbdfcae", "bgc", "cg", "cgb"},
				},
				{
					signal: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, output: []string{"fgae", "cfgab", "fg", "bagce"},
				},
			},
			expectedValue: 61229,
		},
	}

	for _, tc := range tt {
		t.Run("TestPart2", func(t *testing.T) {
			sortedInput := make([]entry, 0)
			for _, e := range tc.input {
				newEntry := entry{}
				for _, word := range e.signal {
					newEntry.signal = append(newEntry.signal, SortString(word))
				}
				for _, word := range e.output {
					newEntry.output = append(newEntry.output, SortString(word))
				}
				sortedInput = append(sortedInput, newEntry)
			}
			fmt.Println(sortedInput)
			actual := part2(sortedInput)

			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}

		})
	}
}
