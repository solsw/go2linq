package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeWhileTest.cs

func TestTake_int(t *testing.T) {
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
			want: Empty[int](),
		},
		{name: "ZeroCount",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  0,
			},
			want: Empty[int](),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  3,
			},
			want: VarAll(0, 1, 2),
		},
		{name: "CountShorterThanSource2",
			args: args{
				source: VarAll(1, 2, 3, 4),
				count:  3,
			},
			want: VarAll(1, 2, 3),
		},
		{name: "CountEqualToSourceLength",
			args: args{
				source: errorhelper.Must(Range(1, 5)),
				count:  5,
			},
			want: VarAll(1, 2, 3, 4, 5),
		},
		{name: "CountGreaterThanSourceLength",
			args: args{
				source: errorhelper.Must(Range(2, 5)),
				count:  100,
			},
			want: VarAll(2, 3, 4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Take(tt.args.source, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Take() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Take() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestTakeLast_int(t *testing.T) {
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
			want: Empty[int](),
		},
		{name: "ZeroCount",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  0,
			},
			want: Empty[int](),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: errorhelper.Must(Range(0, 5)),
				count:  3,
			},
			want: VarAll(2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TakeLast(tt.args.source, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("TakeLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("TakeLast() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestTakeWhile_string(t *testing.T) {
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
				source:    VarAll("zero", "one", "two", "three", "four", "five", "six"),
				predicate: func(s string) bool { return len(s) > 4 },
			},
			want: Empty[string](),
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    VarAll("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 5 },
			},
			want: VarAll("zero", "one", "two"),
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    VarAll("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 100 },
			},
			want: VarAll("zero", "one", "two", "three", "four", "five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TakeWhile(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("TakeWhile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("TakeWhile() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestTakeWhileIdx_string(t *testing.T) {
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
				source:    VarAll("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return idx+len(s) > 4 },
			},
			want: Empty[string](),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source:    VarAll("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return len(s) != idx },
			},
			want: VarAll("zero", "one", "two", "three"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source:    VarAll("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, _ int) bool { return len(s) < 100 },
			},
			want: VarAll("zero", "one", "two", "three", "four", "five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TakeWhileIdx(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("TakeWhileIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("TakeWhileIdx() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func ExampleTakeWhileIdx() {
	fruits := []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry"}
	takeWhileIdx, _ := TakeWhileIdx(
		SliceAll(fruits),
		func(fruit string, index int) bool {
			return len(fruit) >= index
		},
	)
	for fruit := range takeWhileIdx {
		fmt.Println(fruit)
	}
	// Output:
	// apple
	// passionfruit
	// banana
	// mango
	// orange
	// blueberry
}

// second example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func ExampleTakeWhile() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	takeWhile, _ := TakeWhile(
		SliceAll(fruits),
		func(fruit string) bool { return fruit != "orange" },
	)
	for fruit := range takeWhile {
		fmt.Println(fruit)
	}
	// Output:
	// apple
	// banana
	// mango
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.take
func ExampleTake() {
	grades := []int{59, 82, 70, 56, 92, 98, 85}
	orderedGrades, _ := OrderByDesc(SliceAll(grades))
	topThreeGrades, _ := Take[int](orderedGrades, 3)
	fmt.Println("The top three grades are:")
	for grade := range topThreeGrades {
		fmt.Println(grade)
	}
	// Output:
	// The top three grades are:
	// 98
	// 92
	// 85
}
