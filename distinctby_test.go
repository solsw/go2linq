package go2linq

import (
	"cmp"
	"iter"
	"testing"
)

func TestDistinctBy_string_int(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSource",
			args: args{
				source:      nil,
				keySelector: func(s string) int { return len(s) },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSelector",
			args: args{
				source:      Empty[string](),
				keySelector: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "1",
			args: args{
				source:      VarAll("one", "two", "three", "four", "five"),
				keySelector: func(s string) int { return len(s) },
			},
			want: VarAll("one", "three", "four"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctBy(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("DistinctBy() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctBy_Planet_PlanetType(t *testing.T) {
	type args struct {
		source      iter.Seq[Planet]
		keySelector func(Planet) PlanetType
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[Planet]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#distinct-and-distinctby
		{name: "DistinctBy",
			args: args{
				source:      VarAll(Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto),
				keySelector: func(p Planet) PlanetType { return p.Type },
			},
			want: VarAll(Mercury, Jupiter, Uranus, Pluto),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctBy(tt.args.source, tt.args.keySelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctByEq_string_int(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
		equal       func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "1",
			args: args{
				source:      VarAll("one", "two", "three", "four", "five"),
				keySelector: func(s string) int { return len(s) % 2 },
				equal:       func(i1, i2 int) bool { return i1 == i2 },
			},
			want: VarAll("one", "four"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctByEq(tt.args.source, tt.args.keySelector, tt.args.equal)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DistinctByEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctByCmp_string_rune(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) rune
		compare     func(rune, rune) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "1",
			args: args{
				source:      VarAll("one", "two", "three", "four", "five"),
				keySelector: func(s string) rune { return []rune(s)[0] },
				compare:     cmp.Compare[rune],
			},
			want: VarAll("one", "two", "four"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctByCmp(tt.args.source, tt.args.keySelector, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DistinctByCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
