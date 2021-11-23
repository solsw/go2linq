//go:build go1.18

package go2linq

import (
	"math"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RangeTest.cs

func Test_RangeErr(t *testing.T) {
	type args struct {
		start int
		count int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
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
		{name: "CountTooLarge1",
			args: args{
				start: math.MaxInt32,
				count: 2,
			},
			wantErr:     true,
			expectedErr: ErrStartCount,
		},
		{name: "CountTooLarge2",
			args: args{
				start: 2,
				count: math.MaxInt32,
			},
			wantErr:     true,
			expectedErr: ErrStartCount,
		},
		{name: "CountTooLarge3",
			args: args{
				start: math.MaxInt32 / 2,
				count: math.MaxInt32/2 + 3,
			},
			wantErr:     true,
			expectedErr: ErrStartCount,
		},
		{name: "LargeButValidCount1",
			args: args{
				start: math.MaxInt32,
				count: 1,
			},
			want: NewOnSlice(math.MaxInt32),
		},
		{name: "ValidRange",
			args: args{
				start: 5,
				count: 3,
			},
			want: NewOnSlice(5, 6, 7),
		},
		{name: "NegativeStart",
			args: args{
				start: -2,
				count: 5,
			},
			want: NewOnSlice(-2, -1, 0, 1, 2),
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
			want: NewOnSlice(math.MaxInt32),
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
			got, err := RangeErr(tt.args.start, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("RangeErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("RangeErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("RangeErr() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

/*
func Test_Range_LargeButValidCount(t *testing.T) {
	// max length of Enumerator depends on available memory
	t.Log(Count(Range(1, math.MaxInt32)))
	t.Log(Count(Range(math.MaxInt32 / 2, math.MaxInt32 / 2 + 2)))
}
*/
