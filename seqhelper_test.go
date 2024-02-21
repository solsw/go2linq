package go2linq

import (
	"context"
	"fmt"
	"iter"
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/solsw/errorhelper"
)

func TestStringFmt_int(t *testing.T) {
	type args struct {
		seq   iter.Seq[int]
		sep   string
		lrim  string
		rrim  string
		ledge string
		redge string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				seq:   VarAll(1, 2, 3, 4),
				sep:   "-",
				lrim:  "<",
				rrim:  ">",
				ledge: "[",
				redge: "]",
			},
			want: "[<1>-<2>-<3>-<4>]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFmt(tt.args.seq, tt.args.sep, tt.args.lrim, tt.args.rrim, tt.args.ledge, tt.args.redge); got != tt.want {
				t.Errorf("StringFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

type intStringer int

func (i intStringer) String() string {
	return fmt.Sprintf("%d+%d", i, i*i)
}

func TestStringDef(t *testing.T) {
	type args struct {
		seq iter.Seq[intStringer]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				seq: VarAll(intStringer(1), intStringer(2), intStringer(3)),
			},
			want: "[1+1 2+4 3+9]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringDef(tt.args.seq); got != tt.want {
				t.Errorf("StringDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringDef_any(t *testing.T) {
	type args struct {
		seq iter.Seq[any]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				seq: VarAll(any(intStringer(1)), any(2), any(intStringer(3))),
			},
			want: "[1+1 2 3+9]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringDef(tt.args.seq); got != tt.want {
				t.Errorf("StringDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach_int(t *testing.T) {
	var acc1 int
	ctx1, cancel := context.WithCancel(context.Background())
	type args struct {
		ctx    context.Context
		seq    iter.Seq[int]
		action func(int) error
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "01",
			args: args{
				ctx: context.Background(),
				action: func(i int) error {
					acc1 += i * i
					return nil
				},
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "02",
			args: args{
				ctx: context.Background(),
				seq: VarAll(1, 2, 3),
			},
			wantErr:     true,
			expectedErr: ErrNilAction,
		},
		{name: "03",
			args: args{
				ctx: ctx1,
				seq: VarAll(1, 2, 3),
				action: func(i int) error {
					if i == 2 {
						cancel()
					}
					return nil
				},
			},
			wantErr:     true,
			expectedErr: context.Canceled,
		},
		{name: "04",
			args: args{
				ctx: context.Background(),
				seq: VarAll(1, 2, 3),
				action: func(i int) error {
					if i == 2 {
						return ErrTestError
					}
					acc1 += i * i
					return nil
				},
			},
			wantErr:     true,
			expectedErr: ErrTestError,
		},
		{name: "1",
			args: args{
				ctx: context.Background(),
				seq: VarAll(1, 2, 3),
				action: func(i int) error {
					acc1 += i * i
					return nil
				},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc1 = 0
			err := ForEach(tt.args.ctx, tt.args.seq, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEach() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ForEach() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(acc1, tt.want) {
				t.Errorf("ForEach() = %v, want %v", acc1, tt.want)
			}
		})
	}
}

func TestForEachConcurrent_int(t *testing.T) {
	var acc1 int64
	canceledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	type args struct {
		ctx    context.Context
		seq    iter.Seq[int]
		action func(int) error
	}
	tests := []struct {
		name        string
		args        args
		want        int64
		wantErr     bool
		expectedErr error
	}{
		{name: "01",
			args: args{
				ctx:    canceledCtx,
				seq:    VarAll(1, 2, 3),
				action: func(int) error { return nil },
			},
			wantErr:     true,
			expectedErr: context.Canceled,
		},
		{name: "02",
			args: args{
				ctx: context.Background(),
				seq: VarAll(1, 2, 3),
				action: func(i int) error {
					if i == 2 {
						return ErrTestError
					}
					atomic.AddInt64(&acc1, int64(i*i))
					return nil
				},
			},
			wantErr:     true,
			expectedErr: ErrTestError,
		},
		{name: "1",
			args: args{
				ctx: context.Background(),
				seq: errorhelper.Must(Range(1, 1000)),
				action: func(i int) error {
					// acc1 += int64(i * i) // <- demonstrates race error
					atomic.AddInt64(&acc1, int64(i*i))
					return nil
				},
			},
			want: 333833500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc1 = 0
			err := ForEachConcurrent(tt.args.ctx, tt.args.seq, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEachConcurrent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ForEachConcurrent() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(acc1, tt.want) {
				t.Errorf("ForEachConcurrent() = %v, want %v", acc1, tt.want)
			}
		})
	}
}

func TestSeqString_int(t *testing.T) {
	type args struct {
		seq iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "1",
			args: args{
				seq: VarAll(1, 2, 3),
			},
			want: VarAll("1", "2", "3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SeqString(tt.args.seq)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SeqString() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSeqString_any(t *testing.T) {
	type args struct {
		seq iter.Seq[any]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "1",
			args: args{
				seq: VarAll(any(1), any(intStringer(2)), any(3)),
			},
			want: VarAll("1", "2+4", "3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SeqString(tt.args.seq)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SeqString() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestStrings_int(t *testing.T) {
	type args struct {
		seq iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",
			args: args{
				seq: VarAll(1, 2, 3),
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Strings(tt.args.seq)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrings_intStringer(t *testing.T) {
	type args struct {
		seq iter.Seq[intStringer]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",
			args: args{
				seq: VarAll(intStringer(1), 2, 3),
			},
			want: []string{"1+1", "2+4", "3+9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Strings(tt.args.seq)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
