//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/IntersectTest.cs

func Test_IntersectMust_int(t *testing.T) {
	e1 := NewEnSlice(1, 2, 3, 4)
	e2 := NewEnSlice(1, 2, 3, 4)
	e3 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1",
			args: args{
				first:  NewEnSlice(1, 2),
				second: NewEnSlice(2, 3),
			},
			want: NewEnSlice(2),
		},
		{name: "IntWithoutComparer",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: NewEnSlice(4, 5, 6, 7, 8),
		},
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  e2,
				second: SkipMust(e2, 1),
			},
			want: NewEnSlice(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 3),
				second: e3,
			},
			want: NewEnSlice(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_IntersectMust_string(t *testing.T) {
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "NoComparerSpecified",
			args: args{
				first:  NewEnSlice("A", "a", "b", "c", "b"),
				second: NewEnSlice("b", "a", "d", "a"),
			},
			want: NewEnSlice("a", "b"),
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#intersect-and-intersectby
		{name: "Intersect",
			args: args{
				first:  NewEnSlice("Mercury", "Venus", "Earth", "Jupiter"),
				second: NewEnSlice("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: NewEnSlice("Mercury", "Earth", "Jupiter"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_IntersectEqMust_int(t *testing.T) {
	type args struct {
		first   Enumerable[int]
		second  Enumerable[int]
		equaler Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:   NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second:  NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				equaler: Order[int]{}},
			want: NewEnSlice(4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectEqMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_IntersectEqMust_string(t *testing.T) {
	type args struct {
		first   Enumerable[string]
		second  Enumerable[string]
		equaler Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:   NewEnSlice("A", "a", "b", "c", "b"),
				second:  NewEnSlice("b", "a", "d", "a"),
				equaler: CaseInsensitiveEqualer,
			},
			want: NewEnSlice("A", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectEqMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_IntersectCmpMust_int(t *testing.T) {
	e1 := NewEnSlice(4, 3, 2, 1)
	e2 := NewEnSlice(1, 2, 3, 4)
	e3 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
		cmp    Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				cmp:    Order[int]{},
			},
			want: NewEnSlice(4, 5, 6, 7, 8),
		},
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
				cmp:    Order[int]{},
			},
			want: NewEnSlice(4, 3, 2, 1),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  e2,
				second: SkipMust(e2, 1),
				cmp:    Order[int]{},
			},
			want: NewEnSlice(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 3),
				second: e3,
				cmp:    Order[int]{},
			},
			want: NewEnSlice(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectCmpMust(tt.args.first, tt.args.second, tt.args.cmp)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectCmpMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_IntersectCmpMust_string(t *testing.T) {
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
		cmp    Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:  NewEnSlice("A", "a", "b", "c", "b"),
				second: NewEnSlice("b", "a", "d", "a"),
				cmp:    CaseInsensitiveComparer,
			},
			want: NewEnSlice("A", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectCmpMust(tt.args.first, tt.args.second, tt.args.cmp)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectCmpMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}
