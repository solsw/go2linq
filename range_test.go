//go:build go1.18

package go2linq

import (
	"math"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RangeTest.cs

func Test_Range(t *testing.T) {
	type args struct {
		start int
		count int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NegativeCount",
			args: args{
				start: 10,
				count: -1,
			},
			wantErr:     true,
			expectedErr: ErrNegativeCount,
		},
		{name: "LargeButValidCount1",
			args: args{
				start: math.MaxInt32,
				count: 1,
			},
			want: NewEnSlice(math.MaxInt32),
		},
		{name: "ValidRange",
			args: args{
				start: 5,
				count: 3,
			},
			want: NewEnSlice(5, 6, 7),
		},
		{name: "NegativeStart",
			args: args{
				start: -2,
				count: 5,
			},
			want: NewEnSlice(-2, -1, 0, 1, 2),
		},
		{name: "EmptyRange",
			args: args{
				start: 100,
				count: 0,
			},
			want: Empty[int](),
		},
		{name: "SingleValueOfMaxInt32",
			args: args{
				start: math.MaxInt32,
				count: 1,
			},
			want: NewEnSlice(math.MaxInt32),
		},
		{name: "EmptyRangeStartingAtMinInt32",
			args: args{
				start: math.MinInt32,
				count: 0,
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Range(tt.args.start, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Range() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Range() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Range() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}
