//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ExceptTest.cs

func Test_ExceptMust_int(t *testing.T) {
	i4 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "IntWithoutComparer",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "IdenticalEnumerable",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4),
				second: NewEnSlice(1, 2, 3, 4),
			},
			want: Empty[int](),
		},
		{name: "IdenticalEnumerable2",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4),
				second: SkipMust(NewEnSlice(1, 2, 3, 4), 2),
			},
			want: NewEnSlice(1, 2),
		},
		{name: "SameEnumerable",
			args: args{
				first:  i4,
				second: SkipMust(i4, 2),
			},
			want: NewEnSlice(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_ExceptMust_string(t *testing.T) {
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
				first:  NewEnSlice("A", "a", "b", "c", "b", "c"),
				second: NewEnSlice("b", "a", "d", "a"),
			},
			want: NewEnSlice("A", "c"),
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#except-and-exceptby
		{name: "Except",
			args: args{
				first:  NewEnSlice("Mercury", "Venus", "Earth", "Jupiter"),
				second: NewEnSlice("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: NewEnSlice("Venus"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_ExceptEqMust_int(t *testing.T) {
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
				equaler: Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptEqMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_ExceptEqMust_string(t *testing.T) {
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
			want: NewEnSlice("c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptEqMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_ExceptCmpMust_int(t *testing.T) {
	i4 := NewEnSlice(1, 2, 3, 4)
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
		{name: "IntComparerSpecified",
			args: args{
				first:    NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second:   NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				comparer: Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "SameEnumerable",
			args: args{
				first:    i4,
				second:   SkipMust(i4, 2),
				comparer: Order[int]{},
			},
			want: NewEnSlice(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptCmpMust(tt.args.first, tt.args.second, tt.args.comparer)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptCmpMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_ExceptCmpMust_string(t *testing.T) {
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
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:    NewEnSlice("A", "a", "b", "c", "b"),
				second:   NewEnSlice("b", "a", "d", "a"),
				comparer: CaseInsensitiveComparer,
			},
			want: NewEnSlice("c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptCmpMust(tt.args.first, tt.args.second, tt.args.comparer)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptCmpMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}
