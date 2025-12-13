package go2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/generichelper"
	"github.com/solsw/iterhelper"
)

func TestIndex_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
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
				source: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "EmptySource",
			args: args{
				source: Empty[int](),
			},
			want: Empty[generichelper.Tuple2[int, int]](),
		},
		{name: "RegularSource",
			args: args{
				source: iterhelper.Var(1, 2),
			},
			want: iterhelper.Var(generichelper.NewTuple2(0, 1), generichelper.NewTuple2(1, 2)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Index(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Index() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Index() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestIndex_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[generichelper.Tuple2[int, string]]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSource",
			args: args{
				source: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "EmptySource",
			args: args{
				source: Empty[string](),
			},
			want: Empty[generichelper.Tuple2[int, string]](),
		},
		{name: "RegularSource",
			args: args{
				source: iterhelper.Var("one", "two"),
			},
			want: iterhelper.Var(
				generichelper.NewTuple2(0, "one"),
				generichelper.NewTuple2(1, "two"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Index(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Index() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Index() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}
