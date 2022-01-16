//go:build go1.18

package go2linq

import (
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ContainsTest.cs

func Test_Contains_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		value  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchNoComparer",
			args: args{
				source: NewOnSlice("foo", "bar", "baz"),
				value:  "BAR",
			},
			want: false,
		},
		{name: "MatchNoComparer",
			args: args{
				source: NewOnSlice("foo", "bar", "baz"),
				value:  strings.ToLower("BAR"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Contains(tt.args.source, tt.args.value); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ContainsEq_string(t *testing.T) {
	type args struct {
		source  Enumerator[string]
		value   string
		equaler Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchWithCustomComparer",
			args: args{
				source:  NewOnSlice("foo", "bar", "baz"),
				value:   "gronk",
				equaler: CaseInsensitiveEqualer,
			},
			want: false,
		},
		{name: "MatchWithCustomComparer",
			args: args{
				source:  NewOnSlice("foo", "bar", "baz"),
				value:   "BAR",
				equaler: CaseInsensitiveEqualer,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ContainsEq(tt.args.source, tt.args.value, tt.args.equaler); got != tt.want {
				t.Errorf("ContainsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ContainsEq_int(t *testing.T) {
	type args struct {
		source  Enumerator[int]
		value   int
		equaler Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ImmediateReturnWhenMatchIsFound",
			args: args{
				source:  NewOnSlice(10, 1, 5, 0),
				value:   2,
				equaler: EqualerFunc[int](func(i1, i2 int) bool { return i1 == 10/i2 }),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ContainsEq(tt.args.source, tt.args.value, tt.args.equaler); got != tt.want {
				t.Errorf("ContainsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
