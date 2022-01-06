//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ExceptTest.cs

func Test_Except_int(t *testing.T) {
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "IntWithoutComparer",
			args: args{
				first:  NewOnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewOnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: NewOnSlice(1, 2, 3),
		},
		{name: "IdenticalEnumerable",
			args: args{
				first:  NewOnSlice(1, 2, 3, 4),
				second: NewOnSlice(1, 2, 3, 4),
			},
			want: Empty[int](),
		},
		{name: "IdenticalEnumerable2",
			args: args{
				first:  NewOnSlice(1, 2, 3, 4),
				second: SkipMust(NewOnSliceEn(1, 2, 3, 4), 2),
			},
			want: NewOnSlice(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Except(tt.args.first, tt.args.second); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Except() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Except_string(t *testing.T) {
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
				first:  NewOnSlice("A", "a", "b", "c", "b", "c"),
				second: NewOnSlice("b", "a", "d", "a"),
			},
			want: NewOnSlice("A", "c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Except(tt.args.first, tt.args.second); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Except() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptSelf_int(t *testing.T) {
	i4 := NewOnSliceEn(1, 2, 3, 4)
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SameEnumerable",
			args: args{
				first:  i4,
				second: SkipMust(i4, 2),
			},
			want: NewOnSlice(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ExceptSelf(tt.args.first, tt.args.second); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptEq_int(t *testing.T) {
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
				eq:     Orderer[int]{},
			},
			want: NewOnSlice(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ExceptEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptEq_string(t *testing.T) {
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
			want: NewOnSlice("c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ExceptEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptCmp_int(t *testing.T) {
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
				cmp:    Orderer[int]{},
			},
			want: NewOnSlice(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ExceptCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptCmp_string(t *testing.T) {
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
			want: NewOnSlice("c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ExceptCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptCmpSelf_int(t *testing.T) {
	i4 := NewOnSliceEn(1, 2, 3, 4)
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
		{name: "SameEnumerable",
			args: args{
				first:  i4,
				second: SkipMust(i4, 2),
				cmp:    Orderer[int]{},
			},
			want: NewOnSlice(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ExceptCmpSelf(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptCmpSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
