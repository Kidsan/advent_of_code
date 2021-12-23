package main

import (
	"reflect"
	"testing"
)

func Test_parseSnailfishNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want snailfishNumber
	}{
		{
			name: "Basic Number",
			args: args{
				input: "[1,2]",
			},
			want: snailfishNumber{
				value: 1,
				depth: 1,
				next: &snailfishNumber{
					value: 2,
					depth: 1,
				},
			},
		},
		{
			name: "Depth of 2",
			args: args{
				input: "[[1,2],3]",
			},
			want: snailfishNumber{
				value: 1,
				depth: 2,
				next: &snailfishNumber{
					value: 2,
					depth: 2,
					next: &snailfishNumber{
						depth: 1,
						value: 3,
					},
				},
			},
		},
		{
			name: "Complicated",
			args: args{
				input: "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			},
			want: snailfishNumber{
				value: 0,
				depth: 3,
				next: &snailfishNumber{
					value: 4,
					depth: 4,
					next: &snailfishNumber{
						depth: 4,
						value: 5,
						next: &snailfishNumber{
							value: 0,
							depth: 3,
							next: &snailfishNumber{
								value: 0,
								depth: 3,
								next: &snailfishNumber{
									value: 4,
									depth: 4,
									next: &snailfishNumber{
										value: 5,
										depth: 4,
										next: &snailfishNumber{
											value: 2,
											depth: 4,
											next: &snailfishNumber{
												value: 6,
												depth: 4,
												next: &snailfishNumber{
													value: 9,
													depth: 3,
													next: &snailfishNumber{
														value: 5,
														depth: 3,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseSnailfishNumber(tt.args.input)

			if !reflect.DeepEqual(got.value, tt.want.value) {
				t.Errorf("got.value() = %v, want %v", got.value, tt.want.value)
			}
			if !reflect.DeepEqual(got.depth, tt.want.depth) {
				t.Errorf("got.depth() = %v, want %v", got.depth, tt.want.depth)
			}

			received := got.next
			next := tt.want.next

			for next != nil {

				if !reflect.DeepEqual(next.value, received.value) {
					t.Errorf("got.next.value() = %v, want %v", next.value, received.value)
				}
				if !reflect.DeepEqual(next.depth, received.depth) {
					t.Errorf("got.next.depth() = %v, want %v", next.depth, received.depth)
				}

				next = next.next
				received = received.next
			}
		})
	}
}

func Test_snailfishNumber_reduce(t *testing.T) {
	tests := []struct {
		want *snailfishNumber
		name string
		args []string
	}{
		{
			name: "Example From Description",
			args: []string{
				"[[[[4,3],4],4],[7,[[8,4],9]]]",
				"[1,1]",
			},
			want: &snailfishNumber{
				prev: &snailfishNumber{},
				next: &snailfishNumber{
					prev: &snailfishNumber{},
					next: &snailfishNumber{
						prev: &snailfishNumber{},
						next: &snailfishNumber{
							prev: &snailfishNumber{},
							next: &snailfishNumber{
								prev: &snailfishNumber{},
								next: &snailfishNumber{
									prev: &snailfishNumber{},
									next: &snailfishNumber{
										prev: &snailfishNumber{},
										next: &snailfishNumber{
											prev: &snailfishNumber{},
											next: &snailfishNumber{
												prev: &snailfishNumber{},
												next: &snailfishNumber{
													prev:  &snailfishNumber{},
													next:  nil,
													value: 1,
													depth: 2,
												},
												value: 8,
												depth: 2,
											},
											value: 0,
											depth: 4,
										},
										value: 6,
										depth: 4,
									},
									value: 8,
									depth: 4,
								},
								value: 7,
								depth: 4,
							},
							value: 4,
							depth: 3,
						},
						value: 7,
						depth: 4,
					},
					value: 3,
					depth: 5,
				},
				value: 4,
				depth: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numbers []*snailfishNumber
			for _, num := range tt.args {
				newNum := parseSnailfishNumber(num)
				numbers = append(numbers, newNum)
			}
			got := numbers[0].reduce()
			for i := 1; i < len(numbers); i++ {
				got.add(numbers[i])
			}

			received := got
			next := tt.want

			for next != nil {

				if !reflect.DeepEqual(next.value, received.value) {
					t.Errorf("got.next.value() = %v, want %v", received.value, next.value)
				}
				if !reflect.DeepEqual(next.depth, received.depth) {
					t.Errorf("got.next.depth() = %v, want %v", received.depth, next.depth)
				}

				next = next.next
				received = received.next
			}
		})
	}
}

func Test_snailfishNumber_magnitude(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Example 1",
			args: "[[1,2],[[3,4],5]]",
			want: 143,
		},
		{
			name: "Example 2",
			args: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			want: 1384,
		},
		{
			name: "Example 3",
			args: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			want: 445,
		},
		{
			name: "Example 4",
			args: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			want: 791,
		},
		{
			name: "Example 5",
			args: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			want: 1137,
		},
		{
			name: "Example 6",
			args: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			want: 3488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := parseSnailfishNumber(tt.args)
			if got := n.magnitude(); got != tt.want {
				t.Errorf("snailfishNumber.magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "Example",
			args: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want: 4140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numbers []*snailfishNumber

			for _, num := range tt.args {
				numbers = append(numbers, parseSnailfishNumber(num))
			}
			if got := part1(numbers); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
