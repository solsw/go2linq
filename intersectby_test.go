package go2linq

import (
	"iter"
	"testing"

	"github.com/solsw/iterhelper"
)

func TestIntersectBy_Planet(t *testing.T) {
	type args struct {
		first       iter.Seq[Planet]
		second      iter.Seq[Planet]
		keySelector func(Planet) Planet
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[Planet]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#intersect-and-intersectby
		{name: "IntersectBy",
			args: args{
				first:       iterhelper.Var(Mercury, Venus, Earth, Mars, Jupiter),
				second:      iterhelper.Var(Mars, Jupiter, Saturn, Uranus, Neptune),
				keySelector: Identity[Planet],
			},
			want: iterhelper.Var(Mars, Jupiter),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IntersectBy(tt.args.first, tt.args.second, tt.args.keySelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("IntersectBy() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}
