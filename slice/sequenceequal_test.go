//go:build go1.18

package slice

import (
	"testing"
)

func TestSequenceEqualMust_int(t *testing.T) {
	r0 := []int{}
	r2 := []int{0, 1}
	type args struct {
		first  []int
		second []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NilNil",
			args: args{
				first:  nil,
				second: nil,
			},
			want: true,
		},
		{name: "NilEmpty",
			args: args{
				first:  nil,
				second: []int{},
			},
			want: true,
		},
		{name: "NilFirst",
			args: args{
				first:  nil,
				second: []int{1},
			},
			want: false,
		},
		{name: "EmptyEmpty",
			args: args{
				first:  []int{},
				second: []int{},
			},
			want: true,
		},
		{name: "EmptySecond",
			args: args{
				first:  []int{1},
				second: []int{},
			},
			want: false,
		},
		{name: "UnequalLengthsBothArrays",
			args: args{
				first:  []int{1, 5, 3},
				second: []int{1, 5, 3, 10},
			},
			want: false,
		},
		{name: "UnequalLengths",
			args: args{
				first:  []int{1, 5, 3},
				second: []int{1, 5, 3, 10},
			},
			want: false,
		},
		{name: "UnequalData",
			args: args{
				first:  []int{1, 5, 3, 9},
				second: []int{1, 5, 3, 10},
			},
			want: false,
		},
		{name: "EqualSlices",
			args: args{
				first:  []int{1, 5, 3, 10},
				second: []int{1, 5, 3, 10},
			},
			want: true,
		},
		{name: "OrderMatters",
			args: args{
				first:  []int{1, 2},
				second: []int{2, 1},
			},
			want: false,
		},
		{name: "Same0",
			args: args{
				first:  r0,
				second: r0,
			},
			want: true,
		},
		{name: "Same2",
			args: args{
				first:  r2,
				second: r2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SequenceEqualMust(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("SequenceEqualMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
