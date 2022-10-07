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
