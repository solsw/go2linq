//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ExceptTest.cs

func Test_Except_int(t *testing.T) {
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "IntWithoutComparer",
			args: args{
				first: NewOnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewOnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: NewOnSlice(1, 2, 3),
		},
		{name: "IdenticalEnumerable",
			args: args{
				first: NewOnSlice(1, 2, 3, 4),
				second: NewOnSlice(1, 2, 3, 4),
			},
			want: Empty[int](),
		},
		{name: "IdenticalEnumerable2",
			args: args{
				first: NewOnSlice(1, 2, 3, 4),
				second: Skip(NewOnSlice(1, 2, 3, 4), 2),
			},
			want: NewOnSlice(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {			
			if got := Except(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Except() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Except_string(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "NoComparerSpecified",
			args: args{
				first: NewOnSlice("A", "a", "b", "c", "b", "c"),
				second: NewOnSlice("b", "a", "d", "a"),
			},
			want: NewOnSlice("A", "c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {			
			if got := Except(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Except() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptSelf_int(t *testing.T) {
	i4 := NewOnSlice(1, 2, 3, 4)
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SameEnumerable",
			args: args{
				first: i4,
				second: Skip(i4, 2),
			},
			want: NewOnSlice(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptSelf(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptEq_int(t *testing.T) {
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
		{name: "IntComparerSpecified",
			args: args{
				first: NewOnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewOnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				eq: IntEqualer,
			},
			want: NewOnSlice(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptEq_string(t *testing.T) {
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
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first: NewOnSlice("A", "a", "b", "c", "b"),
				second: NewOnSlice("b", "a", "d", "a"),
				eq: CaseInsensitiveEqualer,
			},
			want: NewOnSlice("c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptEq(tt.args.first, tt.args.second, tt.args.eq); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptCmp_int(t *testing.T) {
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
		{name: "IntComparerSpecified",
			args: args{
				first: NewOnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewOnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				cmp: IntComparer,
			},
			want: NewOnSlice(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptCmp_string(t *testing.T) {
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
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first: NewOnSlice("A", "a", "b", "c", "b"),
				second: NewOnSlice("b", "a", "d", "a"),
				cmp: CaseInsensitiveComparer,
			},
			want: NewOnSlice("c")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptCmp(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ExceptCmpSelf_int(t *testing.T) {
	i4 := NewOnSlice(1, 2, 3, 4)
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
		{name: "SameEnumerable",
			args: args{
				first: i4,
				second: Skip(i4, 2),
				cmp: IntComparer,
			},
			want: NewOnSlice(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptCmpSelf(tt.args.first, tt.args.second, tt.args.cmp); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ExceptCmpSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
