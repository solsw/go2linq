//go:build go1.18

package go2linq

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectManyTest.cs

func Test_SelectManyMust_int_rune(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int) Enumerable[rune]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[rune]
	}{
		{name: "SimpleFlatten",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				selector: func(x int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x))...)
				},
			},
			want: NewEnSlice('3', '5', '2', '0', '1', '5'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SelectManyMust_int_int(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int) Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				selector: func(i int) Enumerable[int] {
					return NewEnSlice(i, i*i)
				},
			},
			want: NewEnSlice(1, 1, 2, 4, 3, 9, 4, 16),
		},
		{name: "2",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				selector: func(i int) Enumerable[int] {
					if i%2 == 0 {
						return Empty[int]()
					}
					return NewEnSlice(i, i*i)
				},
			},
			want: NewEnSlice(1, 1, 3, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SelectManyMust_string_string(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#selectmany
		{name: "SelectMany",
			args: args{
				source: NewEnSlice("an apple a day", "the quick brown fox"),
				selector: func(s string) Enumerable[string] {
					return NewEnSlice(strings.Fields(s)...)
				},
			},
			want: NewEnSlice("an", "apple", "a", "day", "the", "quick", "brown", "fox"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SelectManyIdxMust_int_rune(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int, int) Enumerable[rune]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[rune]
	}{
		{name: "SimpleFlattenWithIndex",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				selector: func(x, index int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x + index))...)
				},
			},
			want: NewEnSlice('3', '6', '2', '2', '1', '8'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyIdxMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SelectManyIdxMust_int_int(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int, int) Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				selector: func(i, idx int) Enumerable[int] {
					if idx%2 == 0 {
						return Empty[int]()
					}
					return NewEnSlice(i, i*i)
				},
			},
			want: NewEnSlice(2, 4, 4, 16),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyIdxMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SelectManyCollMust_int_rune_string(t *testing.T) {
	type args struct {
		source             Enumerable[int]
		collectionSelector func(int) Enumerable[rune]
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "FlattenWithProjection",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				collectionSelector: func(x int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x))...)
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: NewEnSlice("3: 3", "5: 5", "20: 2", "20: 0", "15: 1", "15: 5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyCollMust(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyCollMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SelectManyCollIdxMust_int_rune_string(t *testing.T) {
	type args struct {
		source             Enumerable[int]
		collectionSelector func(int, int) Enumerable[rune]
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "FlattenWithProjectionAndIndex",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				collectionSelector: func(x, index int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x + index))...)
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: NewEnSlice("3: 3", "5: 6", "20: 2", "20: 2", "15: 1", "15: 8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyCollIdxMust(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyCollIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}
