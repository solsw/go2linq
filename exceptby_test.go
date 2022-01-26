//go:build go1.18

package go2linq

import (
	"testing"
)

func Test_ExceptBy2_Planet(t *testing.T) {
	PlanetNameSelector := func(planet Planet) string { return planet.Name }
	type args struct {
		first       Enumerable[Planet]
		second      Enumerable[Planet]
		keySelector func(Planet) string
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[Planet]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilFirst",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSecond",
			args: args{
				first: NewEnSlice(Mercury, Venus, Earth, Jupiter),
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSelector",
			args: args{
				first:  NewEnSlice(Mercury, Venus, Earth, Jupiter),
				second: NewEnSlice(Mercury, Earth, Mars, Jupiter),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#except-and-exceptby
		{name: "ExceptBy",
			args: args{
				first:       NewEnSlice(Mercury, Venus, Earth, Jupiter),
				second:      NewEnSlice(Mercury, Earth, Mars, Jupiter),
				keySelector: PlanetNameSelector,
			},
			want: NewEnSlice(Venus),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Enumerable[Planet]
			enr2, err := Select(tt.args.second, tt.args.keySelector)
			if err == nil {
				got, err = ExceptBy(tt.args.first, enr2, tt.args.keySelector)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ExceptBy() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ExceptBy() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptBy() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}
