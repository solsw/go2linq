//go:build go1.18

package go2linq

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync/atomic"
	"testing"
)

func Test_EnOnChan_int(t *testing.T) {
	in1 := make(chan int)
	go func() {
		for i := 1; i <= 4; i++ {
			in1 <- i
		}
		close(in1)
	}()
	type args struct {
		chn <-chan int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1",
			args: args{
				chn: in1,
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnOnChan[int](tt.args.chn)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("EnOnChan() failed")
			}
		})
	}
}

func TestEnToSlice_int(t *testing.T) {
	type args struct {
		en Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "NilSource",
			args: args{
				en: nil,
			},
			want: nil,
		},
		{name: "SimpleSlice",
			args: args{
				en: NewEnSlice[int](1, 2, 3, 4),
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnToSlice[int](tt.args.en)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnToSlice() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestEnToSliceErr_interface_int(t *testing.T) {
	type args struct {
		en Enumerable[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "CastExceptionOnWrongElementType",
			args: args{
				en: CastMust[any, int](NewEnSlice[any](1.0, 2.0, 3.0, 4.0, "five")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EnToSliceErr[int](tt.args.en)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnToSliceErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnToSliceErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

type intStringer int

func (i intStringer) String() string {
	return fmt.Sprintf("%d+%d", i, i*i)
}

func TestEnToString_Stringer(t *testing.T) {
	type args struct {
		en Enumerable[intStringer]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				en: NewEnSlice(intStringer(1), intStringer(2), intStringer(3)),
			},
			want: "[1+1 2+4 3+9]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnToString(tt.args.en)
			if got != tt.want {
				t.Errorf("EnToString_Stringer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnToStringEn_int(t *testing.T) {
	type args struct {
		en Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "1",
			args: args{
				en: NewEnSlice(1, 2, 3),
			},
			want: NewEnSlice("1", "2", "3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnToStringEn(tt.args.en)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("EnToStringEn() = %v, want %v", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func TestEnToStrings_int(t *testing.T) {
	type args struct {
		en Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",
			args: args{
				en: NewEnSlice(1, 2, 3),
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnToStrings(tt.args.en)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnToStrings_Stringer(t *testing.T) {
	type args struct {
		en Enumerable[intStringer]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",
			args: args{
				en: NewEnSlice(intStringer(1), 2, 3),
			},
			want: []string{"1+1", "2+4", "3+9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnToStrings(tt.args.en)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnToStrings_Stringer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEachEn_int(t *testing.T) {
	var acc1 int
	type args struct {
		ctx    context.Context
		en     Enumerable[int]
		action func(context.Context, int) error
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
				action: func(_ context.Context, i int) error {
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
				en:  NewEnSlice(1, 2, 3),
			},
			wantErr:     true,
			expectedErr: ErrNilAction,
		},
		{name: "03",
			args: args{
				ctx: context.Background(),
				en:  NewEnSlice(1, 2, 3),
				action: func(_ context.Context, i int) error {
					if i == 2 {
						return errors.New("ForEachEn error")
					}
					acc1 += i * i
					return nil
				},
			},
			wantErr:     true,
			expectedErr: errors.New("ForEachEn error"),
		},
		{name: "1",
			args: args{
				ctx: context.Background(),
				en:  NewEnSlice(1, 2, 3),
				action: func(_ context.Context, i int) error {
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
			err := ForEachEn(tt.args.ctx, tt.args.en, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEachEn() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.expectedErr) {
					t.Errorf("ForEachEn() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(acc1, tt.want) {
				t.Errorf("ForEachEn() = %v, want %v", acc1, tt.want)
			}
		})
	}
}

func TestForEachEnConcurrent_int(t *testing.T) {
	var acc1 int64
	type args struct {
		ctx    context.Context
		en     Enumerable[int]
		action func(context.Context, int) error
	}
	tests := []struct {
		name        string
		args        args
		want        int64
		wantErr     bool
		expectedErr error
	}{
		{name: "03",
			args: args{
				ctx: context.Background(),
				en:  NewEnSlice(1, 2, 3),
				action: func(_ context.Context, i int) error {
					if i == 2 {
						return errors.New("ForEachEnConcurrent error")
					}
					atomic.AddInt64(&acc1, int64(i*i))
					return nil
				},
			},
			wantErr:     true,
			expectedErr: errors.New("ForEachEnConcurrent error"),
		},
		{name: "1",
			args: args{
				ctx: context.Background(),
				en:  RangeMust(1, 1000),
				action: func(_ context.Context, i int) error {
					// acc1 += int64(i*i) - this will demonstrate race error
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
			err := ForEachEnConcurrent(tt.args.ctx, tt.args.en, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEachEnConcurrent() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.expectedErr) {
					t.Errorf("ForEachEnConcurrent() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(acc1, tt.want) {
				t.Errorf("ForEachEnConcurrent() = %v, want %v", acc1, tt.want)
			}
		})
	}
}
