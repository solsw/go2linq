//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/UnionTest.cs

func Test_UnionMust_string(t *testing.T) {
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "UnionWithTwoEmptySequences",
			args: args{
				first:  Empty[string](),
				second: Empty[string](),
			},
			want: Empty[string](),
		},
		{name: "FirstEmpty",
			args: args{
				first:  Empty[string](),
				second: NewEnSlice("one", "two", "three", "four"),
			},
			want: NewEnSlice("one", "two", "three", "four"),
		},
		{name: "SecondEmpty",
			args: args{
				first:  NewEnSlice("one", "two", "three", "four"),
				second: Empty[string](),
			},
			want: NewEnSlice("one", "two", "three", "four"),
		},
		{name: "UnionWithoutComparer",
			args: args{
				first:  NewEnSlice("a", "b", "B", "c", "b"),
				second: NewEnSlice("d", "e", "d", "a"),
			},
			want: NewEnSlice("a", "b", "B", "c", "d", "e"),
		},
		{name: "UnionWithoutComparer2",
			args: args{
				first:  NewEnSlice("a", "b"),
				second: NewEnSlice("b", "a"),
			},
			want: NewEnSlice("a", "b"),
		},
		{name: "UnionWithEmptyFirstSequence",
			args: args{
				first:  Empty[string](),
				second: NewEnSlice("d", "e", "d", "a"),
			},
			want: NewEnSlice("d", "e", "a"),
		},
		{name: "UnionWithEmptySecondSequence",
			args: args{
				first:  NewEnSlice("a", "b", "B", "c", "b"),
				second: Empty[string](),
			},
			want: NewEnSlice("a", "b", "B", "c"),
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#union-and-unionby
		{name: "Union",
			args: args{
				first:  NewEnSlice("Mercury", "Venus", "Earth", "Jupiter"),
				second: NewEnSlice("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: NewEnSlice("Mercury", "Venus", "Earth", "Jupiter", "Mars"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_UnionMust_int(t *testing.T) {
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
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  TakeMust(e2, 1),
				second: SkipMust(e2, 3),
			},
			want: NewEnSlice(1, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 2),
				second: e3,
			},
			want: NewEnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_UnionEqMust_int(t *testing.T) {
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
		{name: "UnionWithIntEquality",
			args: args{
				first:   NewEnSlice(1, 2),
				second:  NewEnSlice(2, 3),
				equaler: Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionEqMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_UnionEqMust_string(t *testing.T) {
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
		{name: "UnionWithCaseInsensitiveComparerEq",
			args: args{
				first:   NewEnSlice("a", "b", "B", "c", "b"),
				second:  NewEnSlice("d", "e", "d", "a"),
				equaler: CaseInsensitiveEqualer,
			},
			want: NewEnSlice("a", "b", "c", "d", "e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionEqMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_UnionCmpMust_int(t *testing.T) {
	e1 := NewEnSlice(1, 2, 3, 4)
	e2 := NewEnSlice(1, 2, 3, 4)
	e3 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first    Enumerable[int]
		second   Enumerable[int]
		comparer Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "UnionWithIntComparer1",
			args: args{
				first:    NewEnSlice(1, 2, 2),
				second:   Empty[int](),
				comparer: Order[int]{},
			},
			want: NewEnSlice(1, 2),
		},
		{name: "UnionWithIntComparer2",
			args: args{
				first:    NewEnSlice(1, 2),
				second:   NewEnSlice(2, 3),
				comparer: Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "SameEnumerable1",
			args: args{
				first:    e1,
				second:   e1,
				comparer: Order[int]{},
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:    SkipMust[int](e2, 2),
				second:   TakeMust[int](e2, 1),
				comparer: Order[int]{},
			},
			want: NewEnSlice(3, 4, 1),
		},
		{name: "SameEnumerable3",
			args: args{
				first:    SkipMust[int](e3, 2),
				second:   e3,
				comparer: Order[int]{},
			},
			want: NewEnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionCmpMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_UnionCmpMust_string(t *testing.T) {
	type args struct {
		first    Enumerable[string]
		second   Enumerable[string]
		comparer Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "UnionWithCaseInsensitiveComparerCmp",
			args: args{
				first:    NewEnSlice("a", "b", "B", "c", "b"),
				second:   NewEnSlice("d", "e", "d", "a"),
				comparer: CaseInsensitiveComparer,
			},
			want: NewEnSlice("a", "b", "c", "d", "e")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionCmpMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}
