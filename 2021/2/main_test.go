package main

import (
	"testing"
)

func TestFinalPosition(t *testing.T) {
	tt := []struct {
		values         []Instruction
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values:         []Instruction{},
			expectedResult: 0,
		},
		{
			values:         []Instruction{},
			expectedResult: 0,
		},
		{
			values: []Instruction{
				Instruction{
					direction: "forward",
					distance:  5,
				},
				Instruction{
					direction: "down",
					distance:  5,
				},
				Instruction{
					direction: "forward",
					distance:  8,
				},
				Instruction{
					direction: "up",
					distance:  3,
				},
				Instruction{
					direction: "down",
					distance:  8,
				},
				Instruction{
					direction: "forward",
					distance:  2,
				},
			},
			expectedResult: 150,
		},
	}

	for _, tc := range tt {
		t.Run("FinalPosition", func(t *testing.T) {
			if actual := FinalPosition(tc.values...); actual != tc.expectedResult {
				t.Fail()
			}
		})
	}
}

func TestAimingSubmarine(t *testing.T) {
	tt := []struct {
		values         []Instruction
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values:         []Instruction{},
			expectedResult: 0,
		},
		{
			values:         []Instruction{},
			expectedResult: 0,
		},
		{
			values: []Instruction{
				Instruction{
					direction: "forward",
					distance:  5,
				},
				Instruction{
					direction: "down",
					distance:  5,
				},
				Instruction{
					direction: "forward",
					distance:  8,
				},
				Instruction{
					direction: "up",
					distance:  3,
				},
				Instruction{
					direction: "down",
					distance:  8,
				},
				Instruction{
					direction: "forward",
					distance:  2,
				},
			},
			expectedResult: 900,
		},
	}

	for _, tc := range tt {
		t.Run("AimingSubmarine", func(t *testing.T) {
			if actual := AimingSubmarine(tc.values...); actual != tc.expectedResult {
				t.Fail()
			}
		})
	}
}
