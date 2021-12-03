//go:build go1.18

package go2linq

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/JoinTest.cs

func Test_Join_string_rune(t *testing.T) {
	type args struct {
		outer            Enumerator[string]
		inner            Enumerator[string]
		outerKeySelector func(string) rune
		innerKeySelector func(string) rune
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "SimpleJoin",
			args: args{
				outer:            NewOnSlice("first", "second", "third"),
				inner:            NewOnSlice("essence", "offer", "eating", "psalm"),
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewOnSlice("first:offer", "second:essence", "second:psalm"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Join(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Join() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Join_string(t *testing.T) {
	type args struct {
		outer            Enumerator[string]
		inner            Enumerator[string]
		outerKeySelector func(string) string
		innerKeySelector func(string) string
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "CustomComparer",
			args: args{
				outer: NewOnSlice("ABCxxx", "abcyyy", "defzzz", "ghizzz"),
				inner: NewOnSlice("000abc", "111gHi", "222333"),
				outerKeySelector: func(oel string) string {
					return strings.ToLower(oel[:3])
				},
				innerKeySelector: func(iel string) string {
					return strings.ToLower(iel[3:])
				},
				resultSelector: func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewOnSlice("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Join(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Join() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_JoinEq_CustomComparer(t *testing.T) {
	outer := NewOnSlice("ABCxxx", "abcyyy", "defzzz", "ghizzz")
	inner := NewOnSlice("000abc", "111gHi", "222333")
	got, _ := JoinEq(outer, inner,
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel, iel string) string { return oel + ":" + iel },
		CaseInsensitiveEqualer,
	)
	want := NewOnSlice("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi")
	if !SequenceEqualMust(got, want) {
		got.Reset()
		want.Reset()
		t.Errorf("JoinEq_CustomComparer = '%v', want '%v'", String(got), String(want))
	}
}

func Test_Join_DifferentSourceTypes(t *testing.T) {
	outer := NewOnSlice(5, 3, 7)
	inner := NewOnSlice("bee", "giraffe", "tiger", "badger", "ox", "cat", "dog")
	got, _ := Join(outer, inner,
		Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iel string) string { return fmt.Sprintf("%d:%s", oel, iel) },
	)
	want := NewOnSlice("5:tiger", "3:bee", "3:cat", "3:dog", "7:giraffe")
	if !SequenceEqualMust(got, want) {
		got.Reset()
		want.Reset()
		t.Errorf("Join_DifferentSourceTypes = '%v', want '%v'", String(got), String(want))
	}
}

func Test_JoinSelf(t *testing.T) {
	en := NewOnSlice("fs", "sf", "ff", "ss")
	type args struct {
		outer            Enumerator[string]
		inner            Enumerator[string]
		outerKeySelector func(string) rune
		innerKeySelector func(string) rune
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "SameEnumerable",
			args: args{
				outer:            en,
				inner:            en,
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewOnSlice("fs:sf", "fs:ff", "sf:fs", "sf:ss", "ff:sf", "ff:ff", "ss:fs", "ss:ss"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := JoinSelf(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("JoinSelf() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
