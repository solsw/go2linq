package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/collate"
)

func TestIntersect_int(t *testing.T) {
	e1 := []int{1, 2, 3, 4}
	e2 := []int{1, 2, 3, 4}
	e3 := []int{1, 2, 3, 4}
	type args struct {
		first   []int
		second  []int
		equaler collate.Equaler[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
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
				equaler: collate.Order[int]{}},
			want: []int{4, 5, 6, 7, 8},
		},
		{name: "SameSlice1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: []int{1, 2, 3, 4},
		},
		{name: "SameSlice2",
			args: args{
				first:  e2,
				second: e2[1:],
			},
			want: []int{2, 3, 4},
		},
		{name: "SameSlice3",
			args: args{
				first:  e3[3:],
				second: e3,
			},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Intersect(tt.args.first, tt.args.second, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Intersect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect_string(t *testing.T) {
	type args struct {
		first   []string
		second  []string
		equaler collate.Equaler[string]
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
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
				equaler: collate.CaseInsensitiveEqualer,
			},
			want: []string{"A", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Intersect(tt.args.first, tt.args.second, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Intersect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectCmp_int(t *testing.T) {
	e1 := []int{4, 3, 2, 1}
	e2 := []int{1, 2, 3, 4}
	e3 := []int{1, 2, 3, 4}
	type args struct {
		first    []int
		second   []int
		comparer collate.Comparer[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:    []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
				second:   []int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
				comparer: collate.Order[int]{},
			},
			want: []int{4, 5, 6, 7, 8},
		},
		{name: "SameSlice1",
			args: args{
				first:    e1,
				second:   e1,
				comparer: collate.Order[int]{},
			},
			want: []int{4, 3, 2, 1},
		},
		{name: "SameSlice2",
			args: args{
				first:    e2,
				second:   e2[1:],
				comparer: collate.Order[int]{},
			},
			want: []int{2, 3, 4},
		},
		{name: "SameSlice3",
			args: args{
				first:    e3[3:],
				second:   e3,
				comparer: collate.Order[int]{},
			},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntersectCmp(tt.args.first, tt.args.second, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntersectCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectCmp(t *testing.T) {
	type args struct {
		first    []string
		second   []string
		comparer collate.Comparer[string]
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:    []string{"A", "a", "b", "c", "b"},
				second:   []string{"b", "a", "d", "a"},
				comparer: collate.CaseInsensitiveComparer,
			},
			want: []string{"A", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntersectCmp(tt.args.first, tt.args.second, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntersectCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
