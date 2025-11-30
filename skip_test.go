package go2linq

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipWhileTest.cs

func TestSkip_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
		count  int
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "NegativeCount",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  -5,
			},
			want: iterhelper.VarSeq(0, 1, 2, 3, 4),
		},
		{name: "ZeroCount",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  0,
			},
			want: iterhelper.VarSeq(0, 1, 2, 3, 4),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  3,
			},
			want: iterhelper.VarSeq(3, 4),
		},
		{name: "CountEqualToSourceLength",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  5,
			},
			want: Empty[int](),
		},
		{name: "CountGreaterThanSourceLength",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  100,
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Skip(tt.args.source, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Skip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Skip() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestSkipLast_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
		count  int
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "NegativeCount",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  -5,
			},
			want: errorhelper.Must(Range(0, 5)),
		},
		{name: "ZeroCount",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  0,
			},
			want: errorhelper.Must(Range(0, 5)),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  3,
			},
			want: iterhelper.VarSeq(0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SkipLast(tt.args.source, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("SkipLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SkipLast() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestSkipWhile_string(t *testing.T) {
	type args struct {
		source    iter.Seq[string]
		predicate func(string) bool
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[string]
		wantErr bool
	}{
		{name: "PredicateFailingFirstElement",
			args: args{
				source:    iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) > 4 },
			},
			want: iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 5 },
			},
			want: iterhelper.VarSeq("three", "four", "five"),
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 100 },
			},
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SkipWhile(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SkipWhile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SkipWhile() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestSkipWhileIdx_string(t *testing.T) {
	type args struct {
		source    iter.Seq[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[string]
		wantErr bool
	}{
		{name: "PredicateWithIndexFailingFirstElement",
			args: args{
				source:    iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return idx+len(s) > 4 },
			},
			want: iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source:    iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return len(s) > idx },
			},
			want: iterhelper.VarSeq("four", "five"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source:    iterhelper.VarSeq("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, _ int) bool { return len(s) < 100 },
			},
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SkipWhileIdx(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SkipWhileIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SkipWhileIdx() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.skip#examples
func ExampleSkip() {
	grades := []int{59, 82, 70, 56, 92, 98, 85}
	orderedGrades, _ := OrderByDesc(slices.Values(grades))
	lowerGrades, _ := Skip(orderedGrades, 3)
	fmt.Println("All grades except the top three are:")
	for grade := range lowerGrades {
		fmt.Println(grade)
	}
	// Output:
	// All grades except the top three are:
	// 82
	// 70
	// 59
	// 56
}

// second example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile
func ExampleSkipWhile() {
	grades := []int{59, 82, 70, 56, 92, 98, 85}
	orderedGrades, _ := OrderByDesc(slices.Values(grades))
	lowerGrades, _ := SkipWhile(orderedGrades, func(grade int) bool { return grade >= 80 })
	fmt.Println("All grades below 80:")
	for grade := range lowerGrades {
		fmt.Println(grade)
	}
	// Output:
	// All grades below 80:
	// 70
	// 59
	// 56
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile
func ExampleSkipWhileIdx() {
	amounts := []int{5000, 2500, 9000, 8000, 6500, 4000, 1500, 5500}
	skipWhileIdx, _ := SkipWhileIdx(slices.Values(amounts), func(amount, index int) bool { return amount > index*1000 })
	for amount := range skipWhileIdx {
		fmt.Println(amount)
	}
	// Output:
	// 4000
	// 1500
	// 5500
}
