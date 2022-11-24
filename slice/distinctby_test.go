package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestDistinctBy_string_int(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:      nil,
				keySelector: func(s string) int { return len(s) },
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source:      []string{},
				keySelector: func(s string) int { return len(s) },
			},
			want: []string{},
		},
		{name: "NilSelector",
			args: args{
				source:      []string{"one", "two", "three", "four", "five"},
				keySelector: nil,
			},
			wantErr: true,
		},
		{name: "DistinctBy",
			args: args{
				source:      []string{"one", "two", "three", "four", "five"},
				keySelector: func(s string) int { return len(s) },
			},
			want: []string{"one", "three", "four"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctBy(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctByEqMust_string_int(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) int
		equaler     go2linq.Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "DistinctByEqMust",
			args: args{
				source:      []string{"one", "two", "three", "four", "five"},
				keySelector: func(s string) int { return len(s) % 2 },
				equaler:     go2linq.Equaler[int](go2linq.Order[int]{}),
			},
			want: []string{"one", "four"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctByEqMust(tt.args.source, tt.args.keySelector, tt.args.equaler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctByEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctByCmpMust_string_rune(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) rune
		comparer    go2linq.Comparer[rune]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "DistinctByCmpMust",
			args: args{
				source:      []string{"one", "two", "three", "four", "five"},
				keySelector: func(s string) rune { return []rune(s)[0] },
				comparer:    go2linq.Comparer[rune](go2linq.Order[rune]{}),
			},
			want: []string{"one", "two", "four"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctByCmpMust(tt.args.source, tt.args.keySelector, tt.args.comparer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctByCmpMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
