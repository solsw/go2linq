package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v3"
)

func TestIntersectBy(t *testing.T) {
	type args struct {
		first       []Planet
		second      []Planet
		keySelector func(Planet) Planet
		equaler     collate.Equaler[Planet]
	}
	tests := []struct {
		name    string
		args    args
		want    []Planet
		wantErr bool
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
			got, err := IntersectBy(tt.args.first, tt.args.second, tt.args.keySelector, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntersectBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
