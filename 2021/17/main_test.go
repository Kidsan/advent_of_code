package main

import (
	"testing"
)

func Test_ApplyVelocity(t *testing.T) {
	type args struct {
		velocityX int
		velocityY int
	}
	tests := []struct {
		name string
		args args
		want probe
	}{
		{
			name: "example",
			args: args{
				velocityX: 7,
				velocityY: 2,
			},
			want: probe{
				posX:      7,
				posY:      2,
				velocityX: 6,
				velocityY: 1,
			},
		},
		{
			name: "example",
			args: args{
				velocityX: 17,
				velocityY: -4,
			},
			want: probe{
				posX:      17,
				posY:      -4,
				velocityX: 16,
				velocityY: -5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := probe{
				posX:      0,
				posY:      0,
				velocityX: tt.args.velocityX,
				velocityY: tt.args.velocityY,
			}
			p.ApplyVelocity()
			got := p
			if got.posX != tt.want.posX {
				t.Errorf("ApplyVelocity() = %v, want %v", got.posX, tt.want.posX)
			}
			if got.posY != tt.want.posY {
				t.Errorf("ApplyVelocity() = %v, want %v", got.posY, tt.want.posY)
			}
			if got.velocityX != tt.want.velocityX {
				t.Errorf("ApplyVelocity() = %v, want %v", got.velocityX, tt.want.velocityX)
			}
			if got.velocityY != tt.want.velocityY {
				t.Errorf("ApplyVelocity() = %v, want %v", got.velocityY, tt.want.velocityY)
			}
		})
	}
}
