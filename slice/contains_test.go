package slice

import (
	"strings"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestContains(t *testing.T) {
	type args struct {
		source []string
		value  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NilSource",
			args: args{
				source: nil,
				value:  "BAR",
			},
			want: false,
		},
		{name: "EmptySource",
			args: args{
				source: []string{},
				value:  "BAR",
			},
			want: false,
		},
		{name: "NoMatch",
			args: args{
				source: []string{"foo", "bar", "baz"},
				value:  "BAR",
			},
			want: false,
		},
		{name: "Match",
			args: args{
				source: []string{"foo", "bar", "baz"},
				value:  strings.ToLower("BAR"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.source, tt.args.value); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEq(t *testing.T) {
	type args struct {
		source  []string
		value   string
		equaler go2linq.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchWithCustomComparer",
			args: args{
				source:  []string{"foo", "bar", "baz"},
				value:   "gronk",
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: false,
		},
		{name: "MatchWithCustomComparer",
			args: args{
				source:  []string{"foo", "bar", "baz"},
				value:   "BAR",
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsEq(tt.args.source, tt.args.value, tt.args.equaler); got != tt.want {
				t.Errorf("ContainsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
