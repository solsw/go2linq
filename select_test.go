//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectTest.cs

func Test_SelectErr_int_int(t *testing.T) {
	var count int
	type args struct {
		source   Enumerator[int]
		selector func(int) int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
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
				source: NewOnSlice(1, 3, 7, 9, 10),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "SimpleProjection",
			args: args{
				source:   NewOnSlice(1, 5, 2),
				selector: func(x int) int { return x * 2 },
			},
			want: NewOnSlice(2, 10, 4),
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
				source:   NewOnSlice(3, 2, 1), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewOnSlice(1, 2, 3),
		},
		{name: "SideEffectsInProjection2",
			args: args{
				source:   NewOnSlice(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewOnSlice(4, 5, 6),
		},
		{name: "SideEffectsInProjection3",
			args: args{
				source:   NewOnSlice(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewOnSlice(11, 12, 13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectErr(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SelectErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectErr() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
		if tt.name == "SideEffectsInProjection2" {
			count = 10
		}
	}
}

func Test_SelectIdxErr_int_int(t *testing.T) {
	type args struct {
		source   Enumerator[int]
		selector func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
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
				source: NewOnSlice(1, 3, 7, 9, 10),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "WithIndexSimpleProjection",
			args: args{
				source:   NewOnSlice(1, 5, 2),
				selector: func(x, index int) int { return x + index*10 },
			},
			want: NewOnSlice(1, 15, 22),
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
			got, err := SelectIdxErr(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectIdxErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SelectIdxErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectIdxErr() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Select_int_string(t *testing.T) {
	type args struct {
		source   Enumerator[int]
		selector func(int) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "SimpleProjectionToDifferentType",
			args: args{
				source:   NewOnSlice(1, 5, 2),
				selector: func(x int) string { return fmt.Sprint(x) },
			},
			want: NewOnSlice("1", "5", "2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.source, tt.args.selector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Select() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
