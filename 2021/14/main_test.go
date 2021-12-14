package main

import (
	"fmt"
	"testing"
)

func TestInsertions(t *testing.T) {
	tt := []struct {
		input         string
		expectedValue map[string]int
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
			expectedValue: map[string]int{"B": 1, "BC": 1, "CH": 1, "CN": 1, "HB": 1, "NB": 1},
		},
	}

	for _, tc := range tt {
		t.Run("TestInsertions", func(t *testing.T) {
			chars := make(map[string]int)

			for i := 0; i+1 < len(tc.input); i++ {
				newEntry := string(tc.input[i]) + string(tc.input[i+1])
				chars[newEntry]++
			}
			chars[string(tc.input[len(tc.input)-1])] = 1
			actual := insertions(chars, tc.rules, 1)
			fmt.Println(actual)
			for evKey, ev := range tc.expectedValue {
				if actual[evKey] != ev {
					t.Errorf("Got %d, expected %d", actual[evKey], ev)
				}
			}

		})
	}
}

func TestApplyRules10Times(t *testing.T) {
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
		t.Run("TestApplyRules 10 Times", func(t *testing.T) {
			actual := applyRules(tc.input, tc.rules, 10)
			if actual != tc.expectedValue {
				t.Errorf("Got %v, expected %v", actual, tc.expectedValue)
			}

		})
	}
}

func TestApplyRules(t *testing.T) {
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
			expectedValue: 2188189693529,
		},
	}

	for _, tc := range tt {
		t.Run("TestApplyRules 40 Times", func(t *testing.T) {
			actual := applyRules(tc.input, tc.rules, 40)
			if actual != tc.expectedValue {
				t.Errorf("Got %v, expected %v", actual, tc.expectedValue)
			}
		})
	}
}
