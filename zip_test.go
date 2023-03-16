package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ZipTest.cs

func TestZipMust_string_int_string(t *testing.T) {
	type args struct {
		first          Enumerable[string]
		second         Enumerable[int]
		resultSelector func(string, int) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "ShortFirst",
			args: args{
				first:          NewEnSlice("a", "b", "c"),
				second:         RangeMust(5, 10),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: NewEnSlice("a:5", "b:6", "c:7"),
		},
		{name: "ShortSecond",
			args: args{
				first:          NewEnSlice("a", "b", "c", "d", "e"),
				second:         RangeMust(5, 3),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: NewEnSlice("a:5", "b:6", "c:7"),
		},
		{name: "EqualLengthSequences",
			args: args{
				first:          NewEnSlice("a", "b", "c"),
				second:         RangeMust(5, 3),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: NewEnSlice("a:5", "b:6", "c:7"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZipMust(tt.args.first, tt.args.second, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ZipMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestZipMust_string_string_string(t *testing.T) {
	en1 := NewEnSlice("a", "b", "c")
	ee := NewEnSlice("a", "b", "c", "d", "e")
	type args struct {
		first          Enumerable[string]
		second         Enumerable[string]
		resultSelector func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "1",
			args: args{
				first:          NewEnSlice("one", "two", "three", "four"),
				second:         ReverseMust(NewEnSlice("one", "two", "three", "four")),
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: NewEnSlice("onefour", "twothree", "threetwo", "fourone"),
		},
		{name: "SameEnumerableString1",
			args: args{
				first:          en1,
				second:         en1,
				resultSelector: func(s1, s2 string) string { return fmt.Sprintf("%s:%s", s1, s2) },
			},
			want: NewEnSlice("a:a", "b:b", "c:c"),
		},
		{name: "AdjacentElements",
			args: args{
				first:          ee,
				second:         SkipMust(ee, 1),
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: NewEnSlice("ab", "bc", "cd", "de"),
		},
		{name: "AdjacentElements2",
			args: args{
				first:          SkipMust(ee, 1),
				second:         ee,
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: NewEnSlice("ba", "cb", "dc", "ed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZipMust(tt.args.first, tt.args.second, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ZipMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestZipMust_int_int_string(t *testing.T) {
	en0 := RangeMust(1, 4)
	en1 := TakeMust(RangeMust(1, 4), 2)
	en2 := TakeLastMust(RangeMust(1, 4), 2)
	type args struct {
		first          Enumerable[int]
		second         Enumerable[int]
		resultSelector func(int, int) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SameEnumerableInt00",
			args: args{
				first:          en0,
				second:         en0,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: NewEnSlice("1:1", "2:2", "3:3", "4:4"),
		},
		{name: "SameEnumerableInt01",
			args: args{
				first:          SkipMust(en0, 2),
				second:         en0,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: NewEnSlice("3:1", "4:2"),
		},
		{name: "SameEnumerableInt1",
			args: args{
				first:          en1,
				second:         en1,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: NewEnSlice("1:1", "2:2"),
		},
		{name: "SameEnumerableInt2",
			args: args{
				first:          en2,
				second:         en2,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: NewEnSlice("3:3", "4:4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZipMust(tt.args.first, tt.args.second, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ZipMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestZipMust_string_string_int(t *testing.T) {
	type args struct {
		first          Enumerable[string]
		second         Enumerable[string]
		resultSelector func(string, string) int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1",
			args: args{
				first:          NewEnSlice("a", "b", "c"),
				second:         NewEnSlice("one", "two", "three", "four"),
				resultSelector: func(s1, s2 string) int { return len(s1 + s2) },
			},
			want: NewEnSlice(4, 4, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZipMust(tt.args.first, tt.args.second, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ZipMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestZipMust_int_rune_string(t *testing.T) {
	type args struct {
		first          Enumerable[int]
		second         Enumerable[rune]
		resultSelector func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#zip
		{name: "Zip",
			args: args{
				first:          NewEnSlice(1, 2, 3, 4, 5, 6, 7),
				second:         NewEnSlice('A', 'B', 'C', 'D', 'E', 'F'),
				resultSelector: func(number int, letter rune) string { return fmt.Sprintf("%d = %c (%[2]d)", number, letter) },
			},
			want: NewEnSlice("1 = A (65)", "2 = B (66)", "3 = C (67)", "4 = D (68)", "5 = E (69)", "6 = F (70)"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZipMust(tt.args.first, tt.args.second, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ZipMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the example from Enumerable.Zip help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.zip
func ExampleZipMust() {
	numbers := NewEnSlice(1, 2, 3, 4)
	words := NewEnSlice("one", "two", "three")
	zip := ZipMust(numbers, words,
		func(first int, second string) string { return fmt.Sprintf("%d %s", first, second) },
	)
	enr := zip.GetEnumerator()
	for enr.MoveNext() {
		item := enr.Current()
		fmt.Println(item)
	}
	// Output:
	// 1 one
	// 2 two
	// 3 three
}
