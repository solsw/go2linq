package go2linq

import (
	"cmp"
	"iter"
	"strconv"
	"testing"
)

func TestUnionBy_string_int(t *testing.T) {
	type args struct {
		first       iter.Seq[string]
		second      iter.Seq[string]
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "1",
			args: args{
				first:       VarAll("one", "three", "five"),
				second:      VarAll("two", "four"),
				keySelector: func(s string) int { return len(s) },
			},
			want: VarAll("one", "three", "five"),
		},
		{name: "2",
			args: args{
				first:       VarAll("two", "four"),
				second:      VarAll("one", "three", "five"),
				keySelector: func(s string) int { return len(s) },
			},
			want: VarAll("two", "four", "three"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionBy(tt.args.first, tt.args.second, tt.args.keySelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("UnionBy() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestUnionBy_Planet(t *testing.T) {
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
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#union-and-unionby
		{name: "UnionBy",
			args: args{
				first:       VarAll(Mercury, Venus, Earth, Mars, Jupiter),
				second:      VarAll(Mars, Jupiter, Saturn, Uranus, Neptune),
				keySelector: Identity[Planet],
			},
			want: VarAll(Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionBy(tt.args.first, tt.args.second, tt.args.keySelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("UnionBy() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestUnionByCmp_int_string(t *testing.T) {
	e1, _ := Range(1, 10)
	type args struct {
		first       iter.Seq[int]
		second      iter.Seq[int]
		keySelector func(int) string
		compare     func(string, string) int
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilFirst",
			args: args{
				first:       nil,
				second:      e1,
				keySelector: nil,
				compare:     nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSecond",
			args: args{
				first:       e1,
				second:      nil,
				keySelector: nil,
				compare:     nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilSelector",
			args: args{
				first:       e1,
				second:      e1,
				keySelector: nil,
				compare:     nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "NilComparer",
			args: args{
				first:       e1,
				second:      e1,
				keySelector: func(i int) string { return strconv.Itoa(i) },
				compare:     nil,
			},
			wantErr:     true,
			expectedErr: ErrNilCompare,
		},
		{name: "SameEnumerable1",
			args: args{
				first:       e1,
				second:      e1,
				keySelector: func(i int) string { return strconv.FormatBool(i%2 == 0) },
				compare:     cmp.Compare[string],
			},
			want: VarAll(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionByCmp(tt.args.first, tt.args.second, tt.args.keySelector, tt.args.compare)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionByCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("UnionByCmp() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("UnionByCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
