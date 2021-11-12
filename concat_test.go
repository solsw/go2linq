//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ConcatTest.cs

func Test_Concat_int(t *testing.T) {
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "Empty",
			args: args{
				first: Empty[int](),
				second: Empty[int](),
			},
			want: Empty[int](),
		},
		{name: "SemiEmpty", 
			args: args{
				first: NewOnSlice(1, 2, 3, 4),
				second: Empty[int](),
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "SimpleConcatenation",
			args: args{
				first: NewOnSlice(1, 2, 3, 4),
				second: NewOnSlice(1, 2, 3, 4),
			},
			want: NewOnSlice(1, 2, 3, 4, 1, 2, 3, 4),
		},
		{name: "SimpleConcatenation2",
			args: args{
				first: Range(1, 2),
				second: Repeat(3, 1),
			},
			want: NewOnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Concat() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Concat_int2(t *testing.T) {
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SecondSequenceIsntAccessedBeforeFirstUse",
			args: args{
				first: NewOnSlice(1, 2, 3, 4),
				second: Select(NewOnSlice(0, 1), func(x int) int { return 2 / x }),
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "NotNeededElementsAreNotAccessed",
			args: args{
				first: NewOnSlice(1, 2, 3),
				second: Select(NewOnSlice(1, 0), func(x int) int { return 2 / x }),
			},
			want: NewOnSlice(1, 2, 3, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Take(Concat(tt.args.first, tt.args.second), 4); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Concat() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Concat_string(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "Empty",
			args: args{
				first: Empty[string](),
				second: Empty[string](),
			},
			want: Empty[string](),
		},
		{name: "SemiEmpty",
			args: args{
				first: Empty[string](),
				second: NewOnSlice("one", "two", "three", "four"),
			},
			want: NewOnSlice("one", "two", "three", "four"),
		},
		{name: "SimpleConcatenation",
			args: args{
				first: NewOnSlice("a", "b"),
				second: NewOnSlice("c", "d"),
			},
			want: NewOnSlice("a", "b", "c", "d"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Concat() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ConcatSelf_int(t *testing.T) {
	i4 := NewOnSlice(1, 2, 3, 4)
	rg := Range(1, 4)
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
				second: i4,
			},
			want: NewOnSlice(1, 2, 3, 4, 1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first: Take(rg, 2),
				second: Skip(rg, 2),
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatSelf(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ConcatSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ConcatSelf_string(t *testing.T) {
	rs := Skip(Repeat("q", 2), 1)
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "SameEnumerable",
			args: args{
				first: rs,
				second: rs,
			},
			want: NewOnSlice("q", "q"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatSelf(tt.args.first, tt.args.second); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ConcatSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
