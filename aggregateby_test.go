package go2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/generichelper"
)

func TestAggregateBy_string_int_int(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
		seed        int
		accumulator func(int, string) int
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[generichelper.Tuple2[int, int]]
		wantErr bool
	}{
		{name: "Regular",
			args: args{
				source:      VarToSeq("one", "two", "three", "four", "five", "six"),
				keySelector: func(s string) int { return len(s) },
				seed:        0,
				accumulator: func(ac int, el string) int { return ac + len(el) },
			},
			want: VarToSeq(
				generichelper.NewTuple2(3, 9),
				generichelper.NewTuple2(5, 5),
				generichelper.NewTuple2(4, 8),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateBy(tt.args.source, tt.args.keySelector, tt.args.seed, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("AggregateBy() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestAggregateByEq_string_int_int(t *testing.T) {
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
		seed        int
		accumulator func(int, string) int
		keyEqual    func(int, int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[generichelper.Tuple2[int, int]]
		wantErr bool
	}{
		{name: "Regular",
			args: args{
				source:      VarToSeq("one", "two", "three", "four", "five", "six"),
				keySelector: func(s string) int { return len(s) },
				seed:        0,
				accumulator: func(ac int, el string) int { return ac + len(el) },
				keyEqual:    func(a, b int) bool { return a == b },
			},
			want: VarToSeq(
				generichelper.NewTuple2(3, 9),
				generichelper.NewTuple2(5, 5),
				generichelper.NewTuple2(4, 8),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateByEq(tt.args.source, tt.args.keySelector, tt.args.seed, tt.args.accumulator, tt.args.keyEqual)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateByEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("AggregateByEq() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestAggregateBySelEq_string_int_int(t *testing.T) {
	type args struct {
		source       iter.Seq[string]
		keySelector  func(string) int
		seedSelector func(int) int
		accumulator  func(int, string) int
		keyEqual     func(int, int) bool
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
				source:       nil,
				keySelector:  func(s string) int { return len(s) },
				seedSelector: func(k int) int { return k },
				accumulator:  func(ac int, el string) int { return ac + len(el) },
				keyEqual:     func(a, b int) bool { return a == b },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilKeySelector",
			args: args{
				source:       VarToSeq("one", "two", "three", "four", "five", "six"),
				keySelector:  nil,
				seedSelector: func(k int) int { return k },
				accumulator:  func(ac int, el string) int { return ac + len(el) },
				keyEqual:     func(a, b int) bool { return a == b },
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "NilSeedSelector",
			args: args{
				source:       VarToSeq("one", "two", "three", "four", "five", "six"),
				keySelector:  func(s string) int { return len(s) },
				seedSelector: nil,
				accumulator:  func(ac int, el string) int { return ac + len(el) },
				keyEqual:     func(a, b int) bool { return a == b },
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "NilAccumulator",
			args: args{
				source:       VarToSeq("one", "two", "three", "four", "five", "six"),
				keySelector:  func(s string) int { return len(s) },
				seedSelector: func(k int) int { return k },
				accumulator:  nil,
				keyEqual:     func(a, b int) bool { return a == b },
			},
			wantErr:     true,
			expectedErr: ErrNilAccumulator,
		},
		{name: "NilEqual",
			args: args{
				source:       VarToSeq("one", "two", "three", "four", "five", "six"),
				keySelector:  func(s string) int { return len(s) },
				seedSelector: func(k int) int { return k },
				accumulator:  func(ac int, el string) int { return ac + len(el) },
				keyEqual:     nil,
			},
			wantErr:     true,
			expectedErr: ErrNilEqual,
		},
		{name: "Regular",
			args: args{
				source:       VarToSeq("one", "two", "three", "four", "five", "six"),
				keySelector:  func(s string) int { return len(s) },
				seedSelector: func(k int) int { return k },
				accumulator:  func(ac int, el string) int { return ac + len(el) },
				keyEqual:     func(a, b int) bool { return a == b },
			},
			want: VarToSeq(
				generichelper.NewTuple2(3, 12),
				generichelper.NewTuple2(5, 10),
				generichelper.NewTuple2(4, 12),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateBySelEq(tt.args.source, tt.args.keySelector, tt.args.seedSelector, tt.args.accumulator, tt.args.keyEqual)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateBySelEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("AggregateByEq() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("AggregateBySelEq() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
