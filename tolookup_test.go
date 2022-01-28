//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToLookupTest.cs

func Test_ToLookupMust_string_int(t *testing.T) {
	lk1 := newLookup[int, string]()
	lk1.add(3, "abc")
	lk1.add(3, "def")
	lk1.add(1, "x")
	lk1.add(1, "y")
	lk1.add(3, "ghi")
	lk1.add(1, "z")
	lk1.add(2, "00")
	type args struct {
		source      Enumerable[string]
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want *Lookup[int, string]
	}{
		{name: "LookupWithNoComparerOrElementSelector",
			args: args{
				source:      NewEnSlice("abc", "def", "x", "y", "ghi", "z", "00"),
				keySelector: func(s string) int { return len(s) },
			},
			want: lk1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupMust(tt.args.source, tt.args.keySelector)
			if !got.Equal(tt.want) {
				t.Errorf("ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToLookupMust_string_string(t *testing.T) {
	lk2 := newLookup[string, string]()
	lk2.add("abc", "abc")
	lk2.add("def", "def")
	lk2.add("ABC", "ABC")
	type args struct {
		source      Enumerable[string]
		keySelector func(string) string
	}
	tests := []struct {
		name string
		args args
		want *Lookup[string, string]
	}{
		{name: "LookupWithNilComparerButNoElementSelector",
			args: args{
				source:      NewEnSlice("abc", "def", "ABC"),
				keySelector: Identity[string],
			},
			want: lk2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupMust(tt.args.source, tt.args.keySelector)
			if !got.Equal(tt.want) {
				t.Errorf("ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToLookupSelMust(t *testing.T) {
	lk := newLookup[int, string]()
	lk.add(3, "a")
	lk.add(3, "d")
	lk.add(1, "x")
	lk.add(1, "y")
	lk.add(3, "g")
	lk.add(1, "z")
	lk.add(2, "0")
	type args struct {
		source          Enumerable[string]
		keySelector     func(string) int
		elementSelector func(string) string
	}
	tests := []struct {
		name string
		args args
		want *Lookup[int, string]
	}{
		{name: "LookupWithElementSelectorButNoComparer",
			args: args{
				source:          NewEnSlice("abc", "def", "x", "y", "ghi", "z", "00"),
				keySelector:     func(s string) int { return len(s) },
				elementSelector: func(s string) string { return string(s[0]) },
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupSelMust(tt.args.source, tt.args.keySelector, tt.args.elementSelector)
			if !got.Equal(tt.want) {
				t.Errorf("ToLookupSelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToLookupEqMust(t *testing.T) {
	lk := newLookup[string, string]()
	lk.add("abc", "abc")
	lk.add("def", "def")
	lk.add("abc", "ABC")
	type args struct {
		source      Enumerable[string]
		keySelector func(string) string
		equaler     Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want *Lookup[string, string]
	}{
		{name: "LookupWithComparerButNoElementSelector",
			args: args{
				source:      NewEnSlice("abc", "def", "ABC"),
				keySelector: Identity[string],
				equaler:     CaseInsensitiveEqualer,
			},
			want: lk},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupEqMust(tt.args.source, tt.args.keySelector, tt.args.equaler)
			if !got.Equal(tt.want) {
				t.Errorf("ToLookupEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
