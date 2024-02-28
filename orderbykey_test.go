package go2linq

import (
	"cmp"
	"fmt"
	"iter"
	"math"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByDescendingTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByDescendingTest.cs

func TestOrderByKey_string_int(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[string]
		wantErr bool
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#primary-ascending-sort
		{name: "Primary Ascending Sort",
			args: args{
				source:      VarAll("the", "quick", "brown", "fox", "jumps", "over"),
				keySelector: func(s string) int { return len(s) },
			},
			want: VarAll("the", "fox", "over", "quick", "brown", "jumps"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderByKey(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OrderByKey() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestOrderByKeyLs_intint(t *testing.T) {
	type args struct {
		source      iter.Seq[elel[int]]
		keySelector func(elel[int]) int
		less        func(x, y int) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      VarAll(elel[int]{1, 10}, elel[int]{2, 12}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				less:        cmp.Less[int],
			},
			want: VarAll(1, 3, 2),
		},
		{name: "OrderingIsStable",
			args: args{
				source:      VarAll(elel[int]{1, 10}, elel[int]{2, 11}, elel[int]{3, 11}, elel[int]{4, 10}),
				keySelector: func(e elel[int]) int { return e.e2 },
				less:        cmp.Less[int],
			},
			want: VarAll(1, 4, 2, 3),
		},
		{name: "CustomLess",
			args: args{
				source:      VarAll(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				less: func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				},
			},
			want: VarAll(3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Select[elel[int], int](
				errorhelper.Must(OrderByKeyLs(tt.args.source, tt.args.keySelector, tt.args.less)),
				func(e elel[int]) int { return e.e1 },
			)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OrderByKeyLs() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestOrderByKeyDesc_string_rune(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) rune
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#primary-descending-sort
		{name: "Primary Descending Sort",
			args: args{
				source:      VarAll("the", "quick", "brown", "fox", "jumps", "over"),
				keySelector: func(s string) rune { return []rune(s)[0] },
			},
			want: VarAll("the", "quick", "over", "jumps", "fox", "brown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := OrderByKeyDesc(tt.args.source, tt.args.keySelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OrderByKeyDesc() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestOrderByKeyDescLs_intint(t *testing.T) {
	type args struct {
		source      iter.Seq[elel[int]]
		keySelector func(elel[int]) int
		less        func(x, y int) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      VarAll(elel[int]{1, 10}, elel[int]{2, 12}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				less:        cmp.Less[int],
			},
			want: VarAll(2, 3, 1),
		},
		{name: "OrderingIsStable",
			args: args{
				source:      VarAll(elel[int]{1, 10}, elel[int]{2, 11}, elel[int]{3, 11}, elel[int]{4, 10}),
				keySelector: func(e elel[int]) int { return e.e2 },
				less:        cmp.Less[int],
			},
			want: VarAll(2, 3, 1, 4),
		},
		{name: "CustomLess",
			args: args{
				source:      VarAll(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				less: func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				},
			},
			want: VarAll(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Select[elel[int], int](
				errorhelper.Must(OrderByKeyDescLs(tt.args.source, tt.args.keySelector, tt.args.less)),
				func(e elel[int]) int { return e.e1 },
			)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OrderByKeyDescLs() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func ExampleOrderByKeyDesc() {
	fmt.Println(StringDef[string](
		errorhelper.Must(OrderByKeyDesc(
			VarAll("zero", "one", "two", "three", "four", "five"),
			func(s string) int { return len(s) },
		)),
	))
	// Output:
	// [three zero four five one two]
}
