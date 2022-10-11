//go:build go1.18

package go2linq

import (
	"testing"
)

func TestIntersectByMust_Planet(t *testing.T) {
	type args struct {
		first       Enumerable[Planet]
		second      Enumerable[Planet]
		keySelector func(Planet) Planet
	}
	tests := []struct {
		name string
		args args
		want Enumerable[Planet]
	}{
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#intersect-and-intersectby
		{name: "IntersectBy",
			args: args{
				first:       NewEnSlice(Mercury, Venus, Earth, Mars, Jupiter),
				second:      NewEnSlice(Mars, Jupiter, Saturn, Uranus, Neptune),
				keySelector: Identity[Planet],
			},
			want: NewEnSlice(Mars, Jupiter),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectByMust(tt.args.first, tt.args.second, tt.args.keySelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectByMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}
