package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestToLookupMust_string_int(t *testing.T) {
	lk := &go2linq.Lookup[int, string]{KeyEq: go2linq.DeepEqualer[int]{}}
	lk.Add(3, "abc")
	lk.Add(3, "def")
	lk.Add(1, "x")
	lk.Add(1, "y")
	lk.Add(3, "ghi")
	lk.Add(1, "z")
	lk.Add(2, "00")
	type args struct {
		source      []string
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want *go2linq.Lookup[int, string]
	}{
		{name: "EmptySource",
			args: args{
				source:      []string{},
				keySelector: func(s string) int { return len(s) },
			},
			want: &go2linq.Lookup[int, string]{},
		},
		{name: "LookupWithNoComparerOrElementSelector",
			args: args{
				source:      []string{"abc", "def", "x", "y", "ghi", "z", "00"},
				keySelector: func(s string) int { return len(s) },
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLookupMust(tt.args.source, tt.args.keySelector, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLookupMust_string_string(t *testing.T) {
	lk := &go2linq.Lookup[string, string]{KeyEq: go2linq.DeepEqualer[string]{}}
	lk.Add("abc", "abc")
	lk.Add("def", "def")
	lk.Add("ABC", "ABC")
	type args struct {
		source      []string
		keySelector func(string) string
		equaler     go2linq.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want *go2linq.Lookup[string, string]
	}{
		{name: "LookupWithNilComparerButNoElementSelector",
			args: args{
				source:      []string{"abc", "def", "ABC"},
				keySelector: go2linq.Identity[string],
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLookupMust(tt.args.source, tt.args.keySelector, tt.args.equaler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLookupMust(t *testing.T) {
	lk := &go2linq.Lookup[string, string]{KeyEq: go2linq.DeepEqualer[string]{}}
	lk.Add("abc", "abc")
	lk.Add("def", "def")
	lk.Add("abc", "ABC")
	type args struct {
		source      []string
		keySelector func(string) string
		equaler     go2linq.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want *go2linq.Lookup[string, string]
	}{
		{name: "LookupWithComparerButNoElementSelector",
			args: args{
				source:      []string{"abc", "def", "ABC"},
				keySelector: go2linq.Identity[string],
				equaler:     go2linq.CaseInsensitiveEqualer,
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLookupMust(tt.args.source, tt.args.keySelector, tt.args.equaler); !got.EqualTo(tt.want) {
				t.Errorf("ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
