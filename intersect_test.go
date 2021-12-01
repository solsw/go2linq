//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/IntersectTest.cs

func Test_Intersect_int(t *testing.T) {
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "1",
			args: args{
				first:  NewOnSlice(1, 2),
				second: NewOnSlice(2, 3),
			},
			want: NewOnSlice(2),
		},
		{name: "IntWithoutComparer",
			args: args{
				first:  NewOnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewOnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: NewOnSlice(4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Intersect(tt.args.first, tt.args.second); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Intersect() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Intersect_string(t *testing.T) {
	type args struct {
		first  Enumerator[string]
		second Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "NoComparerSpecified",
			args: args{
				first:  NewOnSlice("A", "a", "b", "c", "b"),
				second: NewOnSlice("b", "a", "d", "a"),
			},
			want: NewOnSlice("a", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Intersect(tt.args.first, tt.args.second); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Intersect() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_IntersectSelf_int(t *testing.T) {
	e1 := NewOnSlice(1, 2, 3, 4)
	e2 := NewOnSlice(1, 2, 3, 4)
	e3 := NewOnSlice(1, 2, 3, 4)
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
				first:  e2,
				second: SkipMust(e2, 1),
			},
			want: NewOnSlice(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 3),
				second: e3,
			},
			want: NewOnSlice(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IntersectSelf(tt.args.first, tt.args.second); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("IntersectSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_IntersectEq_int(t *testing.T) {
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
		eq     Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:  NewOnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewOnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				eq:     IntEqualer},
			want: NewOnSlice(4, 5, 6, 7, 8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IntersectEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("IntersectEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_IntersectEq_string(t *testing.T) {
	type args struct {
		first  Enumerator[string]
		second Enumerator[string]
		eq     Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:  NewOnSlice("A", "a", "b", "c", "b"),
				second: NewOnSlice("b", "a", "d", "a"),
				eq:     CaseInsensitiveEqualer,
			},
			want: NewOnSlice("A", "b")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IntersectEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("IntersectEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_IntersectCmp_int(t *testing.T) {
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
		cmp    Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:  NewOnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewOnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				cmp:    IntComparer,
			},
			want: NewOnSlice(4, 5, 6, 7, 8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IntersectCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("IntersectCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_IntersectCmp_string(t *testing.T) {
	type args struct {
		first  Enumerator[string]
		second Enumerator[string]
		cmp    Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:  NewOnSlice("A", "a", "b", "c", "b"),
				second: NewOnSlice("b", "a", "d", "a"),
				cmp:    CaseInsensitiveComparer,
			},
			want: NewOnSlice("A", "b")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IntersectCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("IntersectCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_IntersectCmpSelf(t *testing.T) {
	e1 := NewOnSlice(4, 3, 2, 1)
	e2 := NewOnSlice(1, 2, 3, 4)
	e3 := NewOnSlice(1, 2, 3, 4)
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
		cmp    Comparer[int]
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
				cmp:    IntComparer,
			},
			want: NewOnSlice(4, 3, 2, 1),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  e2,
				second: SkipMust(e2, 1),
				cmp:    IntComparer,
			},
			want: NewOnSlice(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 3),
				second: e3,
				cmp:    IntComparer,
			},
			want: NewOnSlice(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IntersectCmpSelf(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("IntersectCmpSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
