//go:build go1.18

package go2linq

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/JoinTest.cs

func Test_JoinMust_string_rune(t *testing.T) {
	en := NewEnSlice("fs", "sf", "ff", "ss")
	type args struct {
		outer            Enumerable[string]
		inner            Enumerable[string]
		outerKeySelector func(string) rune
		innerKeySelector func(string) rune
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SimpleJoin",
			args: args{
				outer:            NewEnSlice("first", "second", "third"),
				inner:            NewEnSlice("essence", "offer", "eating", "psalm"),
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewEnSlice("first:offer", "second:essence", "second:psalm"),
		},
		{name: "SameEnumerable",
			args: args{
				outer:            en,
				inner:            en,
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewEnSlice("fs:sf", "fs:ff", "sf:fs", "sf:ss", "ff:sf", "ff:ff", "ss:fs", "ss:ss"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinMust(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("JoinMust() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_JoinMust_string(t *testing.T) {
	type args struct {
		outer            Enumerable[string]
		inner            Enumerable[string]
		outerKeySelector func(string) string
		innerKeySelector func(string) string
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "CustomComparer",
			args: args{
				outer: NewEnSlice("ABCxxx", "abcyyy", "defzzz", "ghizzz"),
				inner: NewEnSlice("000abc", "111gHi", "222333"),
				outerKeySelector: func(oel string) string {
					return strings.ToLower(oel[:3])
				},
				innerKeySelector: func(iel string) string {
					return strings.ToLower(iel[3:])
				},
				resultSelector: func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewEnSlice("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinMust(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("JoinMust() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_JoinEqMust_CustomComparer(t *testing.T) {
	outer := NewEnSlice("ABCxxx", "abcyyy", "defzzz", "ghizzz")
	inner := NewEnSlice("000abc", "111gHi", "222333")
	got := JoinEqMust(outer, inner,
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel, iel string) string { return oel + ":" + iel },
		CaseInsensitiveEqualer,
	)
	want := NewEnSlice("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi")
	if !SequenceEqualMust(got, want) {
		t.Errorf("JoinEqMust_CustomComparer = '%v', want '%v'", EnToString(got), EnToString(want))
	}
}

func Test_JoinMust_DifferentSourceTypes(t *testing.T) {
	outer := NewEnSlice(5, 3, 7)
	inner := NewEnSlice("bee", "giraffe", "tiger", "badger", "ox", "cat", "dog")
	got := JoinMust(outer, inner,
		Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iel string) string { return fmt.Sprintf("%d:%s", oel, iel) },
	)
	want := NewEnSlice("5:tiger", "3:bee", "3:cat", "3:dog", "7:giraffe")
	if !SequenceEqualMust(got, want) {
		t.Errorf("JoinMust_DifferentSourceTypes = '%v', want '%v'", EnToString(got), EnToString(want))
	}
}
