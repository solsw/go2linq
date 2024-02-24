package go2linq

import (
	"iter"
	"testing"
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
				first:       VarAll(Mercury, Venus, Earth, Mars, Jupiter),
				second:      VarAll(Mars, Jupiter, Saturn, Uranus, Neptune),
				keySelector: Identity[Planet],
			},
			want: VarAll(Mars, Jupiter),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IntersectBy(tt.args.first, tt.args.second, tt.args.keySelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("IntersectBy() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
