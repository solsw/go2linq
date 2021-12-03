//go:build go1.18

package go2linq

import (
	"math"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByDescendingTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ThenByDescendingTest.cs

func Test_OrderByLsMust_int(t *testing.T) {
	type args struct {
		source      Enumerator[int]
		keySelector func(int) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "1234",
			args: args{
				source:      NewOnSlice(4, 1, 3, 2),
				keySelector: Identity[int],
				lesser:      IntLesser,
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderByLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser).GetEnumerator()
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OrderByLs() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_OrderByLsMust_intint(t *testing.T) {
	type args struct {
		source      Enumerator[elel[int]]
		keySelector func(elel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      NewOnSlice(elel[int]{1, 10}, elel[int]{2, 12}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(1, 3, 2),
		},
		{name: "OrderingIsStable",
			args: args{
				source:      NewOnSlice(elel[int]{1, 10}, elel[int]{2, 11}, elel[int]{3, 11}, elel[int]{4, 10}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(1, 4, 2, 3),
		},
		{name: "CustomLess",
			args: args{
				source:      NewOnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewOnSlice(3, 2, 1),
		},
		{name: "CustomComparer",
			args: args{
				source:      NewOnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
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
			want: NewOnSlice(3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectMust(
				OrderByLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser).GetEnumerator(),
				func(e elel[int]) int { return e.e1 },
			)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OrderByLs() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_OrderByDescendingLsMust_intint(t *testing.T) {
	type args struct {
		source      Enumerator[elel[int]]
		keySelector func(elel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      NewOnSlice(elel[int]{1, 10}, elel[int]{2, 12}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(2, 3, 1),
		},
		{name: "OrderingIsStable",
			args: args{
				source:      NewOnSlice(elel[int]{1, 10}, elel[int]{2, 11}, elel[int]{3, 11}, elel[int]{4, 10}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(2, 3, 1, 4),
		},
		{name: "CustomLess",
			args: args{
				source:      NewOnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewOnSlice(1, 2, 3),
		},
		{name: "CustomComparer",
			args: args{
				source:      NewOnSlice(elel[int]{1, 15}, elel[int]{2, -13}, elel[int]{3, 11}),
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
			want: NewOnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectMust(
				OrderByDescendingLsMust(tt.args.source, tt.args.keySelector, tt.args.lesser).GetEnumerator(),
				func(e elel[int]) int { return e.e1 },
			)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OrderByDescendingLs() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ThenByLsMust_1(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelel[int]]
		keySelector func(elelel[int]) elelel[int]
		lesser      Lesser[elelel[int]]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: Identity[elelel[int]],
				lesser:      LesserFunc[elelel[int]](func(e1, e2 elelel[int]) bool { return e1.e3 < e2.e3 }),
			},
			want: NewOnSlice(3, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByLsMust(tt.args.oe, tt.args.keySelector, tt.args.lesser).GetEnumerator()
			got2 := SelectMust(got1, func(e elelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				got2.Reset()
				tt.want.Reset()
				t.Errorf("ThenByLs() = '%v', want '%v'", String(got2), String(tt.want))
			}
		})
	}
}

func Test_ThenByLsMust_2(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelel[int]]
		keySelector func(elelel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "PrimaryOrderingTakesPrecedence",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 10, 20}, elelel[int]{2, 12, 21}, elelel[int]{3, 11, 22}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(1, 3, 2),
		},
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			args: args{
				oe: OrderByLsMust(NewOnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(3, 1, 2),
		},
		{name: "ThenByAfterOrderByDescending",
			args: args{
				oe: OrderByDescendingLsMust(
					NewOnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(2, 3, 1),
		},
		{name: "OrderingIsStable",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 1, 10}, elelel[int]{2, 1, 11}, elelel[int]{3, 1, 11}, elelel[int]{4, 1, 10}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(1, 4, 2, 3),
		},
		{name: "CustomLess",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser: LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewOnSlice(3, 2, 1),
		},
		{name: "CustomComparer",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
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
			want: NewOnSlice(3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByLsMust(tt.args.oe, tt.args.keySelector, tt.args.lesser).GetEnumerator()
			got2 := SelectMust(got1, func(e elelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				got2.Reset()
				tt.want.Reset()
				t.Errorf("ThenByLs() = '%v', want '%v'", String(got2), String(tt.want))
			}
		})
	}
}

func Test_ThenByLsMust_3(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelelel[int]]
		keySelector func(elelelel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "TertiaryKeys",
			args: args{
				oe: ThenByLsMust(
					OrderByLsMust(
						NewOnSlice(elelelel[int]{1, 10, 22, 30}, elelelel[int]{2, 12, 21, 31}, elelelel[int]{3, 10, 20, 33}, elelelel[int]{4, 10, 20, 32}),
						func(e elelelel[int]) int { return e.e2 },
						IntLesser,
					),
					func(e elelelel[int]) int { return e.e3 },
					IntLesser,
				),
				keySelector: func(e elelelel[int]) int { return e.e4 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(4, 3, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByLsMust(tt.args.oe, tt.args.keySelector, tt.args.lesser).GetEnumerator()
			got2 := SelectMust(got1, func(e elelelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				got2.Reset()
				tt.want.Reset()
				t.Errorf("ThenByLs() = '%v', want '%v'", String(got2), String(tt.want))
			}
		})
	}
}

func Test_ThenByDescendingLsMust_1(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelel[int]]
		keySelector func(elelel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "PrimaryOrderingTakesPrecedence",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 10, 20}, elelel[int]{2, 12, 21}, elelel[int]{3, 11, 22}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(1, 3, 2),
		},
		{name: "SecondOrderingIsUsedWhenPrimariesAreEqual",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 10, 19}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(3, 1, 2),
		},
		{name: "ThenByDescendingAfterOrderByDescending",
			args: args{
				oe: OrderByDescendingLsMust(
					NewOnSlice(elelel[int]{1, 10, 22}, elelel[int]{2, 12, 21}, elelel[int]{3, 10, 20}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(2, 1, 3),
		},
		{name: "OrderingIsStable",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 1, 10}, elelel[int]{2, 1, 11}, elelel[int]{3, 1, 11}, elelel[int]{4, 1, 10}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(2, 3, 1, 4),
		},
		{name: "CustomLess",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
				lesser: LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: NewOnSlice(1, 2, 3),
		},
		{name: "CustomComparer",
			args: args{
				oe: OrderByLsMust(
					NewOnSlice(elelel[int]{1, 1, 15}, elelel[int]{2, 1, -13}, elelel[int]{3, 1, 11}),
					func(e elelel[int]) int { return e.e2 },
					IntLesser,
				),
				keySelector: func(e elelel[int]) int { return e.e3 },
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
			want: NewOnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByDescendingLsMust(tt.args.oe, tt.args.keySelector, tt.args.lesser).GetEnumerator()
			got2 := SelectMust(got1, func(e elelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				got2.Reset()
				tt.want.Reset()
				t.Errorf("ThenByDescendingLs() = '%v', want '%v'", String(got2), String(tt.want))
			}
		})
	}
}

func Test_ThenByDescendingLsMust_2(t *testing.T) {
	type args struct {
		oe          *OrderedEnumerable[elelelel[int]]
		keySelector func(elelelel[int]) int
		lesser      Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "TertiaryKeys",
			args: args{
				oe: ThenByDescendingLsMust(
					OrderByLsMust(
						NewOnSlice(elelelel[int]{1, 10, 22, 30}, elelelel[int]{2, 12, 21, 31}, elelelel[int]{3, 10, 20, 33}, elelelel[int]{4, 10, 20, 32}),
						func(e elelelel[int]) int { return e.e2 },
						IntLesser,
					),
					func(e elelelel[int]) int { return e.e3 },
					IntLesser,
				),
				keySelector: func(e elelelel[int]) int { return e.e4 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(1, 3, 4, 2),
		},
		{name: "TertiaryKeysWithMixedOrdering",
			args: args{
				oe: ThenByLsMust(
					OrderByLsMust(
						NewOnSlice(elelelel[int]{1, 10, 22, 30}, elelelel[int]{2, 12, 21, 31}, elelelel[int]{3, 10, 20, 33}, elelelel[int]{4, 10, 20, 32}),
						func(e elelelel[int]) int { return e.e2 },
						IntLesser,
					),
					func(e elelelel[int]) int { return e.e3 },
					IntLesser,
				),
				keySelector: func(e elelelel[int]) int { return e.e4 },
				lesser:      IntLesser,
			},
			want: NewOnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := ThenByDescendingLsMust(tt.args.oe, tt.args.keySelector, tt.args.lesser).GetEnumerator()
			got2 := SelectMust(got1, func(e elelelel[int]) int { return e.e1 })
			if !SequenceEqualMust(got2, tt.want) {
				got2.Reset()
				tt.want.Reset()
				t.Errorf("ThenByDescendingLs() = '%v', want '%v'", String(got2), String(tt.want))
			}
		})
	}
}
