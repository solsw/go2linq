package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestIntersectMust_int(t *testing.T) {
	e1 := []int{1, 2, 3, 4}
	e2 := []int{1, 2, 3, 4}
	e3 := []int{1, 2, 3, 4}
	type args struct {
		first   []int
		second  []int
		equaler go2linq.Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1",
			args: args{
				first:  []int{1, 2},
				second: []int{2, 3},
			},
			want: []int{2},
		},
		{name: "IntWithoutComparer",
			args: args{
				first:  []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
				second: []int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
			},
			want: []int{4, 5, 6, 7, 8},
		},
		{name: "IntComparerSpecified",
			args: args{
				first:   []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
				second:  []int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
				equaler: go2linq.Order[int]{}},
			want: []int{4, 5, 6, 7, 8},
		},
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: []int{1, 2, 3, 4},
		},
		{name: "SameEnumerable2",
			args: args{
				first:  e2,
				second: e2[1:],
			},
			want: []int{2, 3, 4},
		},
		{name: "SameEnumerable3",
			args: args{
				first:  e3[3:],
				second: e3,
			},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectMust(tt.args.first, tt.args.second, tt.args.equaler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectMust_string(t *testing.T) {
	type args struct {
		first   []string
		second  []string
		equaler go2linq.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "NoComparerSpecified",
			args: args{
				first:  []string{"A", "a", "b", "c", "b"},
				second: []string{"b", "a", "d", "a"},
			},
			want: []string{"a", "b"},
		},
		{name: "Intersect",
			args: args{
				first:  []string{"Mercury", "Venus", "Earth", "Jupiter"},
				second: []string{"Mercury", "Earth", "Mars", "Jupiter"},
			},
			want: []string{"Mercury", "Earth", "Jupiter"},
		},
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:   []string{"A", "a", "b", "c", "b"},
				second:  []string{"b", "a", "d", "a"},
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: []string{"A", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectMust(tt.args.first, tt.args.second, tt.args.equaler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectCmpMust_int(t *testing.T) {
	e1 := []int{4, 3, 2, 1}
	e2 := []int{1, 2, 3, 4}
	e3 := []int{1, 2, 3, 4}
	type args struct {
		first    []int
		second   []int
		comparer go2linq.Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:    []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
				second:   []int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
				comparer: go2linq.Order[int]{},
			},
			want: []int{4, 5, 6, 7, 8},
		},
		{name: "SameEnumerable1",
			args: args{
				first:    e1,
				second:   e1,
				comparer: go2linq.Order[int]{},
			},
			want: []int{4, 3, 2, 1},
		},
		{name: "SameEnumerable2",
			args: args{
				first:    e2,
				second:   e2[1:],
				comparer: go2linq.Order[int]{},
			},
			want: []int{2, 3, 4},
		},
		{name: "SameEnumerable3",
			args: args{
				first:    e3[3:],
				second:   e3,
				comparer: go2linq.Order[int]{},
			},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectCmpMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectCmpMust_string(t *testing.T) {
	type args struct {
		first    []string
		second   []string
		comparer go2linq.Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:    []string{"A", "a", "b", "c", "b"},
				second:   []string{"b", "a", "d", "a"},
				comparer: go2linq.CaseInsensitiveComparer,
			},
			want: []string{"A", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectCmpMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
