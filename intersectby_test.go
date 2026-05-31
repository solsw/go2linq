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

// distinct elements of 'first' that share a key must be deduplicated by key,
// yielding only the first occurrence (matching .NET's set semantics).
func TestIntersectBy_DupKeyInFirst(t *testing.T) {
	type kv struct {
		k int
		v string
	}
	first := iterhelper.Var(kv{1, "a"}, kv{1, "b"}, kv{2, "c"})
	second := iterhelper.Var(1)
	got, _ := IntersectBy(first, second, func(x kv) int { return x.k })
	want := iterhelper.Var(kv{1, "a"})
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("IntersectBy() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(want))
	}
}
