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

func TestOrderBy_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "1234",
			args: args{
				source: VarAll(4, 1, 3, 2),
			},
			want: VarAll(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderBy(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OrderBy() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestOrderByLs_elelel(t *testing.T) {
	type args struct {
		source iter.Seq[elelel[int]]
		less   func(elelel[int], elelel[int]) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "1",
			args: args{
				source: VarAll(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return x.e2 < y.e2 },
					func(x, y elelel[int]) bool { return x.e3 < y.e3 },
				),
			},
			want: VarAll(3, 1, 2),
		},
		{name: "PrimaryOrderingTakesPrecedence",
			args: args{
				source: VarAll(elelel[int]{1, 10, 20}, elelel[int]{2, 12, 21}, elelel[int]{3, 11, 22}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return x.e2 < y.e2 },
					func(x, y elelel[int]) bool { return x.e3 < y.e3 },
				),
			},
			want: VarAll(1, 3, 2),
		},
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			args: args{
				source: VarAll(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return x.e2 < y.e2 },
					func(x, y elelel[int]) bool { return x.e3 < y.e3 },
				),
			},
			want: VarAll(3, 1, 2),
		},
		{name: "ThenByAfterOrderByDescending",
			args: args{
				source: VarAll(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return y.e2 < x.e2 },
					func(x, y elelel[int]) bool { return x.e3 < y.e3 },
				),
			},
			want: VarAll(2, 3, 1),
		},
		{name: "OrderingIsStable",
			args: args{
				source: VarAll(elelel[int]{1, 1, 10}, elelel[int]{2, 1, 11}, elelel[int]{3, 1, 11}, elelel[int]{4, 1, 10}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return x.e2 < y.e2 },
					func(x, y elelel[int]) bool { return x.e3 < y.e3 },
				),
			},
			want: VarAll(1, 4, 2, 3),
		},
		{name: "CustomLess",
			args: args{
				source: VarAll(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return x.e2 < y.e2 },
					func(x, y elelel[int]) bool { return cmp.Less(math.Abs(float64(x.e3)), math.Abs(float64(y.e3))) },
				),
			},
			want: VarAll(3, 2, 1),
		},
		{name: "CustomComparer",
			args: args{
				source: VarAll(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return x.e2 < y.e2 },
					func(x, y elelel[int]) bool { return cmp.Compare(math.Abs(float64(x.e3)), math.Abs(float64(y.e3))) < 0 },
				),
			},
			want: VarAll(3, 2, 1),
		},
		{name: "ThenByDescendingAfterOrderByDescending",
			args: args{
				source: VarAll(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return cmp.Less(y.e2, x.e2) },
					func(x, y elelel[int]) bool { return y.e3 < x.e3 },
				),
			},
			want: VarAll(2, 1, 3),
		},
		{name: "DescendingOrderingIsStable",
			args: args{
				source: VarAll(elelel[int]{1, 1, 10}, elelel[int]{2, 1, 11}, elelel[int]{3, 1, 11}, elelel[int]{4, 1, 10}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return cmp.Less(y.e2, x.e2) },
					func(x, y elelel[int]) bool { return y.e3 < x.e3 },
				),
			},
			want: VarAll(2, 3, 1, 4),
		},
		{name: "CustomDescendingLess",
			args: args{
				source: VarAll(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return cmp.Less(y.e2, x.e2) },
					func(x, y elelel[int]) bool { return cmp.Less(math.Abs(float64(y.e3)), math.Abs(float64(x.e3))) },
				),
			},
			want: VarAll(1, 2, 3),
		},
		{name: "CustomDescendingComparer",
			args: args{
				source: VarAll(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
				less: ThenLess(
					func(x, y elelel[int]) bool { return cmp.Less(y.e2, x.e2) },
					func(x, y elelel[int]) bool { return cmp.Compare(math.Abs(float64(y.e3)), math.Abs(float64(x.e3))) < 0 },
				),
			},
			want: VarAll(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, _ := OrderByLs(tt.args.source, tt.args.less)
			got2, _ := Select(got1, func(e elelel[int]) int { return e.e1 })
			equal, _ := SequenceEqual(got2, tt.want)
			if !equal {
				t.Errorf("OrderByLs() = %v, want %v", StringDef(got2), StringDef(tt.want))
			}
		})
	}
}

