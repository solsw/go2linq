package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/collate"
)

func TestUnion_int(t *testing.T) {
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
		{name: "SameSlice1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: []int{1, 2, 3, 4},
		},
		{name: "SameSlice2",
			args: args{
				first:  e2[:1],
				second: e2[3:],
			},
			want: []int{1, 4},
		},
		{name: "SameSlice3",
			args: args{
				first:  e3[2:],
				second: e3,
			},
			want: []int{3, 4, 1, 2},
		},
		{name: "UnionWithIntEquality",
			args: args{
				first:   []int{1, 2},
				second:  []int{2, 3},
				equaler: collate.Order[int]{},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Union(tt.args.first, tt.args.second, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion_string(t *testing.T) {
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
		{name: "TwoNils",
			args: args{
				first:  nil,
				second: nil,
			},
			want: nil,
		},
		{name: "UnionWithTwoEmptySlices",
			args: args{
				first:  []string{},
				second: []string{},
			},
			want: []string{},
		},
		{name: "FirstNil",
			args: args{
				first:  nil,
				second: []string{"one", "two", "three", "four"},
			},
			want: []string{"one", "two", "three", "four"},
		},
		{name: "FirstEmpty",
			args: args{
				first:  []string{},
				second: []string{"one", "two", "three", "four"},
			},
			want: []string{"one", "two", "three", "four"},
		},
		{name: "SecondNil",
			args: args{
				first:  []string{"one", "two", "three", "four"},
				second: nil,
			},
			want: []string{"one", "two", "three", "four"},
		},
		{name: "SecondEmpty",
			args: args{
				first:  []string{"one", "two", "three", "four"},
				second: []string{},
			},
			want: []string{"one", "two", "three", "four"},
		},
		{name: "UnionWithoutComparer",
			args: args{
				first:  []string{"a", "b", "B", "c", "b"},
				second: []string{"d", "e", "d", "a"},
			},
			want: []string{"a", "b", "B", "c", "d", "e"},
		},
		{name: "UnionWithoutComparer2",
			args: args{
				first:  []string{"a", "b"},
				second: []string{"b", "a"},
			},
			want: []string{"a", "b"},
		},
		{name: "Union",
			args: args{
				first:  []string{"Mercury", "Venus", "Earth", "Jupiter"},
				second: []string{"Mercury", "Earth", "Mars", "Jupiter"},
			},
			want: []string{"Mercury", "Venus", "Earth", "Jupiter", "Mars"},
		},
		{name: "UnionWithCaseInsensitiveComparerEq",
			args: args{
				first:   []string{"a", "b", "B", "c", "b"},
				second:  []string{"d", "e", "d", "a"},
				equaler: collate.CaseInsensitiveOrder,
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Union(tt.args.first, tt.args.second, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionCmp_int(t *testing.T) {
	e1 := []int{1, 2, 3, 4}
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
		{name: "NilComparer",
			args: args{
				first:    []int{1, 2, 2},
				second:   []int{},
				comparer: nil,
			},
			wantErr: true,
		},
		{name: "UnionWithIntComparer1",
			args: args{
				first:    []int{1, 2, 2},
				second:   []int{},
				comparer: collate.Order[int]{},
			},
			want: []int{1, 2},
		},
		{name: "UnionWithIntComparer2",
			args: args{
				first:    []int{1, 2},
				second:   []int{2, 3},
				comparer: collate.Order[int]{},
			},
			want: []int{1, 2, 3},
		},
		{name: "SameSlice1",
			args: args{
				first:    e1,
				second:   e1,
				comparer: collate.Order[int]{},
			},
			want: []int{1, 2, 3, 4},
		},
		{name: "SameSlice2",
			args: args{
				first:    e2[2:],
				second:   e2[:1],
				comparer: collate.Order[int]{},
			},
			want: []int{3, 4, 1},
		},
		{name: "SameSlice3",
			args: args{
				first:    e3[2:],
				second:   e3,
				comparer: collate.Order[int]{},
			},
			want: []int{3, 4, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionCmp(tt.args.first, tt.args.second, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionCmp_string(t *testing.T) {
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
		{name: "UnionWithCaseInsensitiveComparerCmp",
			args: args{
				first:    []string{"a", "b", "B", "c", "b"},
				second:   []string{"d", "e", "d", "a"},
				comparer: collate.CaseInsensitiveOrder,
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionCmp(tt.args.first, tt.args.second, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
