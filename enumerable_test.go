//go:build go1.18

package go2linq

import (
	"context"
	"fmt"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestOnChan_int(t *testing.T) {
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
			got := OnChan(tt.args.chn)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OnChan() failed")
			}
		})
	}
}

type intStringer int

func (i intStringer) String() string {
	return fmt.Sprintf("%d+%d", i, i*i)
}

func TestToStringDef_Stringer(t *testing.T) {
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
			got := ToStringDef(tt.args.en)
			if got != tt.want {
				t.Errorf("ToStringDef_Stringer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToEnString(t *testing.T) {
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
			got := ToEnString(tt.args.en)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ToEnString() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestToStrings(t *testing.T) {
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
			got := ToStrings(tt.args.en)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrings_Stringer(t *testing.T) {
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
			got := ToStrings(tt.args.en)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrings_Stringer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	var acc1 int
	ctx1, cancel := context.WithCancel(context.Background())
	type args struct {
		ctx    context.Context
		en     Enumerable[int]
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
				en:  NewEnSlice(1, 2, 3),
			},
			wantErr:     true,
			expectedErr: ErrNilAction,
		},
		{name: "03",
			args: args{
				ctx: ctx1,
				en:  NewEnSlice(1, 2, 3),
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
				en:  NewEnSlice(1, 2, 3),
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
				en:  NewEnSlice(1, 2, 3),
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
			err := ForEach(tt.args.ctx, tt.args.en, tt.args.action)
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

func TestForEachConcurrent(t *testing.T) {
	var acc1 int64
	canceledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	type args struct {
		ctx    context.Context
		en     Enumerable[int]
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
				en:     NewEnSlice(1, 2, 3),
				action: func(int) error { return nil },
			},
			wantErr:     true,
			expectedErr: context.Canceled,
		},
		{name: "02",
			args: args{
				ctx: context.Background(),
				en:  NewEnSlice(1, 2, 3),
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
				en:  RangeMust(1, 1000),
				action: func(i int) error {
					// acc1 += int64(i*i) <- demonstrates race error
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
			err := ForEachConcurrent(tt.args.ctx, tt.args.en, tt.args.action)
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
