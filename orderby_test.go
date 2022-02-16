//go:build go1.18

package go2linq

import (
	"fmt"
	"math"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByDescendingTest.cs

func Test_OrderByLsMust_int(t *testing.T) {
	type args struct {
		source      Enumerable[int]
		keySelector func(int) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1234",
			args: args{
				source:      NewEnSlice(4, 1, 3, 2),
				keySelector: Identity[int],
				lesser:      Order[int]{},
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderByLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser)
			if !SequenceEqualMust[int](got, tt.want) {
				t.Errorf("OrderByLsMust() = %v, want %v", ToStringDef[int](got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_OrderByLsMust_intint(t *testing.T) {
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
				OrderByLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser),
				func(e elel[int]) int { return e.e1 },
			)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OrderByLsMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_OrderByDescendingLsMust_intint(t *testing.T) {
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
				OrderByDescendingLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser),
				func(e elel[int]) int { return e.e1 },
			)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OrderByDescendingLsMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_OrderByLsMust_string_int(t *testing.T) {
	type args struct {
		source      Enumerable[string]
		keySelector func(string) int
		lesser      Lesser[int]
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
				lesser:      Order[int]{},
			},
			want: NewEnSlice("the", "fox", "quick", "brown", "jumps"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderByLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser)
			if !SequenceEqualMust[string](got, tt.want) {
				t.Errorf("OrderByLsMust() = %v, want %v", ToStringDef[string](got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_OrderByDescendingLsMust_string_rune(t *testing.T) {
	type args struct {
		source      Enumerable[string]
		keySelector func(string) rune
		lesser      Lesser[rune]
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
				lesser:      Order[rune]{},
			},
			want: NewEnSlice("the", "quick", "jumps", "fox", "brown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderByDescendingLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser)
			if !SequenceEqualMust[string](got, tt.want) {
				t.Errorf("OrderByDescendingLsMust() = %v, want %v", ToStringDef[string](got), ToStringDef(tt.want))
			}
		})
	}
}

func Example_OrderBySelfMust() {
	fmt.Println(ToStringDef[string](
		OrderBySelfMust(
			NewEnSlice("zero", "one", "two", "three", "four", "five"),
		),
	))
	// Output:
	// [five four one three two zero]
}

func Example_OrderByDescendingMust() {
	fmt.Println(ToStringDef[string](
		OrderByDescendingMust(
			NewEnSlice("zero", "one", "two", "three", "four", "five"),
			func(s string) int { return len(s) },
		),
	))
	// Output:
	// [three zero four five one two]
}
