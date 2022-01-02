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

func TestSlice_int(t *testing.T) {
	type args struct {
		en Enumerator[int]
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
				en: NewOnSlice[int](1, 2, 3, 4),
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Slice[int](tt.args.en)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestSliceErr_interface_int(t *testing.T) {
	type args struct {
		en Enumerator[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "CastExceptionOnWrongElementType",
			args: args{
				en: CastMust[any, int](NewOnSlice[any](1.0, 2.0, 3.0, 4.0, "five")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SliceErr[int](tt.args.en)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestStringFmt_int(t *testing.T) {
	type args struct {
		en        Enumerator[int]
		separator string
		leftRim   string
		rightRim  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				en:        NewOnSlice(1, 2, 3),
				separator: "-",
				leftRim:   "*",
				rightRim:  "^",
			},
			want: "*1^-*2^-*3^",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFmt(tt.args.en, tt.args.separator, tt.args.leftRim, tt.args.rightRim); got != tt.want {
				t.Errorf("StringFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

type intStringer int

func (i intStringer) String() string {
	return fmt.Sprintf("%d+%d", i, i*i)
}

func TestStringFmt_intStringer(t *testing.T) {
	type args struct {
		en        Enumerator[intStringer]
		separator string
		leftRim   string
		rightRim  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				en:        NewOnSlice(intStringer(1), intStringer(2), intStringer(3)),
				separator: "-",
			},
			want: "1+1-2+4-3+9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFmt(tt.args.en, tt.args.separator, tt.args.leftRim, tt.args.rightRim); got != tt.want {
				t.Errorf("StringFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrings_int(t *testing.T) {
	type args struct {
		en Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "1",
			args: args{
				en: NewOnSlice(1, 2, 3),
			},
			// want: NewOnSlice("1", "2"),
			want: NewOnSlice("1", "2", "3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStrings(tt.args.en); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ToStrings() = %v, want %v", String(got), String(tt.want))
			}
		})
	}
}

func TestStrings_int(t *testing.T) {
	type args struct {
		en Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",
			args: args{
				en: NewOnSlice(1, 2, 3),
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strings(tt.args.en); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrings_intStringer(t *testing.T) {
	type args struct {
		en Enumerator[intStringer]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",
			args: args{
				en: NewOnSlice(intStringer(1), 2, 3),
			},
			want: []string{"1+1", "2+4", "3+9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strings(tt.args.en); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach_int(t *testing.T) {
	var acc1 int
	type args struct {
		ctx    context.Context
		en     Enumerator[int]
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
				en:  NewOnSlice(1, 2, 3),
			},
			wantErr:     true,
			expectedErr: ErrNilAction,
		},
		{name: "03",
			args: args{
				ctx: context.Background(),
				en:  NewOnSlice(1, 2, 3),
				action: func(_ context.Context, i int) error {
					if i == 2 {
						return errors.New("ForEach error")
					}
					acc1 += i * i
					return nil
				},
			},
			wantErr:     true,
			expectedErr: errors.New("ForEach error"),
		},
		{name: "1",
			args: args{
				ctx: context.Background(),
				en:  NewOnSlice(1, 2, 3),
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
			err := ForEach(tt.args.ctx, tt.args.en, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEach() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.expectedErr) {
					t.Errorf("ForEach() error = '%v', expectedErr '%v'", err, tt.expectedErr)
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
	type args struct {
		ctx    context.Context
		en     Enumerator[int]
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
				en:  NewOnSlice(1, 2, 3),
				action: func(_ context.Context, i int) error {
					if i == 2 {
						return errors.New("ForEachConcurrent error")
					}
					atomic.AddInt64(&acc1, int64(i*i))
					return nil
				},
			},
			wantErr:     true,
			expectedErr: errors.New("ForEachConcurrent error"),
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
			err := ForEachConcurrent(tt.args.ctx, tt.args.en, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEachConcurrent() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.expectedErr) {
					t.Errorf("ForEachConcurrent() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(acc1, tt.want) {
				t.Errorf("ForEachConcurrent() = %v, want %v", acc1, tt.want)
			}
		})
	}
}
