//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectTest.cs

func Test_Select_int_int(t *testing.T) {
	var count int
	type args struct {
		source   Enumerable[int]
		selector func(int) int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceThrowsNullArgumentException",
			args: args{
				selector: func(x int) int { return x + 1 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullProjectionThrowsNullArgumentException",
			args: args{
				source: NewEnSlice(1, 3, 7, 9, 10),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "SimpleProjection",
			args: args{
				source:   NewEnSlice(1, 5, 2),
				selector: func(x int) int { return x * 2 },
			},
			want: NewEnSlice(2, 10, 4),
		},
		{name: "EmptySource",
			args: args{
				source:   Empty[int](),
				selector: func(x int) int { return x * 2 },
			},
			want: Empty[int](),
		},
		{name: "SideEffectsInProjection1",
			args: args{
				source:   NewEnSlice(3, 2, 1), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "SideEffectsInProjection2",
			args: args{
				source:   NewEnSlice(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewEnSlice(4, 5, 6),
		},
		{name: "SideEffectsInProjection3",
			args: args{
				source:   NewEnSlice(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewEnSlice(11, 12, 13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Select() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Select() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
		if tt.name == "SideEffectsInProjection2" {
			count = 10
		}
	}
}

func Test_Select_int_string(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SimpleProjectionToDifferentType",
			args: args{
				source:   NewEnSlice(1, 5, 2),
				selector: func(x int) string { return fmt.Sprint(x) },
			},
			want: NewEnSlice("1", "5", "2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Select(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Select() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_SelectIdx_int_int(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "WithIndexNullSourceThrowsNullArgumentException",
			args: args{
				selector: func(x, index int) int { return x + index },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "WithIndexNullPredicateThrowsNullArgumentException",
			args: args{
				source: NewEnSlice(1, 3, 7, 9, 10),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "WithIndexSimpleProjection",
			args: args{
				source:   NewEnSlice(1, 5, 2),
				selector: func(x, index int) int { return x + index*10 },
			},
			want: NewEnSlice(1, 15, 22),
		},
		{name: "WithIndexEmptySource",
			args: args{
				source:   Empty[int](),
				selector: func(x, index int) int { return x + index },
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectIdx(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectIdx() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SelectIdx() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectIdx() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}
