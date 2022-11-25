package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestExceptBy_Planet(t *testing.T) {
	type args struct {
		first       []Planet
		second      []Planet
		keySelector func(Planet) Planet
	}
	tests := []struct {
		name    string
		args    args
		want    []Planet
		wantErr bool
	}{
		{name: "NilFirst",
			args: args{
				first:  nil,
				second: []Planet{Mercury, Earth, Mars, Jupiter},
			},
			want: nil,
		},
		{name: "NilSecond",
			args: args{
				first:  []Planet{Mercury, Venus, Earth, Jupiter},
				second: nil,
			},
			want: []Planet{Mercury, Venus, Earth, Jupiter},
		},
		{name: "NilSelector",
			args: args{
				first:  []Planet{Mercury, Venus, Earth, Jupiter},
				second: []Planet{Mercury, Earth, Mars, Jupiter},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExceptBy(tt.args.first, tt.args.second, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExceptBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExceptByEqMust(t *testing.T) {
	type args struct {
		first       []Planet
		second      []Planet
		keySelector func(Planet) Planet
		equaler     go2linq.Equaler[Planet]
	}
	tests := []struct {
		name string
		args args
		want []Planet
	}{
		{name: "ExceptByEqMust",
			args: args{
				first:       []Planet{Mercury, Venus, Earth, Jupiter},
				second:      []Planet{Mercury, Earth, Mars, Jupiter},
				keySelector: go2linq.Identity[Planet],
				equaler:     go2linq.EqualerFunc[Planet](func(p1, p2 Planet) bool { return p1.Name == p2.Name }),
			},
			want: []Planet{Venus},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceptByEqMust(tt.args.first, tt.args.second, tt.args.keySelector, tt.args.equaler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExceptByEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
