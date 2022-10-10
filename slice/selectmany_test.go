//go:build go1.18

package slice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSelectManyMust_int_rune(t *testing.T) {
	type args struct {
		source   []int
		selector func(int) []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{name: "SimpleFlatten",
			args: args{
				source: []int{3, 5, 20, 15},
				selector: func(x int) []rune {
					return []rune(fmt.Sprint(x))
				},
			},
			want: []rune{'3', '5', '2', '0', '1', '5'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyMust(tt.args.source, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyMust_int_int(t *testing.T) {
	type args struct {
		source   []int
		selector func(int) []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "SimpleFlatten1",
			args: args{
				source: []int{1, 2, 3, 4},
				selector: func(i int) []int {
					return []int{i, i * i}
				},
			},
			want: []int{1, 1, 2, 4, 3, 9, 4, 16},
		},
		{name: "SimpleFlatten2",
			args: args{
				source: []int{1, 2, 3, 4},
				selector: func(i int) []int {
					if i%2 == 0 {
						return []int{}
					}
					return []int{i, i * i}
				},
			},
			want: []int{1, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyMust(tt.args.source, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyMust_string_string(t *testing.T) {
	type args struct {
		source   []string
		selector func(string) []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "SelectMany",
			args: args{
				source: []string{"an apple a day", "the quick brown fox"},
				selector: func(s string) []string {
					return strings.Fields(s)
				},
			},
			want: []string{"an", "apple", "a", "day", "the", "quick", "brown", "fox"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyMust(tt.args.source, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyIdxMust_int_rune(t *testing.T) {
	type args struct {
		source   []int
		selector func(int, int) []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{name: "SimpleFlatten",
			args: args{
				source: []int{3, 5, 20, 15},
				selector: func(x, idx int) []rune {
					return []rune(fmt.Sprint(x + idx))
				},
			},
			want: []rune{'3', '6', '2', '2', '1', '8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyIdxMust(tt.args.source, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyIdxMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyIdxMust_int_int(t *testing.T) {
	type args struct {
		source   []int
		selector func(int, int) []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "SimpleFlatten",
			args: args{
				source: []int{1, 2, 3, 4},
				selector: func(x, idx int) []int {
					if idx%2 == 0 {
						return []int{}
					}
					return []int{x, x * x}
				},
			},
			want: []int{2, 4, 4, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyIdxMust(tt.args.source, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyIdxMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyCollMust_int_rune_string(t *testing.T) {
	type args struct {
		source             []int
		collectionSelector func(int) []rune
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "FlattenWithProjection",
			args: args{
				source: []int{3, 5, 20, 15},
				collectionSelector: func(x int) []rune {
					return []rune(fmt.Sprint(x))
				},
				resultSelector: func(i int, r rune) string {
					return fmt.Sprintf("%d: %s", i, string(r))
				},
			},
			want: []string{"3: 3", "5: 5", "20: 2", "20: 0", "15: 1", "15: 5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyCollMust(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyCollMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyCollIdxMust_int_rune_string(t *testing.T) {
	type args struct {
		source             []int
		collectionSelector func(int, int) []rune
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "FlattenWithProjectionAndIndex",
			args: args{
				source: []int{3, 5, 20, 15},
				collectionSelector: func(x, idx int) []rune {
					return []rune(fmt.Sprint(x + idx))
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: []string{"3: 3", "5: 6", "20: 2", "20: 2", "15: 1", "15: 8"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyCollIdxMust(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyCollIdxMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
