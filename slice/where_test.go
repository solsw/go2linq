package slice

import (
	"reflect"
	"strings"
	"testing"
)

func TestWhere_int(t *testing.T) {
	type args struct {
		source    []int
		predicate func(int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:    nil,
				predicate: func(i int) bool { return i > 5 },
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source:    []int{},
				predicate: func(i int) bool { return i > 5 },
			},
			want: []int{},
		},
		{name: "NilPredicate",
			args: args{
				source:    []int{1, 3, 4, 2, 8, 1},
				predicate: nil,
			},
			wantErr: true,
		},
		{name: "SimpleFiltering",
			args: args{
				source:    []int{1, 3, 4, 2, 8, 1},
				predicate: func(i int) bool { return i < 4 },
			},
			want: []int{1, 3, 2, 1},
		},
		{name: "AlwaysFalsePredicate",
			args: args{
				source:    []int{1, 2, 3, 4},
				predicate: func(int) bool { return false },
			},
			want: []int{},
		},
		{name: "AlwaysTruePredicate",
			args: args{
				source:    []int{1, 2, 3, 4},
				predicate: func(int) bool { return true },
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Where(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Where() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhereMust_string(t *testing.T) {
	type args struct {
		source    []string
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "AlwaysTruePredicate",
			args: args{
				source:    []string{"one", "two", "three", "four", "five"},
				predicate: func(string) bool { return true },
			},
			want: []string{"one", "two", "three", "four", "five"},
		},
		{name: "SimpleFiltering1",
			args: args{
				source:    []string{"one", "two", "three", "four", "five"},
				predicate: func(s string) bool { return strings.HasPrefix(s, "t") },
			},
			want: []string{"two", "three"},
		},
		{name: "SimpleFiltering2",
			args: args{
				source:    []string{"one", "two", "three", "four", "five"},
				predicate: func(s string) bool { return len(s) == 3 },
			},
			want: []string{"one", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhereMust(tt.args.source, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WhereMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhereIdx_int(t *testing.T) {
	type args struct {
		source    []int
		predicate func(int, int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSourceWithIndex",
			args: args{
				source:    nil,
				predicate: func(x, _ int) bool { return x > 5 },
			},
			want: nil,
		},
		{name: "EmptySourceWithIndex",
			args: args{
				source:    []int{},
				predicate: func(x, _ int) bool { return x > 5 },
			},
			want: []int{},
		},
		{name: "NilPredicateWithIndex",
			args: args{
				source:    []int{1, 3, 4, 2, 8, 1},
				predicate: nil,
			},
			wantErr: true,
		},
		{name: "SimpleFilteringWithIndex",
			args: args{
				source:    []int{1, 3, 4, 2, 8, 1},
				predicate: func(x, idx int) bool { return x < idx },
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WhereIdx(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("WhereIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WhereIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhereIdxMust_string(t *testing.T) {
	type args struct {
		source    []string
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "SimpleFiltering",
			args: args{
				source:    []string{"one", "two", "three", "four", "five"},
				predicate: func(s string, idx int) bool { return len(s) == idx },
			},
			want: []string{"five"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhereIdxMust(tt.args.source, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WhereIdxMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
