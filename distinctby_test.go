//go:build go1.18

package go2linq

import (
	"testing"
)

func Test_DistinctBy_string_int(t *testing.T) {
	type args struct {
		source      Enumerable[string]
		keySelector func(string) int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSource",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSelector",
			args: args{
				source: Empty[string](),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "1",
			args: args{
				source:      NewEnSlice("one", "two", "three", "four", "five"),
				keySelector: func(s string) int { return len(s) },
			},
			want: NewEnSlice("one", "three", "four"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctBy(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctBy() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("DistinctBy() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctBy() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_DistinctByMust_Planet_PlanetType(t *testing.T) {
	type args struct {
		source      Enumerable[Planet]
		keySelector func(Planet) PlanetType
	}
	tests := []struct {
		name string
		args args
		want Enumerable[Planet]
	}{
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#distinct-and-distinctby
		{name: "DistinctBy",
			args: args{
				source:      NewEnSlice(Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto),
				keySelector: func(p Planet) PlanetType { return p.Type },
			},
			want: NewEnSlice(Mercury, Jupiter, Uranus, Pluto),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctBy(tt.args.source, tt.args.keySelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctBy() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_DistinctByEqMust_string_int(t *testing.T) {
	type args struct {
		source      Enumerable[string]
		keySelector func(string) int
		equaler     Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "1",
			args: args{
				source:      NewEnSlice("one", "two", "three", "four", "five"),
				keySelector: func(s string) int { return len(s) % 2 },
				equaler:     Equaler[int](Order[int]{}),
			},
			want: NewEnSlice("one", "four"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistinctByEqMust(tt.args.source, tt.args.keySelector, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctByEqMust() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_DistinctByCmpMust_string_rune(t *testing.T) {
	type args struct {
		source      Enumerable[string]
		keySelector func(string) rune
		comparer    Comparer[rune]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "1",
			args: args{
				source:      NewEnSlice("one", "two", "three", "four", "five"),
				keySelector: func(s string) rune { return []rune(s)[0] },
				comparer:    Comparer[rune](Order[rune]{}),
			},
			want: NewEnSlice("one", "two", "four"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistinctByCmpMust(tt.args.source, tt.args.keySelector, tt.args.comparer)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctByCmpMust() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}