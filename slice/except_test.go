package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/collate"
)

func TestExcept_int(t *testing.T) {
	i4 := []int{1, 2, 3, 4}
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
				equaler: collate.Order[int]{},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Except(tt.args.first, tt.args.second, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Except() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Except() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExcept_string(t *testing.T) {
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
				first:  []string{"A", "a", "b", "c", "b", "c"},
				second: []string{"b", "a", "d", "a"},
			},
			want: []string{"A", "c"},
		},
		{name: "Except",
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
				equaler: collate.CaseInsensitiveEqualer,
			},
			want: []string{"c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Except(tt.args.first, tt.args.second, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Except() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Except() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExceptCmp_int(t *testing.T) {
	i4 := []int{1, 2, 3, 4}
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
			want: []int{1, 2, 3},
		},
		{name: "SameSlice",
			args: args{
				first:    i4,
				second:   i4[2:],
				comparer: collate.Order[int]{},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExceptCmp(tt.args.first, tt.args.second, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExceptCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExceptCmp_string(t *testing.T) {
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
			want: []string{"c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExceptCmp(tt.args.first, tt.args.second, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExceptCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
