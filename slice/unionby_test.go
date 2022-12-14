package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestUnionBy_string_int(t *testing.T) {
	type args struct {
		first       []string
		second      []string
		keySelector func(string) int
		equaler     go2linq.Equaler[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "1",
			args: args{
				first:       []string{"one", "three", "five"},
				second:      []string{"two", "four"},
				keySelector: func(s string) int { return len(s) },
			},
			want: []string{"one", "three", "five"},
		},
		{name: "2",
			args: args{
				first:       []string{"two", "four"},
				second:      []string{"one", "three", "five"},
				keySelector: func(s string) int { return len(s) },
			},
			want: []string{"two", "four", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionBy(tt.args.first, tt.args.second, tt.args.keySelector, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionBy_Planet(t *testing.T) {
	type args struct {
		first       []Planet
		second      []Planet
		keySelector func(Planet) Planet
		equaler     go2linq.Equaler[Planet]
	}
	tests := []struct {
		name    string
		args    args
		want    []Planet
		wantErr bool
	}{
		{name: "UnionBy",
			args: args{
				first:       []Planet{Mercury, Venus, Earth, Mars, Jupiter},
				second:      []Planet{Mars, Jupiter, Saturn, Uranus, Neptune},
				keySelector: go2linq.Identity[Planet],
			},
			want: []Planet{Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionBy(tt.args.first, tt.args.second, tt.args.keySelector, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionByCmp_int_bool(t *testing.T) {
	e1, _ := Range(1, 10)
	type args struct {
		first       []int
		second      []int
		keySelector func(int) bool
		comparer    go2linq.Comparer[bool]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilBoth",
			args: args{
				first:       nil,
				second:      nil,
				keySelector: func(i int) bool { return i%2 == 0 },
				comparer:    go2linq.BoolComparer,
			},
			want: nil,
		},
		{name: "EmptyBoth",
			args: args{
				first:       []int{},
				second:      []int{},
				keySelector: func(i int) bool { return i%2 == 0 },
				comparer:    go2linq.BoolComparer,
			},
			want: []int{},
		},
		{name: "NilFirst",
			args: args{
				first:       nil,
				second:      []int{2},
				keySelector: func(i int) bool { return i%2 == 0 },
				comparer:    go2linq.BoolComparer,
			},
			want: []int{2},
		},
		{name: "NilSecond",
			args: args{
				first:       []int{1},
				second:      nil,
				keySelector: func(i int) bool { return i%2 == 0 },
				comparer:    go2linq.BoolComparer,
			},
			want: []int{1},
		},
		{name: "NilSelector",
			args: args{
				first:       []int{1},
				second:      []int{2},
				keySelector: nil,
				comparer:    go2linq.BoolComparer,
			},
			wantErr: true,
		},
		{name: "NilComparer",
			args: args{
				first:       []int{1},
				second:      []int{2},
				keySelector: func(i int) bool { return i%2 == 0 },
				comparer:    nil,
			},
			wantErr: true,
		},
		{name: "SameSlice",
			args: args{
				first:       e1,
				second:      e1,
				keySelector: func(i int) bool { return i%2 == 0 },
				comparer:    go2linq.BoolComparer,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionByCmp(tt.args.first, tt.args.second, tt.args.keySelector, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionByCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionByCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
