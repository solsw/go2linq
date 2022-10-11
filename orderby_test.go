//go:build go1.18

package go2linq

import (
	"fmt"
	"math"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByDescendingTest.cs

func TestOrderByMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1234",
			args: args{
				source: NewEnSlice(4, 1, 3, 2),
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderByMust(tt.args.source)
			if !SequenceEqualMust[int](got, tt.want) {
				t.Errorf("OrderByMust() = %v, want %v", ToStringDef[int](got), ToStringDef(tt.want))
			}
		})
	}
}

func TestOrderByKeyLsMust_intint(t *testing.T) {
	type args struct {
		source      Enumerable[elel[int]]
		keySelector func(elel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      NewEnSlice(elel[int]{1, 10}, elel[int]{2, 12}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      Order[int]{},
			},
			want: NewEnSlice(1, 3, 2),
		},
		{name: "OrderingIsStable",
			args: args{
				source:      NewEnSlice(elel[int]{1, 10}, elel[int]{2, 11}, elel[int]{3, 11}, elel[int]{4, 10}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      Order[int]{},
			},
			want: NewEnSlice(1, 4, 2, 3),
		},
		{name: "CustomLess",
			args: args{
				source:      NewEnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewEnSlice(3, 2, 1),
		},
		{name: "CustomComparer",
			args: args{
				source:      NewEnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: ComparerFunc[int](func(i1, i2 int) int {
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
			got := SelectMust[elel[int], int](
				OrderByKeyLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser),
				func(e elel[int]) int { return e.e1 },
			)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OrderByKeyLsMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestOrderByKeyDescLsMust_intint(t *testing.T) {
	type args struct {
		source      Enumerable[elel[int]]
		keySelector func(elel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      NewEnSlice(elel[int]{1, 10}, elel[int]{2, 12}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      Order[int]{},
			},
			want: NewEnSlice(2, 3, 1),
		},
		{name: "OrderingIsStable",
			args: args{
				source:      NewEnSlice(elel[int]{1, 10}, elel[int]{2, 11}, elel[int]{3, 11}, elel[int]{4, 10}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      Order[int]{},
			},
			want: NewEnSlice(2, 3, 1, 4),
		},
		{name: "CustomLess",
			args: args{
				source:      NewEnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "CustomComparer",
			args: args{
				source:      NewEnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: ComparerFunc[int](func(i1, i2 int) int {
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
			got := SelectMust[elel[int], int](
				OrderByKeyDescLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser),
				func(e elel[int]) int { return e.e1 },
			)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OrderByKeyDescLsMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestOrderByKeyMust_string_int(t *testing.T) {
	type args struct {
		source      Enumerable[string]
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#primary-ascending-sort
		{name: "Primary Ascending Sort",
			args: args{
				source:      NewEnSlice("the", "quick", "brown", "fox", "jumps"),
				keySelector: func(s string) int { return len(s) },
			},
			want: NewEnSlice("the", "fox", "quick", "brown", "jumps"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderByKeyMust(tt.args.source, tt.args.keySelector)
			if !SequenceEqualMust[string](got, tt.want) {
				t.Errorf("OrderByKeyMust() = %v, want %v", ToStringDef[string](got), ToStringDef(tt.want))
			}
		})
	}
}

func TestOrderByKeyDescMust_string_rune(t *testing.T) {
	type args struct {
		source      Enumerable[string]
		keySelector func(string) rune
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/sorting-data#primary-descending-sort
		{name: "Primary Descending Sort",
			args: args{
				source:      NewEnSlice("the", "quick", "brown", "fox", "jumps"),
				keySelector: func(s string) rune { return []rune(s)[0] },
			},
			want: NewEnSlice("the", "quick", "jumps", "fox", "brown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderByKeyDescMust(tt.args.source, tt.args.keySelector)
			if !SequenceEqualMust[string](got, tt.want) {
				t.Errorf("OrderByKeyDescMust() = %v, want %v", ToStringDef[string](got), ToStringDef(tt.want))
			}
		})
	}
}

func ExampleOrderByMust() {
	fmt.Println(ToStringDef[string](
		OrderByMust(
			NewEnSlice("zero", "one", "two", "three", "four", "five"),
		),
	))
	// Output:
	// [five four one three two zero]
}

// see OrderByEx1 example from Enumerable.OrderBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby
func ExampleOrderByLsMust() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	var ls Lesser[Pet] = LesserFunc[Pet](func(p1, p2 Pet) bool { return p1.Age < p2.Age })
	query := OrderByLsMust(pets, ls)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		pet := enr.Current()
		fmt.Printf("%s - %d\n", pet.Name, pet.Age)
	}
	// Output:
	// Whiskers - 1
	// Boots - 4
	// Barley - 8
}

// see OrderByDescendingEx1 example from Enumerable.OrderByDescending help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
func ExampleOrderByDescLsMust() {
	decimals := NewEnSlice(6.2, 8.3, 0.5, 1.3, 6.3, 9.7)
	var ls Lesser[float64] = LesserFunc[float64](
		func(f1, f2 float64) bool {
			_, fr1 := math.Modf(f1)
			_, fr2 := math.Modf(f2)
			if math.Abs(fr1-fr2) < 0.001 {
				return f1 < f2
			}
			return fr1 < fr2
		},
	)
	query := OrderByDescLsMust(decimals, ls)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		num := enr.Current()
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

func ExampleOrderByKeyDescMust() {
	fmt.Println(ToStringDef[string](
		OrderByKeyDescMust(
			NewEnSlice("zero", "one", "two", "three", "four", "five"),
			func(s string) int { return len(s) },
		),
	))
	// Output:
	// [three zero four five one two]
}
