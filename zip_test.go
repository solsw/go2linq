package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ZipTest.cs

func TestZip_string_int_string(t *testing.T) {
	type args struct {
		first          iter.Seq[string]
		second         iter.Seq[int]
		resultSelector func(string, int) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "ShortFirst",
			args: args{
				first:          VarAll("a", "b", "c"),
				second:         errorhelper.Must(Range(5, 10)),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: VarAll("a:5", "b:6", "c:7"),
		},
		{name: "ShortSecond",
			args: args{
				first:          VarAll("a", "b", "c", "d", "e"),
				second:         errorhelper.Must(Range(5, 3)),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: VarAll("a:5", "b:6", "c:7"),
		},
		{name: "EqualLengthSequences",
			args: args{
				first:          VarAll("a", "b", "c"),
				second:         errorhelper.Must(Range(5, 3)),
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: VarAll("a:5", "b:6", "c:7"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Zip() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestZip_string_string_string(t *testing.T) {
	en1 := VarAll("a", "b", "c")
	ee := VarAll("a", "b", "c", "d", "e")
	type args struct {
		first          iter.Seq[string]
		second         iter.Seq[string]
		resultSelector func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "1",
			args: args{
				first:          VarAll("one", "two", "three", "four"),
				second:         errorhelper.Must(Reverse(VarAll("one", "two", "three", "four"))),
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: VarAll("onefour", "twothree", "threetwo", "fourone"),
		},
		{name: "SameEnumerableString1",
			args: args{
				first:          en1,
				second:         en1,
				resultSelector: func(s1, s2 string) string { return fmt.Sprintf("%s:%s", s1, s2) },
			},
			want: VarAll("a:a", "b:b", "c:c"),
		},
		{name: "AdjacentElements",
			args: args{
				first:          ee,
				second:         errorhelper.Must(Skip(ee, 1)),
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: VarAll("ab", "bc", "cd", "de"),
		},
		{name: "AdjacentElements2",
			args: args{
				first:          errorhelper.Must(Skip(ee, 1)),
				second:         ee,
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: VarAll("ba", "cb", "dc", "ed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Zip() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestZip_int_int_string(t *testing.T) {
	en0, _ := Range(1, 4)
	en1, _ := Take(errorhelper.Must(Range(1, 4)), 2)
	en2, _ := TakeLast(errorhelper.Must(Range(1, 4)), 2)
	type args struct {
		first          iter.Seq[int]
		second         iter.Seq[int]
		resultSelector func(int, int) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "SameEnumerableInt00",
			args: args{
				first:          en0,
				second:         en0,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: VarAll("1:1", "2:2", "3:3", "4:4"),
		},
		{name: "SameEnumerableInt01",
			args: args{
				first:          errorhelper.Must(Skip(en0, 2)),
				second:         en0,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: VarAll("3:1", "4:2"),
		},
		{name: "SameEnumerableInt1",
			args: args{
				first:          en1,
				second:         en1,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: VarAll("1:1", "2:2"),
		},
		{name: "SameEnumerableInt2",
			args: args{
				first:          en2,
				second:         en2,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d:%d", i1, i2) },
			},
			want: VarAll("3:3", "4:4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Zip() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestZip_string_string_int(t *testing.T) {
	type args struct {
		first          iter.Seq[string]
		second         iter.Seq[string]
		resultSelector func(string, string) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "1",
			args: args{
				first:          VarAll("a", "b", "c"),
				second:         VarAll("one", "two", "three", "four"),
				resultSelector: func(s1, s2 string) int { return len(s1 + s2) },
			},
			want: VarAll(4, 4, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Zip() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestZip_int_rune_string(t *testing.T) {
	type args struct {
		first          iter.Seq[int]
		second         iter.Seq[rune]
		resultSelector func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#zip
		{name: "Zip",
			args: args{
				first:          VarAll(1, 2, 3, 4, 5, 6, 7),
				second:         VarAll('A', 'B', 'C', 'D', 'E', 'F'),
				resultSelector: func(number int, letter rune) string { return fmt.Sprintf("%d = %c (%[2]d)", number, letter) },
			},
			want: VarAll("1 = A (65)", "2 = B (66)", "3 = C (67)", "4 = D (68)", "5 = E (69)", "6 = F (70)"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Zip() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.zip
func ExampleZip() {
	numbers := []int{1, 2, 3, 4}
	words := []string{"one", "two", "three"}
	zip, _ := Zip(SliceAll(numbers), SliceAll(words),
		func(first int, second string) string { return fmt.Sprintf("%d %s", first, second) },
	)
	for item := range zip {
		fmt.Println(item)
	}
	// Output:
	// 1 one
	// 2 two
	// 3 three
}
