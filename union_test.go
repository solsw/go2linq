//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/UnionTest.cs

func Test_Union_string(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "UnionWithTwoEmptySequences",
			args: args{
				first: Empty[string](),
				second: Empty[string](),
			},
			want: Empty[string](),
		},
		{name: "FirstEmpty",
			args: args{
				first: Empty[string](),
				second: NewOnSlice[string]("one", "two", "three", "four"),
			},
			want: NewOnSlice[string]("one", "two", "three", "four"),
		},
		{name: "SecondEmpty",
			args: args{
				first: NewOnSlice[string]("one", "two", "three", "four"),
				second: Empty[string](),
			},
			want: NewOnSlice[string]("one", "two", "three", "four"),
		},
		{name: "UnionWithoutComparer",
			args: args{
				first: NewOnSlice("a", "b", "B", "c", "b"),
				second: NewOnSlice("d", "e", "d", "a"),
			},
			want: NewOnSlice("a", "b", "B", "c", "d", "e"),
		},
		{name: "UnionWithoutComparer2",
			args: args{
				first: NewOnSlice("a", "b"),
				second: NewOnSlice("b", "a"),
			},
			want: NewOnSlice("a", "b"),
		},
		{name: "UnionWithEmptyFirstSequence",
			args: args{
				first: Empty[string](),
				second: NewOnSlice("d", "e", "d", "a"),
			},
			want: NewOnSlice("d", "e", "a"),
		},
		{name: "UnionWithEmptySecondSequence",
			args: args{
				first: NewOnSlice("a", "b", "B", "c", "b"),
				second: Empty[string](),
			},
			want: NewOnSlice("a", "b", "B", "c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Union() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionSelf_int(t *testing.T) {
	e1 := NewOnSlice(1, 2, 3, 4)
	e2 := NewOnSlice(1, 2, 3, 4)
	e3 := NewOnSlice(1, 2, 3, 4)
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SameEnumerable1",
			args: args{
				first: e1,
				second: e1,
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first: Take(e2, 1),
				second: Skip(e2, 3),
			},
			want: NewOnSlice(1, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first: Skip(e3, 2),
				second: e3,
			},
			want: NewOnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionSelf(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionEq_int(t *testing.T) {
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
		eq Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "UnionWithIntEquality",
			args: args{
				first: NewOnSlice(1, 2),
				second: NewOnSlice(2, 3),
				eq: IntEqualer,
			},
			want: NewOnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionEq_string(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
		eq Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "UnionWithCaseInsensitiveComparerEq",
			args: args{
				first: NewOnSlice("a", "b", "B", "c", "b"),
				second: NewOnSlice("d", "e", "d", "a"),
				eq: CaseInsensitiveEqualer,
			},
			want: NewOnSlice("a", "b", "c", "d", "e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionCmp_int(t *testing.T) {
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
		cmp Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "UnionWithIntComparer1",
			args: args{
				first: NewOnSlice(1, 2, 2),
				second: Empty[int](),
				cmp: IntComparer,
			},
			want: NewOnSlice(1, 2),
		},
		{name: "UnionWithIntComparer2",
			args: args{
				first: NewOnSlice(1, 2),
				second: NewOnSlice(2, 3),
				cmp: IntComparer,
			},
			want: NewOnSlice(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionCmp_string(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
		cmp Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "UnionWithCaseInsensitiveComparerCmp",
			args: args{
				first: NewOnSlice("a", "b", "B", "c", "b"),
				second: NewOnSlice("d", "e", "d", "a"),
				cmp: CaseInsensitiveComparer,
			},
			want: NewOnSlice("a", "b", "c", "d", "e")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqual(got, tt.want) {
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
		first Enumerator[int]
		second Enumerator[int]
		cmp Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SameEnumerable1",
			args: args{
				first: e1,
				second: e1,
				cmp: IntComparer,
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first: Skip(e2, 2),
				second: Take(e2, 1),
				cmp: IntComparer,
			},
			want: NewOnSlice(3, 4, 1),
		},
		{name: "SameEnumerable3",
			args: args{
				first: Skip(e3, 2),
				second: e3,
				cmp: IntComparer,
			},
			want: NewOnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionCmpSelf(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionCmpSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
