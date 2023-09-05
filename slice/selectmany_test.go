package slice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSelectMany_int_rune(t *testing.T) {
	type args struct {
		source   []int
		selector func(int) []rune
	}
	tests := []struct {
		name    string
		args    args
		want    []rune
		wantErr bool
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
			got, err := SelectMany(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectMany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectMany_int_int(t *testing.T) {
	type args struct {
		source   []int
		selector func(int) []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
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
			got, err := SelectMany(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectMany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectMany_string_string(t *testing.T) {
	type args struct {
		source   []string
		selector func(string) []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
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
			got, err := SelectMany(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectMany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyIdx_int_rune(t *testing.T) {
	type args struct {
		source   []int
		selector func(int, int) []rune
	}
	tests := []struct {
		name    string
		args    args
		want    []rune
		wantErr bool
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
			got, err := SelectManyIdx(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectManyIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyIdx_int_int(t *testing.T) {
	type args struct {
		source   []int
		selector func(int, int) []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
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
			got, err := SelectManyIdx(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectManyIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyColl_int_rune_string(t *testing.T) {
	type args struct {
		source             []int
		collectionSelector func(int) []rune
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
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
			got, err := SelectManyColl(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectManyColl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyColl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectManyCollIdx_int_rune_string(t *testing.T) {
	type args struct {
		source             []int
		collectionSelector func(int, int) []rune
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
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
			got, err := SelectManyCollIdx(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectManyCollIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectManyCollIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}
