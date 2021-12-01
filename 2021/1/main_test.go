package main

import (
	"testing"
)

func TestDepthComparison(t *testing.T) {
	tt := []struct {
		values         []int
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values:         []int{},
			expectedResult: 0,
		},
		{
			values:         []int{1},
			expectedResult: 0,
		},
		{
			values:         []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expectedResult: 7,
		},
	}

	for _, tc := range tt {
		t.Run("DepthComparison", func(t *testing.T) {
			if actual := DepthComparison(tc.values...); actual != tc.expectedResult {
				t.Fail()
			}
		})
	}
}

func TestSlidingWindowComparison(t *testing.T) {
	tt := []struct {
		values         []int
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values:         []int{},
			expectedResult: 0,
		},
		{
			values:         []int{1},
			expectedResult: 0,
		},
		{
			values:         []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expectedResult: 5,
		},
	}

	for _, tc := range tt {
		t.Run("DepthComparison", func(t *testing.T) {
			if actual := SlidingWindowComparison(tc.values...); actual != tc.expectedResult {
				t.Fail()
			}
		})
	}
}
