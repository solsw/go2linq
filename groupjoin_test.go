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
	outer := NewEnSlice("first", "second", "third")
	inner := NewEnSlice("essence", "offer", "eating", "psalm")
	got := GroupJoinMust(outer, inner,
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
		})
	want := NewEnSlice("first:offer", "second:essence;psalm", "third:")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupJoinMust_SimpleGroupJoin = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

func Test_GroupJoinMust_SameEnumerable(t *testing.T) {
	outer := NewEnSlice("fs", "sf", "ff", "ss")
	inner := outer
	got := ToSliceMust(GroupJoinMust(outer, inner,
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
		}))
	want := []string{"fs:sf;ff", "sf:fs;ss", "ff:sf;ff", "ss:fs;ss"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinMust_SameEnumerable = %v, want %v", got, want)
	}
}

func Test_GroupJoinEqMust_CustomComparer(t *testing.T) {
	outer := NewEnSlice("ABCxxx", "abcyyy", "defzzz", "ghizzz")
	inner := NewEnSlice("000abc", "111gHi", "222333", "333AbC")
	got := GroupJoinEqMust(outer, inner,
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel string, iels Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
		},
		CaseInsensitiveEqualer)
	want := NewEnSlice("ABCxxx:000abc;333AbC", "abcyyy:000abc;333AbC", "defzzz:", "ghizzz:111gHi")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupJoinEqMust_CustomComparer = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

func Test_GroupJoinMust_DifferentSourceTypes(t *testing.T) {
	outer := NewEnSlice(5, 3, 7, 4)
	inner := NewEnSlice("bee", "giraffe", "tiger", "badger", "ox", "cat", "dog")
	got := GroupJoinMust(outer, inner, Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iels Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
		},
	)
	want := NewEnSlice("5:tiger", "3:bee;cat;dog", "7:giraffe", "4:")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupJoinMust_DifferentSourceTypes = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}
