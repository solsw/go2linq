package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/collate"
)

func TestDistinctBy(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) int
		equaler     collate.Equaler[int]
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
		{name: "DistinctByWithoutEqualer",
			args: args{
				source:      []string{"one", "two", "three", "four", "five"},
				keySelector: func(s string) int { return len(s) },
			},
			want: []string{"one", "three", "four"},
		},
		{name: "DistinctByWithEqualer",
			args: args{
				source:      []string{"one", "two", "three", "four", "five"},
				keySelector: func(s string) int { return len(s) % 2 },
				equaler:     collate.Equaler[int](collate.Order[int]{}),
			},
			want: []string{"one", "four"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctBy(tt.args.source, tt.args.keySelector, tt.args.equaler)
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

func TestDistinctByCmp(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) rune
		comparer    collate.Comparer[rune]
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "DistinctByCmp",
			args: args{
				source:      []string{"one", "two", "three", "four", "five"},
				keySelector: func(s string) rune { return []rune(s)[0] },
				comparer:    collate.Comparer[rune](collate.Order[rune]{}),
			},
			want: []string{"one", "two", "four"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctByCmp(tt.args.source, tt.args.keySelector, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctByCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctByCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
