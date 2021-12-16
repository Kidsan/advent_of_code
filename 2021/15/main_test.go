package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		seen map[string]int
		rows [][]int
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				x:    0,
				y:    0,
				seen: map[string]int{},
				rows: [][]int{
					{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
					{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
					{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
					{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
					{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
					{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
					{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
					{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
					{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
					{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
				},
			},
			want: 41,
		},

		{
			name: "convoluted",
			args: args{
				x:    0,
				y:    0,
				seen: map[string]int{},
				rows: [][]int{
					{1, 1, 6, 3},
					{1, 3, 8, 1},
					{2, 1, 3, 6},
					{3, 6, 9, 4},
				},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.x, tt.args.y, tt.args.rows, tt.args.seen); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
