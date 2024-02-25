package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/generichelper"
)

func TestApplyResultSelector(t *testing.T) {
	lk := Lookup[int, string]{KeyEqual: generichelper.DeepEqual[int]}
	lk.Add(3, "abc")
	lk.Add(3, "def")
	lk.Add(1, "x")
	lk.Add(1, "y")
	lk.Add(3, "ghi")
	lk.Add(1, "z")
	lk.Add(2, "00")
	type args struct {
		lookup         *Lookup[int, string]
		resultSelector func(int, iter.Seq[string]) string
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "00",
			args:        args{lookup: nil, resultSelector: func(i int, ss iter.Seq[string]) string { return fmt.Sprintf("%d:%v", i, ss) }},
			want:        VarAll("12345"),
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "01",
			args:        args{lookup: &lk, resultSelector: nil},
			want:        VarAll("12345"),
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "1",
			args: args{
				lookup:         &lk,
				resultSelector: func(i int, ss iter.Seq[string]) string { return fmt.Sprintf("%d:%v", i, StringDef(ss)) },
			},
			want: VarAll("3:[abc def ghi]", "1:[x y z]", "2:[00]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyResultSelector(tt.args.lookup, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyResultSelector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ApplyResultSelector() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("ApplyResultSelector() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
