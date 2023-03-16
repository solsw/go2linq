package go2linq

import (
	"fmt"
	"math"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByDescendingTest.cs

func TestThenByLsMust_1(t *testing.T) {
	type args struct {
		oe     *OrderedEnumerable[elelel[int]]
		lesser collate.Lesser[elelel[int]]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
				),
				lesser: collate.LesserFunc[elelel[int]](func(e1, e2 elelel[int]) bool { return e1.e3 < e2.e3 }),
			},
			want: NewEnSlice(3, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByLsMust(tt.args.oe, tt.args.lesser)
			got2 := SelectMust[elelel[int], int](got1, func(e elelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				t.Errorf("ThenByLsMust() = %v, want %v", ToStringDef(got2), ToStringDef(tt.want))
			}
		})
	}
}

func TestThenByKeyLsMust_2(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelel[int]]
		keySelector func(elelel[int]) int
		lesser      collate.Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "PrimaryOrderingTakesPrecedence",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 10, 20}, elelel[int]{2, 12, 21}, elelel[int]{3, 11, 22}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(1, 3, 2),
		},
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			args: args{
				oe: OrderByKeyMust(NewEnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(3, 1, 2),
		},
		{name: "ThenByAfterOrderByDescending",
			args: args{
				oe: OrderByDescKeyMust(
					NewEnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(2, 3, 1),
		},
		{name: "OrderingIsStable",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 1, 10}, elelel[int]{2, 1, 11}, elelel[int]{3, 1, 11}, elelel[int]{4, 1, 10}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(1, 4, 2, 3),
		},
		{name: "CustomLess",
			args: args{
				oe: OrderByKeyLsMust(
					NewEnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
					collate.Lesser[int](collate.Order[int]{}),
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser: collate.LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewEnSlice(3, 2, 1),
		},
		{name: "CustomComparer",
			args: args{
				oe: OrderByKeyLsMust(
					NewEnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
					collate.Lesser[int](collate.Order[int]{}),
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser: collate.ComparerFunc[int](func(i1, i2 int) int {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
			},
			want: NewEnSlice(3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByKeyLsMust(tt.args.oe, tt.args.keySelector, tt.args.lesser)
			got2 := SelectMust[elelel[int], int](got1, func(e elelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				t.Errorf("ThenByKeyLsMust() = %v, want %v", ToStringDef(got2), ToStringDef(tt.want))
			}
		})
	}
}

func TestThenByKeyMust_3(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelelel[int]]
		keySelector func(elelelel[int]) int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "TertiaryKeys",
			args: args{
				oe: ThenByKeyMust(
					OrderByKeyMust(
						NewEnSlice(elelelel[int]{1, 10, 22, 30}, elelelel[int]{2, 12, 21, 31}, elelelel[int]{3, 10, 20, 33}, elelelel[int]{4, 10, 20, 32}),
						func(e elelelel[int]) int { return e.e2 },
					),
					func(e elelelel[int]) int { return e.e3 },
				),
				keySelector: func(e elelelel[int]) int { return e.e4 },
			},
			want: NewEnSlice(4, 3, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByKeyMust(tt.args.oe, tt.args.keySelector)
			got2 := SelectMust[elelelel[int], int](got1, func(e elelelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				t.Errorf("ThenByKeyMust() = %v, want %v", ToStringDef(got2), ToStringDef(tt.want))
			}
		})
	}
}

func TestThenByDescKeyLsMust_1(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelel[int]]
		keySelector func(elelel[int]) int
		lesser      collate.Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "PrimaryOrderingTakesPrecedence",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 10, 20}, elelel[int]{2, 12, 21}, elelel[int]{3, 11, 22}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(1, 3, 2),
		},
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 10, 19}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(3, 1, 2),
		},
		{name: "ThenByDescendingAfterOrderByDescending",
			args: args{
				oe: OrderByDescKeyMust(
					NewEnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(2, 1, 3),
		},
		{name: "OrderingIsStable",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 1, 10}, elelel[int]{2, 1, 11}, elelel[int]{3, 1, 11}, elelel[int]{4, 1, 10}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      collate.Order[int]{},
			},
			want: NewEnSlice(2, 3, 1, 4),
		},
		{name: "CustomLess",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser: collate.LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "CustomComparer",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser: collate.ComparerFunc[int](func(i1, i2 int) int {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
			},
			want: NewEnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByDescKeyLsMust(tt.args.oe, tt.args.keySelector, tt.args.lesser)
			got2 := SelectMust[elelel[int], int](got1, func(e elelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				t.Errorf("ThenByDescKeyLsMust() = %v, want %v", ToStringDef(got2), ToStringDef(tt.want))
			}
		})
	}
}

