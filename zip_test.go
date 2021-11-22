//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ZipTest.cs

func Test_Zip_string_int_string(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[int]
		resultSelector func(string, int) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "ShortFirst",
			args: args{
				first: NewOnSlice("a", "b", "c"),
				second: Range(5, 10),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: NewOnSlice("a:5", "b:6", "c:7"),
		},
		{name: "ShortSecond",
			args: args{
				first: NewOnSlice("a", "b", "c", "d", "e"),
				second: Range(5, 3),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: NewOnSlice("a:5", "b:6", "c:7"),
		},
		{name: "EqualLengthSequences",
			args: args{
				first: NewOnSlice("a", "b", "c"),
				second: Range(5, 3),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: NewOnSlice("a:5", "b:6", "c:7"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zip(tt.args.first, tt.args.second, tt.args.resultSelector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Zip() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ZipSelf_string(t *testing.T) {
	ee := NewOnSlice("a", "b", "c", "d", "e")
	r1 := Repeat("q", 2)
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
		resultSelector func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "AdjacentElements",
			args: args{
				first: ee,
				second: Skip(ee, 1),
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: NewOnSlice("ab", "bc", "cd", "de"),
		},
		{name: "SameEnumerable1",
			args: args{
				first: r1,
				second: r1,
				resultSelector: func(s1, s2 string) string { return s1 + ":" + s2 },
			},
			want: NewOnSlice("q:q", "q:q"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZipSelf(tt.args.first, tt.args.second, tt.args.resultSelector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ZipSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ZipSelf_int(t *testing.T) {
	r2 := Range(1, 4)
	type args struct {
		first Enumerator[int]
		second Enumerator[int]
		resultSelector func(int, int) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "SameEnumerable2",
			args: args{
				first: Skip(r2, 2),
				second: r2,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: NewOnSlice("3:1", "4:2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZipSelf(tt.args.first, tt.args.second, tt.args.resultSelector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ZipSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Zip_string_string_int(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
		resultSelector func(string, string) int
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "1",
			args: args{
				first: NewOnSlice("a", "b", "c"),
				second: NewOnSlice("one", "two", "three", "four"),
				resultSelector: func(s1, s2 string) int { return len(s1 + s2) },
			},
			want: NewOnSlice(4, 4, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zip(tt.args.first, tt.args.second, tt.args.resultSelector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Zip() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Zip_string(t *testing.T) {
	type args struct {
		first Enumerator[string]
		second Enumerator[string]
		resultSelector func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "1",
			args: args{
				first: NewOnSlice("one", "two", "three", "four"),
				second: Reverse(NewOnSlice("one", "two", "three", "four")),
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: NewOnSlice("onefour", "twothree", "threetwo", "fourone"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zip(tt.args.first, tt.args.second, tt.args.resultSelector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Zip() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
