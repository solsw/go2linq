package go2linq

import (
	"iter"
	"testing"
)

func TestExceptBy_Planet(t *testing.T) {
	type args struct {
		first       iter.Seq[Planet]
		second      iter.Seq[Planet]
		keySelector func(Planet) string
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[Planet]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilFirst",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSecond",
			args: args{
				first: VarAll(Mercury, Venus, Earth, Jupiter),
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSelector",
			args: args{
				first:       VarAll(Mercury, Venus, Earth, Jupiter),
				second:      VarAll(Mercury, Earth, Mars, Jupiter),
				keySelector: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#except-and-exceptby
		{name: "ExceptBy",
			args: args{
				first:       VarAll(Mercury, Venus, Earth, Jupiter),
				second:      VarAll(Mercury, Earth, Mars, Jupiter),
				keySelector: func(planet Planet) string { return planet.Name },
			},
			want: VarAll(Venus),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got iter.Seq[Planet]
			enr2, err := Select(tt.args.second, tt.args.keySelector)
			if err == nil {
				got, err = ExceptBy(tt.args.first, enr2, tt.args.keySelector)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ExceptBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ExceptBy() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("ExceptBy() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
