//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/UnionTest.cs

func Test_Union_string(t *testing.T) {
	type args struct {
		first  Enumerator[string]
		second Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
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
				second: NewOnSlice("one", "two", "three", "four"),
			},
			want: NewOnSlice("one", "two", "three", "four"),
		},
		{name: "SecondEmpty",
			args: args{
				first:  NewOnSlice("one", "two", "three", "four"),
				second: Empty[string](),
			},
			want: NewOnSlice("one", "two", "three", "four"),
		},
		{name: "UnionWithoutComparer",
			args: args{
				first:  NewOnSlice("a", "b", "B", "c", "b"),
				second: NewOnSlice("d", "e", "d", "a"),
			},
			want: NewOnSlice("a", "b", "B", "c", "d", "e"),
		},
		{name: "UnionWithoutComparer2",
			args: args{
				first:  NewOnSlice("a", "b"),
				second: NewOnSlice("b", "a"),
			},
			want: NewOnSlice("a", "b"),
		},
		{name: "UnionWithEmptyFirstSequence",
			args: args{
				first:  Empty[string](),
				second: NewOnSlice("d", "e", "d", "a"),
			},
			want: NewOnSlice("d", "e", "a"),
		},
		{name: "UnionWithEmptySecondSequence",
			args: args{
				first:  NewOnSlice("a", "b", "B", "c", "b"),
				second: Empty[string](),
			},
			want: NewOnSlice("a", "b", "B", "c"),
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#union-and-unionby
		{name: "Union",
			args: args{
				first:  NewOnSlice("Mercury", "Venus", "Earth", "Jupiter"),
				second: NewOnSlice("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: NewOnSlice("Mercury", "Venus", "Earth", "Jupiter", "Mars"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Union(tt.args.first, tt.args.second); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Union() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionSelf_int(t *testing.T) {
	e1 := NewOnSlice(1, 2, 3, 4)
	e2 := NewOnSliceEn(1, 2, 3, 4)
	e3 := NewOnSliceEn(1, 2, 3, 4)
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  TakeMust(e2, 1),
				second: SkipMust(e2, 3),
			},
			want: NewOnSlice(1, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 2),
				second: e3,
			},
			want: NewOnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionSelf(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionEq_int(t *testing.T) {
	type args struct {
		first   Enumerator[int]
		second  Enumerator[int]
		equaler Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "UnionWithIntEquality",
			args: args{
				first:   NewOnSlice(1, 2),
				second:  NewOnSlice(2, 3),
				equaler: Order[int]{},
			},
			want: NewOnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionEq(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionEq_string(t *testing.T) {
	type args struct {
		first   Enumerator[string]
		second  Enumerator[string]
		equaler Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "UnionWithCaseInsensitiveComparerEq",
			args: args{
				first:   NewOnSlice("a", "b", "B", "c", "b"),
				second:  NewOnSlice("d", "e", "d", "a"),
				equaler: CaseInsensitiveEqualer,
			},
			want: NewOnSlice("a", "b", "c", "d", "e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionEq(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionCmp_int(t *testing.T) {
	type args struct {
		first    Enumerator[int]
		second   Enumerator[int]
		comparer Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "UnionWithIntComparer1",
			args: args{
				first:    NewOnSlice(1, 2, 2),
				second:   Empty[int](),
				comparer: Order[int]{},
			},
			want: NewOnSlice(1, 2),
		},
		{name: "UnionWithIntComparer2",
			args: args{
				first:    NewOnSlice(1, 2),
				second:   NewOnSlice(2, 3),
				comparer: Order[int]{},
			},
			want: NewOnSlice(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := UnionCmp(tt.args.first, tt.args.second, tt.args.comparer); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionCmp_string(t *testing.T) {
	type args struct {
		first    Enumerator[string]
		second   Enumerator[string]
		comparer Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "UnionWithCaseInsensitiveComparerCmp",
			args: args{
				first:    NewOnSlice("a", "b", "B", "c", "b"),
				second:   NewOnSlice("d", "e", "d", "a"),
				comparer: CaseInsensitiveComparer,
			},
			want: NewOnSlice("a", "b", "c", "d", "e")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := UnionCmp(tt.args.first, tt.args.second, tt.args.comparer); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionCmpSelf_int(t *testing.T) {
	e1 := NewOnSlice(1, 2, 3, 4)
	e2 := NewOnSlice(1, 2, 3, 4)
	e3 := NewOnSlice(1, 2, 3, 4)
	type args struct {
		first    Enumerator[int]
		second   Enumerator[int]
		comparer Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SameEnumerable1",
			args: args{
				first:    e1,
				second:   e1,
				comparer: Order[int]{},
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:    SkipMust[int](e2, 2),
				second:   TakeMust[int](e2, 1),
				comparer: Order[int]{},
			},
			want: NewOnSlice(3, 4, 1),
		},
		{name: "SameEnumerable3",
			args: args{
				first:    SkipMust[int](e3, 2),
				second:   e3,
				comparer: Order[int]{},
			},
			want: NewOnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := UnionCmpSelf(tt.args.first, tt.args.second, tt.args.comparer); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionCmpSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
