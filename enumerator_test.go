//go:build go1.18

package go2linq

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringFmt_int(t *testing.T) {
	type args struct {
		en Enumerator[int]
		separator string
		leftRim string
		rightRim string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ name: "1",
			args: args{
				en: NewOnSlice(1, 2, 3),
				separator: "-",
				leftRim: "*",
				rightRim: "^",
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
		en Enumerator[intStringer]
		separator string
		leftRim string
		rightRim string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ name: "1",
			args: args{
				en: NewOnSlice(intStringer(1), intStringer(2), intStringer(3)),
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
		{ name: "1",
			args: args{
				en: NewOnSlice(1, 2, 3),
			},
			// want: NewOnSlice("1", "2"),
			want: NewOnSlice("1", "2", "3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {			
			if got := ToStrings(tt.args.en); !SequenceEqual(got, tt.want) {
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
		{ name: "1",
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
		{ name: "1",
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
