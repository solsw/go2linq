//go:build go1.18

package go2linq

import (
	"fmt"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AggregateTest.cs

func Test_Aggregate_int(t *testing.T) {
	type args struct {
		source      Enumerable[int]
		accumulator func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceUnseeded",
			args: args{
				accumulator: func(x, y int) int { return x + y },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullFuncUnseeded",
			args: args{
				source: NewEnSlice(1, 3),
			},
			wantErr:     true,
			expectedErr: ErrNilAccumulator,
		},
		{name: "UnseededAggregation",
			args: args{
				source:      NewEnSlice(1, 4, 5),
				accumulator: func(current, value int) int { return current*2 + value },
			},
			want: 17,
		},
		{name: "EmptySequenceUnseeded",
			args: args{
				source:      Empty[int](),
				accumulator: func(ac, el int) int { return ac + el },
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "UnseededSingleElementAggregation",
			args: args{
				source:      NewEnSlice(1),
				accumulator: func(ac, el int) int { return ac*2 + el },
			},
			want: 1,
		},
		{name: "FirstElementOfInputIsUsedAsSeedForUnseededOverload",
			args: args{
				source:      NewEnSlice(5, 3, 2),
				accumulator: func(ac, el int) int { return ac * el },
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Aggregate(tt.args.source, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aggregate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Aggregate() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AggregateSeed_int(t *testing.T) {
	type args struct {
		source      Enumerable[int]
		seed        int
		accumulator func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceSeeded",
			args: args{
				accumulator: func(x, y int) int { return x + y },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullFuncSeeded",
			args: args{
				source: NewEnSlice(1, 3),
				seed:   5,
			},
			wantErr:     true,
			expectedErr: ErrNilAccumulator,
		},
		{name: "SeededAggregation",
			args: args{
				source:      NewEnSlice(1, 4, 5),
				seed:        5,
				accumulator: func(current, value int) int { return current*2 + value },
			},
			want: 57,
		},
		{name: "EmptySequenceSeeded",
			args: args{
				source:      Empty[int](),
				seed:        5,
				accumulator: func(x, y int) int { return x + y },
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateSeed(tt.args.source, tt.args.seed, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AggregateSeed() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AggregateSeedMust_int32_int64(t *testing.T) {
	type args struct {
		source      Enumerable[int32]
		seed        int64
		accumulator func(int64, int32) int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "DifferentSourceAndAccumulatorTypes",
			args: args{
				source:      NewEnSlice(int32(2000000000), int32(2000000000), int32(2000000000)),
				seed:        int64(0),
				accumulator: func(ac int64, el int32) int64 { return ac + int64(el) },
			},
			want: int64(6000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AggregateSeedMust(tt.args.source, tt.args.seed, tt.args.accumulator)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeedMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AggregateSeedSel_int_string(t *testing.T) {
	type args struct {
		source         Enumerable[int]
		seed           int
		accumulator    func(int, int) int
		resultSelector func(int) string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceSeededWithResultSelector",
			args: args{
				seed:           5,
				accumulator:    func(x, y int) int { return x + y },
				resultSelector: func(result int) string { return fmt.Sprint(result) },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullFuncSeededWithResultSelector",
			args: args{
				source:         NewEnSlice(1, 3),
				seed:           5,
				resultSelector: func(result int) string { return fmt.Sprint(result) },
			},
			wantErr:     true,
			expectedErr: ErrNilAccumulator,
		},
		{name: "NullProjectionSeededWithResultSelector",
			args: args{
				source:      NewEnSlice(1, 3),
				seed:        5,
				accumulator: func(x, y int) int { return x + y },
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "SeededAggregationWithResultSelector",
			args: args{
				source:         NewEnSlice(1, 4, 5),
				seed:           5,
				accumulator:    func(current, value int) int { return current*2 + value },
				resultSelector: func(result int) string { return fmt.Sprint(result) },
			},
			want: "57",
		},
		{name: "EmptySequenceSeededWithResultSelector",
			args: args{
				source:         Empty[int](),
				seed:           5,
				accumulator:    func(x, y int) int { return x + y },
				resultSelector: func(result int) string { return fmt.Sprint(result) },
			},
			want: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateSeedSel(tt.args.source, tt.args.seed, tt.args.accumulator, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSeedSel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AggregateSeedSel() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeedSel() = %v, want %v", got, tt.want)
			}
		})
	}
}
