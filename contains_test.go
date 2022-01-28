//go:build go1.18

package go2linq

import (
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ContainsTest.cs

func Test_ContainsMust_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
		value  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchNoComparer",
			args: args{
				source: NewEnSlice("foo", "bar", "baz"),
				value:  "BAR",
			},
			want: false,
		},
		{name: "MatchNoComparer",
			args: args{
				source: NewEnSlice("foo", "bar", "baz"),
				value:  strings.ToLower("BAR"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsMust(tt.args.source, tt.args.value)
			if got != tt.want {
				t.Errorf("ContainsMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ContainsEqMust_string(t *testing.T) {
	type args struct {
		source  Enumerable[string]
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
				source:  NewEnSlice("foo", "bar", "baz"),
				value:   "gronk",
				equaler: CaseInsensitiveEqualer,
			},
			want: false,
		},
		{name: "MatchWithCustomComparer",
			args: args{
				source:  NewEnSlice("foo", "bar", "baz"),
				value:   "BAR",
				equaler: CaseInsensitiveEqualer,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsEqMust(tt.args.source, tt.args.value, tt.args.equaler)
			if got != tt.want {
				t.Errorf("ContainsEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ContainsEqMust_int(t *testing.T) {
	type args struct {
		source  Enumerable[int]
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
				source:  NewEnSlice(10, 1, 5, 0),
				value:   2,
				equaler: EqualerFunc[int](func(i1, i2 int) bool { return i1 == 10/i2 }),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsEqMust(tt.args.source, tt.args.value, tt.args.equaler)
			if got != tt.want {
				t.Errorf("ContainsEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
