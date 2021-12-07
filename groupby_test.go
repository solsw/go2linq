//go:build go1.18

package go2linq

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupByTest.cs

func Test_GroupBy(t *testing.T) {
	en := NewOnSliceEn("abc", "hello", "def", "there", "four")
	grs := Slice(GroupByMust(en, func(el string) int { return len(el) }))
	if len(grs) != 3 {
		t.Errorf("len(GroupBy) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupBy[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupBy[0].Key = '%v', want '%v'", gr0.key, 3)
	}
	got0 := NewOnSliceEn(gr0.values...)
	want0 := NewOnSliceEn("abc", "def")
	if !SequenceEqualMust(got0, want0) {
		got0.Reset()
		want0.Reset()
		t.Errorf("GroupBy[0].values = '%v', want '%v'", String(got0), String(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupBy[1].Key = '%v', want '%v'", gr1.key, 5)
	}
	got1 := NewOnSliceEn(gr1.values...)
	want1 := NewOnSliceEn("hello", "there")
	if !SequenceEqualMust(got1, want1) {
		got1.Reset()
		want1.Reset()
		t.Errorf("GroupBy[1].values = '%v', want '%v'", String(got1), String(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupBy[2].Key = '%v', want '%v'", gr2.key, 4)
	}
	got2 := NewOnSliceEn(gr2.values...)
	want2 := NewOnSliceEn("four")
	if !SequenceEqualMust(got2, want2) {
		got2.Reset()
		want2.Reset()
		t.Errorf("GroupBy[2].values = '%v', want '%v'", String(got2), String(want2))
	}
}

func Test_GroupBySel(t *testing.T) {
	en := NewOnSliceEn("abc", "hello", "def", "there", "four")
	grs := Slice(GroupBySelMust(en,
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] }),
	)
	if len(grs) != 3 {
		t.Errorf("len(GroupBySel) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupBySel[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupBySel[0].Key = '%v', want '%v'", gr0.key, 3)
	}
	got0 := NewOnSliceEn(gr0.values...)
	want0 := NewOnSliceEn('a', 'd')
	if !SequenceEqualMust(got0, want0) {
		got0.Reset()
		want0.Reset()
		t.Errorf("GroupBySel[0].values = '%v', want '%v'", String(got0), String(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupBySel[1].Key = '%v', want '%v'", gr1, 3)
	}
	got1 := NewOnSliceEn(gr1.values...)
	want1 := NewOnSliceEn('h', 't')
	if !SequenceEqualMust(got1, want1) {
		got1.Reset()
		want1.Reset()
		t.Errorf("GroupBySel[1].values = '%v', want '%v'", String(got1), String(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupBySel[2].Key = '%v', want '%v'", gr2, 3)
	}
	got2 := NewOnSliceEn(gr2.values...)
	want2 := NewOnSliceEn('f')
	if !SequenceEqualMust(got2, want2) {
		got2.Reset()
		want2.Reset()
		t.Errorf("GroupBySel[2].values = '%v', want '%v'", String(got2), String(want2))
	}
}

func Test_GroupByRes(t *testing.T) {
	en := NewOnSliceEn("abc", "hello", "def", "there", "four")
	grs := Slice(GroupByResMust(en,
		func(el string) int { return len(el) },
		func(el int, en Enumerator[string]) string {
			return fmt.Sprintf("%v:%v", el, strings.Join(Strings(en), ";"))
		}))
	got := NewOnSliceEn(grs...)
	want := NewOnSliceEn("3:abc;def", "5:hello;there", "4:four")
	if !SequenceEqualMust(got, want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupByRes = '%v', want '%v'", String(got), String(want))
	}
}

func Test_GroupBySelRes(t *testing.T) {
	en := NewOnSliceEn("abc", "hello", "def", "there", "four")
	grs := Slice(GroupBySelResMust(en,
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] },
		func(el int, en Enumerator[rune]) string {
			vv := func() []string {
				var r []string
				for en.MoveNext() {
					r = append(r, string(en.Current()))
				}
				return r
			}()
			return fmt.Sprintf("%v:%v", el, strings.Join(vv, ";"))
		}))
	got := NewOnSliceEn(grs...)
	want := NewOnSliceEn("3:a;d", "5:h;t", "4:f")
	if !SequenceEqualMust(got, want) {
		got.Reset()
		want.Reset()
		t.Errorf("GroupBySelRes = '%v', want '%v'", String(got), String(want))
	}
}