func TestThenByDescKeyMust_2(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelelel[int]]
		keySelector func(elelelel[int]) int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "TertiaryKeys",
			args: args{
				oe: ThenByDescKeyMust(
					OrderByKeyMust(
						NewEnSlice(elelelel[int]{1, 10, 22, 30}, elelelel[int]{2, 12, 21, 31}, elelelel[int]{3, 10, 20, 33}, elelelel[int]{4, 10, 20, 32}),
						func(e elelelel[int]) int { return e.e2 },
					),
					func(e elelelel[int]) int { return e.e3 },
				),
				keySelector: func(e elelelel[int]) int { return e.e4 },
			},
			want: NewEnSlice(1, 3, 4, 2),
		},
		{name: "TertiaryKeysWithMixedOrdering",
			args: args{
				oe: ThenByKeyMust(
					OrderByKeyMust(
						NewEnSlice(elelelel[int]{1, 10, 22, 30}, elelelel[int]{2, 12, 21, 31}, elelelel[int]{3, 10, 20, 33}, elelelel[int]{4, 10, 20, 32}),
						func(e elelelel[int]) int { return e.e2 },
					),
					func(e elelelel[int]) int { return e.e3 },
				),
				keySelector: func(e elelelel[int]) int { return e.e4 },
			},
			want: NewEnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByDescKeyMust(tt.args.oe, tt.args.keySelector)
			got2 := SelectMust[elelelel[int], int](got1, func(e elelelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				t.Errorf("ThenByDescKeyMust() = %v, want %v", ToStringDef(got2), ToStringDef(tt.want))
			}
		})
	}
}

func TestThenByKeyMust_string_rune(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[string]
		keySelector func(string) rune
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#secondary-sort-examples
		{name: "Secondary Ascending Sort",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice("the", "quick", "brown", "fox", "jumps"),
					func(s string) int { return len(s) },
				),
				keySelector: func(s string) rune { return []rune(s)[0] },
			},
			want: NewEnSlice("fox", "the", "brown", "jumps", "quick"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThenByKeyMust(tt.args.oe, tt.args.keySelector)
			if !SequenceEqualMust[string](got, tt.want) {
				t.Errorf("ThenByKeyMust() = %v, want %v", ToStringDef[string](got), ToStringDef(tt.want))
			}
		})
	}
}

func TestThenByDescKeyMust_string_rune(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[string]
		keySelector func(string) rune
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#secondary-descending-sort
		{name: "Secondary Ascending Sort",
			args: args{
				oe: OrderByKeyMust(
					NewEnSlice("the", "quick", "brown", "fox", "jumps"),
					func(s string) int { return len(s) },
				),
				keySelector: func(s string) rune { return []rune(s)[0] },
			},
			want: NewEnSlice("the", "fox", "quick", "jumps", "brown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThenByDescKeyMust(tt.args.oe, tt.args.keySelector)
			if !SequenceEqualMust[string](got, tt.want) {
				t.Errorf("ThenByDescKeyMust() = %v, want %v", ToStringDef[string](got), ToStringDef(tt.want))
			}
		})
	}
}

// see the example from Enumerable.ThenBy help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenby
func ExampleThenByMust() {
	fruits := NewEnSlice("grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry")
	// Sort the strings first by their length and then alphabetically.
	query := ThenByMust(
		OrderByKeyMust(fruits, func(fruit string) int { return len(fruit) }),
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
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

// see ThenByDescendingEx1 example from Enumerable.ThenByDescending help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending
func ExampleThenByDescLsMust() {
	fruits := NewEnSlice("apPLe", "baNanA", "apple", "APple", "orange", "BAnana", "ORANGE", "apPLE")
	// Sort the strings first ascending by their length and then descending using a custom case insensitive comparer.
	query := ThenByDescLsMust(
		OrderByKeyMust(fruits, func(fruit string) int { return len(fruit) }),
		collate.CaseInsensitiveLesser,
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
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
