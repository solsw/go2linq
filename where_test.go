//go:build go1.18

package go2linq

import (
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/WhereTest.cs

func Test_Where_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
		predicate func(int) bool
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
				predicate: func(i int) bool { return i > 5 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullPredicateThrowsNullArgumentException",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "SimpleFiltering",
			args: args{
				source:    NewOnSlice(1, 3, 4, 2, 8, 1),
				predicate: func(i int) bool { return i < 4 },
			},
			want: NewOnSlice(1, 3, 2, 1),
		},
		{name: "EmptySource",
			args: args{
				source:    Empty[int](),
				predicate: func(i int) bool { return i > 5 },
			},
			want: Empty[int](),
		},
		{name: "1",
			args: args{
				source:    NewOnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return false },
			},
			want: Empty[int](),
		},
		{name: "2",
			args: args{
				source:    NewOnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return true },
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "3",
			args: args{
				source:    NewOnSlice(1, 2, 3, 4),
				predicate: func(i int) bool { return i%2 == 1 },
			},
			want: NewOnSlice(1, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Where(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Where() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Where() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Where() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Where_string(t *testing.T) {
	type args struct {
		source    Enumerator[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "4",
			args: args{
				source:    NewOnSlice("one", "two", "three", "four", "five"),
				predicate: func(string) bool { return true },
			},
			want: NewOnSlice("one", "two", "three", "four", "five"),
		},
		{name: "5",
			args: args{
				source:    NewOnSlice("one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return strings.HasPrefix(s, "t") },
			},
			want: NewOnSlice("two", "three"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Where(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Where() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_WhereIdx_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
		predicate func(int, int) bool
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
				predicate: func(x, index int) bool { return x > 5 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "WithIndexNullPredicateThrowsNullArgumentException",
			args: args{
				source: NewOnSlice(1, 3, 7, 9, 10),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "WithIndexSimpleFiltering",
			args: args{
				source:    NewOnSlice(1, 3, 4, 2, 8, 1),
				predicate: func(x, index int) bool { return x < index },
			},
			want: NewOnSlice(2, 1),
		},
		{name: "WithIndexEmptySource",
			args: args{
				source:    Empty[int](),
				predicate: func(x, index int) bool { return x < 4 },
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WhereIdx(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("WhereIdx() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("WhereIdx() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("WhereIdx() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_WhereIdx_string(t *testing.T) {
	type args struct {
		source    Enumerator[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "1",
			args: args{
				source:    NewOnSlice("one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) == i },
			},
			want: NewOnSlice("five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := WhereIdx(tt.args.source, tt.args.predicate); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("WhereIdx() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
