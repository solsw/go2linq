package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestExceptMust_int(t *testing.T) {
	i4 := []int{1, 2, 3, 4}
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
		{name: "IntWithoutComparer",
			args: args{
				first:  []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
				second: []int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
			},
			want: []int{1, 2, 3},
		},
		{name: "IdenticalSlices",
			args: args{
				first:  []int{1, 2, 3, 4},
				second: []int{1, 2, 3, 4},
			},
			want: []int{},
		},
		{name: "IdenticalSlices2",
			args: args{
				first:  []int{1, 2, 3, 4},
				second: []int{1, 2, 3, 4}[2:],
			},
			want: []int{1, 2},
		},
		{name: "SameSlice",
			args: args{
				first:  i4,
				second: i4[2:],
			},
			want: []int{1, 2},
		},
		{name: "IntComparerSpecified",
			args: args{
				first:   []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
				second:  []int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
				equaler: go2linq.Order[int]{},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptMust(tt.args.first, tt.args.second, tt.args.equaler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExceptMust_string(t *testing.T) {
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
				first:  []string{"A", "a", "b", "c", "b", "c"},
				second: []string{"b", "a", "d", "a"},
			},
			want: []string{"A", "c"},
		},
		{name: "ExceptMust",
			args: args{
				first:  []string{"Mercury", "Venus", "Earth", "Jupiter"},
				second: []string{"Mercury", "Earth", "Mars", "Jupiter"},
			},
			want: []string{"Venus"},
		},
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:   []string{"A", "a", "b", "c", "b"},
				second:  []string{"b", "a", "d", "a"},
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: []string{"c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptMust(tt.args.first, tt.args.second, tt.args.equaler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExceptCmpMust_int(t *testing.T) {
	i4 := []int{1, 2, 3, 4}
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
			want: []int{1, 2, 3},
		},
		{name: "SameSlice",
			args: args{
				first:    i4,
				second:   i4[2:],
				comparer: go2linq.Order[int]{},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptCmpMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExceptCmpMust_string(t *testing.T) {
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
			want: []string{"c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptCmpMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
