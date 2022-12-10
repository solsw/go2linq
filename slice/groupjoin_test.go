package slice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestGroupJoinMust_SimpleGroupJoin(t *testing.T) {
	outer := []string{"first", "second", "third"}
	inner := []string{"essence", "offer", "eating", "psalm"}
	got := GroupJoinMust(outer, inner,
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels []string) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(iels, ";"))
		},
		nil)
	want := []string{"first:offer", "second:essence;psalm", "third:"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinMust_SimpleGroupJoin = %v, want %v", got, want)
	}
}

func TestGroupJoinMust_SameEnumerable(t *testing.T) {
	outer := []string{"fs", "sf", "ff", "ss"}
	inner := outer
	got := GroupJoinMust(outer, inner,
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels []string) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(iels, ";"))
		},
		nil)
	want := []string{"fs:sf;ff", "sf:fs;ss", "ff:sf;ff", "ss:fs;ss"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinMust_SameEnumerable = %v, want %v", got, want)
	}
}

func TestGroupJoinMust_CustomComparer(t *testing.T) {
	outer := []string{"ABCxxx", "abcyyy", "defzzz", "ghizzz"}
	inner := []string{"000abc", "111gHi", "222333", "333AbC"}
	got := GroupJoinMust(outer, inner,
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel string, iels []string) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(iels, ";"))
		},
		go2linq.CaseInsensitiveEqualer)
	want := []string{"ABCxxx:000abc;333AbC", "abcyyy:000abc;333AbC", "defzzz:", "ghizzz:111gHi"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinMust_CustomComparer = %v, want %v", got, want)
	}
}

func TestGroupJoinMust_DifferentSourceTypes(t *testing.T) {
	outer := []int{5, 3, 7, 4}
	inner := []string{"bee", "giraffe", "tiger", "badger", "ox", "cat", "dog"}
	got := GroupJoinMust(outer, inner, go2linq.Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iels []string) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(iels, ";"))
		},
		nil)
	want := []string{"5:tiger", "3:bee;cat;dog", "7:giraffe", "4:"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinMust_DifferentSourceTypes = %v, want %v", got, want)
	}
}
