//go:build go1.18

package go2linq

import (
	"testing"
)

func Test_UnionByMust_string_int(t *testing.T) {
	type args struct {
		first       Enumerator[string]
		second      Enumerator[string]
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "1",
			args: args{
				first:       NewOnSlice("one", "three", "five"),
				second:      NewOnSlice("two", "four"),
				keySelector: func(s string) int { return len(s) },
			},
			want: NewOnSlice("one", "three", "five"),
		},
		{name: "2",
			args: args{
				first:       NewOnSlice("two", "four"),
				second:      NewOnSlice("one", "three", "five"),
				keySelector: func(s string) int { return len(s) },
			},
			want: NewOnSlice("two", "four", "three"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionByMust(tt.args.first, tt.args.second, tt.args.keySelector)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionByMust() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionByMust_Planet(t *testing.T) {
	type args struct {
		first       Enumerator[Planet]
		second      Enumerator[Planet]
		keySelector func(Planet) Planet
	}
	tests := []struct {
		name string
		args args
		want Enumerator[Planet]
	}{
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#union-and-unionby
		{name: "UnionBy",
			args: args{
				first:       NewOnSlice(Mercury, Venus, Earth, Mars, Jupiter),
				second:      NewOnSlice(Mars, Jupiter, Saturn, Uranus, Neptune),
				keySelector: Identity[Planet],
			},
			want: NewOnSlice(Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionByMust(tt.args.first, tt.args.second, tt.args.keySelector)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionByMust() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_UnionByCmpSelf_int_bool(t *testing.T) {
	e1 := RangeMust(1, 10)
	type args struct {
		first       Enumerator[int]
		second      Enumerator[int]
		keySelector func(int) bool
		comparer    Comparer[bool]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilFirst",
			args: args{
				second: e1,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSecond",
			args: args{
				first: e1,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSelector",
			args: args{
				first:  e1,
				second: e1,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "NilComparer",
			args: args{
				first:       e1,
				second:      e1,
				keySelector: func(i int) bool { return i%2 == 0 },
			},
			wantErr:     true,
			expectedErr: ErrNilComparer,
		},
		{name: "SameEnumerable1",
			args: args{
				first:       e1,
				second:      e1,
				keySelector: func(i int) bool { return i%2 == 0 },
				comparer:    BoolComparer,
			},
			want: NewOnSlice(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionByCmpSelf(tt.args.first, tt.args.second, tt.args.keySelector, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionByCmpSelf() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("UnionByCmpSelf() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("UnionByCmpSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
