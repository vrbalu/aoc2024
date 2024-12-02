package calculator

import "testing"

func TestCalculateSimiliarity(t *testing.T) {
	type args struct {
		leftArray  []int
		rightArray []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "given result",
			args: args{
				leftArray:  []int{3, 4, 2, 1, 3, 3},
				rightArray: []int{4, 3, 5, 3, 9, 3},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSimiliarity(tt.args.leftArray, tt.args.rightArray); got != tt.want {
				t.Errorf("CalculateSimiliarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
