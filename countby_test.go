package go2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/generichelper"
	"github.com/solsw/iterhelper"
)

func TestCountBy_string_int(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[generichelper.Tuple2[int, int]]
	}{
		{name: "Regular",
			args: args{
				source:      iterhelper.VarSeq("one", "two", "three", "four", "five", "six"),
				keySelector: func(s string) int { return len(s) },
			},
			want: iterhelper.VarSeq(
				generichelper.NewTuple2(3, 3),
				generichelper.NewTuple2(5, 1),
				generichelper.NewTuple2(4, 2),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := CountBy(tt.args.source, tt.args.keySelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("CountBy() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestCountByEq_string_int(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
		keyEqual    func(int, int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[generichelper.Tuple2[int, int]]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSource",
			args: args{
				source:      nil,
				keySelector: func(s string) int { return len(s) },
				keyEqual:    func(a, b int) bool { return a == b },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilKeySelector",
			args: args{
				source:      iterhelper.VarSeq("one", "two", "three", "four", "five", "six"),
				keySelector: nil,
				keyEqual:    func(a, b int) bool { return a == b },
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "NilEqual",
			args: args{
				source:      iterhelper.VarSeq("one", "two", "three", "four", "five", "six"),
				keySelector: func(s string) int { return len(s) },
				keyEqual:    nil,
			},
			wantErr:     true,
			expectedErr: ErrNilEqual,
		},
		{name: "Regular",
			args: args{
				source:      iterhelper.VarSeq("one", "two", "three", "four", "five", "six"),
				keySelector: func(s string) int { return len(s) },
				keyEqual:    func(a, b int) bool { return a == b },
			},
			want: iterhelper.VarSeq(
				generichelper.NewTuple2(3, 3),
				generichelper.NewTuple2(5, 1),
				generichelper.NewTuple2(4, 2),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountByEq(tt.args.source, tt.args.keySelector, tt.args.keyEqual)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountByEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("CountByEq() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("CountByEq() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}
