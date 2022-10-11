//go:build go1.18

package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

var (
	testString1 = "test"
	testString2 = "test"
)

func TestDistinct_int(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source: nil,
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distinct(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distinct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctMust_string(t *testing.T) {
	type args struct {
		source []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "DistinctMust",
			args: args{
				source: []string{"A", "a", "b", "c", "b"},
			},
			want: []string{"A", "a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctMust(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctEq_int(t *testing.T) {
	type args struct {
		source  []int
		equaler go2linq.Equaler[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "DistinctEq",
			args: args{
				source:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				equaler: go2linq.EqualerFunc[int](func(i1, i2 int) bool { return i1%2 == i2%2 }),
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctEq(tt.args.source, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctEq_string(t *testing.T) {
	type args struct {
		source  []string
		equaler go2linq.Equaler[string]
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:  nil,
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source: []string{},
			},
			want: []string{},
		},
		{name: "NilEqualerUsesDeepEqualer",
			args: args{
				source:  []string{"xyz", testString1, "XYZ", testString2, "def"},
				equaler: nil,
			},
			want: []string{"xyz", testString1, "XYZ", "def"},
		},
		{name: "NonNullEqualer",
			args: args{
				source:  []string{"xyz", testString1, "XYZ", testString2, "def"},
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: []string{"xyz", testString1, "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctEq(tt.args.source, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctCmpMust_int(t *testing.T) {
	type args struct {
		source   []int
		comparer go2linq.Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "NilSource",
			args: args{
				source:   nil,
				comparer: go2linq.Order[int]{},
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source:   []int{},
				comparer: go2linq.Order[int]{},
			},
			want: []int{},
		},
		{name: "DistinctCmpMust",
			args: args{
				source:   []int{1, 2, 3, 4},
				comparer: go2linq.Order[int]{},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctCmpMust(tt.args.source, tt.args.comparer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctCmpMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctCmpMust_string(t *testing.T) {
	type args struct {
		source   []string
		comparer go2linq.Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "DistinctCmpMust1",
			args: args{
				source:   []string{"xyz", testString1, "XYZ", testString2, "def"},
				comparer: go2linq.CaseInsensitiveComparer,
			},
			want: []string{"xyz", testString1, "def"},
		},
		{name: "DistinctCmpMust2",
			args: args{
				source:   []string{"A", "a", "b", "c", "b"},
				comparer: go2linq.CaseInsensitiveComparer,
			},
			want: []string{"A", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctCmpMust(tt.args.source, tt.args.comparer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctCmpMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
