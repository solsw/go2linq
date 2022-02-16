//go:build go1.18

package go2linq

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupByTest.cs

func Test_GroupByMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupByMust(en, func(el string) int { return len(el) }))
	if len(grs) != 3 {
		t.Errorf("len(GroupByMust) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupByMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupByMust[0].Key = %v, want %v", gr0.key, 3)
	}
	got0 := NewEnSlice(gr0.values...)
	want0 := NewEnSlice("abc", "def")
	if !SequenceEqualMust(got0, want0) {
		t.Errorf("GroupByMust[0].values = %v, want %v", ToStringDef(got0), ToStringDef(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupByMust[1].Key = %v, want %v", gr1.key, 5)
	}
	got1 := NewEnSlice(gr1.values...)
	want1 := NewEnSlice("hello", "there")
	if !SequenceEqualMust(got1, want1) {
		t.Errorf("GroupByMust[1].values = %v, want %v", ToStringDef(got1), ToStringDef(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupByMust[2].Key = %v, want %v", gr2.key, 4)
	}
	got2 := NewEnSlice(gr2.values...)
	want2 := NewEnSlice("four")
	if !SequenceEqualMust(got2, want2) {
		t.Errorf("GroupByMust[2].values = %v, want %v", ToStringDef(got2), ToStringDef(want2))
	}
}

func Test_GroupBySelMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupBySelMust(en,
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] }),
	)
	if len(grs) != 3 {
		t.Errorf("len(GroupBySelMust) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupBySelMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupBySelMust[0].Key = %v, want %v", gr0.key, 3)
	}
	got0 := NewEnSlice(gr0.values...)
	want0 := NewEnSlice('a', 'd')
	if !SequenceEqualMust(got0, want0) {
		t.Errorf("GroupBySelMust[0].values = %v, want %v", ToStringDef(got0), ToStringDef(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupBySelMust[1].Key = %v, want %v", gr1, 3)
	}
	got1 := NewEnSlice(gr1.values...)
	want1 := NewEnSlice('h', 't')
	if !SequenceEqualMust(got1, want1) {
		t.Errorf("GroupBySelMust[1].values = %v, want %v", ToStringDef(got1), ToStringDef(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupBySelMust[2].Key = %v, want %v", gr2, 3)
	}
	got2 := NewEnSlice(gr2.values...)
	want2 := NewEnSlice('f')
	if !SequenceEqualMust(got2, want2) {
		t.Errorf("GroupBySelMust[2].values = %v, want %v", ToStringDef(got2), ToStringDef(want2))
	}
}

func Test_GroupByResMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupByResMust(en,
		func(el string) int { return len(el) },
		func(el int, en Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", el, strings.Join(ToStrings(en), ";"))
		}))
	got := NewEnSlice(grs...)
	want := NewEnSlice("3:abc;def", "5:hello;there", "4:four")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupByResMust = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

func Test_GroupBySelResMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupBySelResMust(en,
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] },
		func(el int, en Enumerable[rune]) string {
			vv := func() []string {
				var r []string
				enr := en.GetEnumerator()
				for enr.MoveNext() {
					r = append(r, string(enr.Current()))
				}
				return r
			}()
			return fmt.Sprintf("%v:%v", el, strings.Join(vv, ";"))
		}))
	got := NewEnSlice(grs...)
	want := NewEnSlice("3:a;d", "5:h;t", "4:f")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupBySelResMust = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}