func ExampleOrderBy() {
	fmt.Println(StringDef[string](
		errorhelper.Must(OrderBy(VarAll("zero", "one", "two", "three", "four", "five"))),
	))
	// Output:
	// [five four one three two zero]
}

// OrderByEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderby
func ExampleOrderByLs() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	orderByLs, _ := OrderByLs(SliceAll(pets), func(p1, p2 Pet) bool { return p1.Age < p2.Age })
	for pet := range orderByLs {
		fmt.Printf("%s - %d\n", pet.Name, pet.Age)
	}
	// Output:
	// Whiskers - 1
	// Boots - 4
	// Barley - 8
}

// OrderByDescendingEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
func ExampleOrderByDescLs() {
	decimals := []float64{6.2, 8.3, 0.5, 1.3, 6.3, 9.7}
	less := func(f1, f2 float64) bool {
		_, fr1 := math.Modf(f1)
		_, fr2 := math.Modf(f2)
		if math.Abs(fr1-fr2) < 0.001 {
			return f1 < f2
		}
		return fr1 < fr2
	}
	orderByDescLs, _ := OrderByDescLs(SliceAll(decimals), less)
	for num := range orderByDescLs {
		fmt.Println(num)
	}
	// Output:
	// 9.7
	// 0.5
	// 8.3
	// 6.3
	// 1.3
	// 6.2
}

// example from https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenby
func ExampleThenLess_1() {
	// Sort the strings first by their length and then alphabetically.
	orderByLs, _ := OrderByLs(
		VarAll("grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry"),
		ThenLess(
			func(s1, s2 string) bool { return len(s1) < len(s2) },
			func(s1, s2 string) bool { return s1 < s2 },
		))
	for fruit := range orderByLs {
		fmt.Println(fruit)
	}
	// Output:
	// apple
	// grape
	// mango
	// banana
	// orange
	// blueberry
	// raspberry
	// passionfruit
}

// ThenByDescendingEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending
func ExampleThenLess_2() {
	// Sort the strings first ascending by their length and then descending using a custom case insensitive comparer.
	orderByLs, _ := OrderByLs(
		VarAll("apPLe", "baNanA", "apple", "APple", "orange", "BAnana", "ORANGE", "apPLE"),
		ThenLess(
			func(s1, s2 string) bool { return len(s1) < len(s2) },
			ReverseLess(caseInsensitiveLess),
		))
	for fruit := range orderByLs {
		fmt.Println(fruit)
	}
	// Output:
	// apPLe
	// apple
	// APple
	// apPLE
	// orange
	// ORANGE
	// baNanA
	// BAnana
}

func TestOrderByLs_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
		less   func(string, string) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#secondary-sort-examples
		{name: "Secondary Ascending Sort",
			args: args{
				source: VarAll("the", "quick", "brown", "fox", "jumps"),
				less: ThenLess(
					func(x, y string) bool { return len(x) < len(y) },
					func(x, y string) bool { return []rune(x)[0] < []rune(y)[0] },
				),
			},
			want: VarAll("fox", "the", "brown", "jumps", "quick"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#secondary-descending-sort
		{name: "Secondary Descending Sort",
			args: args{
				source: VarAll("the", "quick", "brown", "fox", "jumps"),
				less: ThenLess(
					func(x, y string) bool { return len(x) < len(y) },
					func(x, y string) bool { return []rune(y)[0] < []rune(x)[0] },
				),
			},
			want: VarAll("the", "fox", "quick", "jumps", "brown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := OrderByLs(tt.args.source, tt.args.less)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OrderByLs() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
