package counter

import "testing"

func TestCountSafeReports(t *testing.T) {
	type args struct {
		allReports [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "given example", args: args{allReports: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9}}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSafeReports(tt.args.allReports); got != tt.want {
				t.Errorf("CountSafeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}
