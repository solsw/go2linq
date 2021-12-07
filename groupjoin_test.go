//go:build go1.18

package go2linq

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupJoinTest.cs

func Test_GroupJoinMust_SimpleGroupJoin(t *testing.T) {
	outer := NewOnSliceEn("first", "second", "third")
	inner := NewOnSliceEn("essence", "offer", "eating", "psalm")
	got := GroupJoinMust(outer, inner,
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels Enumerator[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(Strings(iels), ";"))
		})
	want := NewOnSliceEn("first:offer", "second:essence;psalm", "third:")
	if !SequenceEqualMust(got, want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupJoinMust_SimpleGroupJoin = '%v', want '%v'", String(got), String(want))
	}
}

func Test_GroupJoinSelfMust_SameEnumerable(t *testing.T) {
	outer := NewOnSliceEn("fs", "sf", "ff", "ss")
	inner := outer
	got := Slice(GroupJoinSelfMust(outer, inner,
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels Enumerator[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(Strings(iels), ";"))
		}))
	want := []string{"fs:sf;ff", "sf:fs;ss", "ff:sf;ff", "ss:fs;ss"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinSelfMust_SameEnumerable = %v, want %v", got, want)
	}
}

func Test_GroupJoinEqMust_CustomComparer(t *testing.T) {
	outer := NewOnSliceEn("ABCxxx", "abcyyy", "defzzz", "ghizzz")
	inner := NewOnSliceEn("000abc", "111gHi", "222333", "333AbC")
	got := GroupJoinEqMust(outer, inner,
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel string, iels Enumerator[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(Strings(iels), ";"))
		},
		CaseInsensitiveEqualer)
	want := NewOnSliceEn("ABCxxx:000abc;333AbC", "abcyyy:000abc;333AbC", "defzzz:", "ghizzz:111gHi")
	if !SequenceEqualMust(got, want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupJoinEqMust_CustomComparer = '%v', want '%v'", String(got), String(want))
	}
}

func Test_GroupJoinMust_DifferentSourceTypes(t *testing.T) {
	outer := NewOnSliceEn(5, 3, 7, 4)
	inner := NewOnSliceEn("bee", "giraffe", "tiger", "badger", "ox", "cat", "dog")
	got := GroupJoinMust(outer, inner, Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iels Enumerator[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(Strings(iels), ";"))
		},
	)
	want := NewOnSliceEn("5:tiger", "3:bee;cat;dog", "7:giraffe", "4:")
	if !SequenceEqualMust(got, want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupJoinMust_DifferentSourceTypes = '%v', want '%v'", String(got), String(want))
	}
}
