package slice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGroupByMust(t *testing.T) {
	en := []string{"abc", "hello", "def", "there", "four"}
	grs := GroupByMust(en, func(el string) int { return len(el) })
	if len(grs) != 3 {
		t.Errorf("len(GroupByMust) = %v, want %v", len(grs), 3)
	}
	lg0 := grs[0].Count()
	if lg0 != 2 {
		t.Errorf("len(GroupByMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.Key() != 3 {
		t.Errorf("GroupByMust[0].Key = %v, want %v", gr0.Key(), 3)
	}
	got0 := gr0.Values()
	want0 := []string{"abc", "def"}
	if !reflect.DeepEqual(got0, want0) {
		t.Errorf("GroupByMust[0].values = %v, want %v", got0, want0)
	}

	gr1 := grs[1]
	if gr1.Key() != 5 {
		t.Errorf("GroupByMust[1].Key = %v, want %v", gr1.Key(), 5)
	}
	got1 := gr1.Values()
	want1 := []string{"hello", "there"}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("GroupByMust[1].values = %v, want %v", got1, want1)
	}

	gr2 := grs[2]
	if gr2.Key() != 4 {
		t.Errorf("GroupByMust[2].Key = %v, want %v", gr2.Key(), 4)
	}
	got2 := gr2.Values()
	want2 := []string{"four"}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("GroupByMust[2].values = %v, want %v", got2, want2)
	}
}

func TestGroupBySelMust(t *testing.T) {
	en := []string{"abc", "hello", "def", "there", "four"}
	grs := GroupBySelMust(en, func(el string) int { return len(el) }, func(el string) rune { return []rune(el)[0] })
	if len(grs) != 3 {
		t.Errorf("len(GroupBySelMust) = %v, want %v", len(grs), 3)
	}
	lg0 := grs[0].Count()
	if lg0 != 2 {
		t.Errorf("len(GroupBySelMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.Key() != 3 {
		t.Errorf("GroupBySelMust[0].Key = %v, want %v", gr0.Key(), 3)
	}
	got0 := gr0.Values()
	want0 := []rune{'a', 'd'}
	if !reflect.DeepEqual(got0, want0) {
		t.Errorf("GroupBySelMust[0].values = %v, want %v", got0, want0)
	}

	gr1 := grs[1]
	if gr1.Key() != 5 {
		t.Errorf("GroupBySelMust[1].Key = %v, want %v", gr1, 3)
	}
	got1 := gr1.Values()
	want1 := []rune{'h', 't'}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("GroupBySelMust[1].values = %v, want %v", got1, want1)
	}

	gr2 := grs[2]
	if gr2.Key() != 4 {
		t.Errorf("GroupBySelMust[2].Key = %v, want %v", gr2, 3)
	}
	got2 := gr2.Values()
	want2 := []rune{'f'}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("GroupBySelMust[2].values = %v, want %v", got2, want2)
	}
}

func TestGroupByResMust(t *testing.T) {
	ss := []string{"abc", "hello", "def", "there", "four"}
	got := GroupByResMust(ss,
		func(s string) int { return len(s) },
		func(i int, ss []string) string {
			return fmt.Sprintf("%v:%v", i, strings.Join(ss, ";"))
		})
	want := []string{"3:abc;def", "5:hello;there", "4:four"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupByResMust = %v, want %v", got, want)
	}
}

func TestGroupBySelResMust(t *testing.T) {
	en := []string{"abc", "hello", "def", "there", "four"}
	got := GroupBySelResMust(en,
		func(s string) int { return len(s) },
		func(s string) rune { return []rune(s)[0] },
		func(i int, rr []rune) string {
			vv := func() []string {
				var ss []string
				for _, v := range rr {
					ss = append(ss, string(v))
				}
				return ss
			}()
			return fmt.Sprintf("%v:%v", i, strings.Join(vv, ";"))
		})
	want := []string{"3:a;d", "5:h;t", "4:f"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBySelResMust = %v, want %v", got, want)
	}
}