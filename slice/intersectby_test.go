package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestIntersectByMust_Planet(t *testing.T) {
	type args struct {
		first       []Planet
		second      []Planet
		keySelector func(Planet) Planet
	}
	tests := []struct {
		name string
		args args
		want []Planet
	}{
		{name: "IntersectBy",
			args: args{
				first:       []Planet{Mercury, Venus, Earth, Mars, Jupiter},
				second:      []Planet{Mars, Jupiter, Saturn, Uranus, Neptune},
				keySelector: go2linq.Identity[Planet],
			},
			want: []Planet{Mars, Jupiter},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectByMust(tt.args.first, tt.args.second, tt.args.keySelector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectByMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
