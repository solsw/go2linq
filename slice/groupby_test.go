package slice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGroupBy(t *testing.T) {
	en := []string{"abc", "hello", "def", "there", "four"}
	grs, _ := GroupBy(en, func(el string) int { return len(el) }, nil)
	if len(grs) != 3 {
		t.Errorf("len(GroupBy) = %v, want %v", len(grs), 3)
	}
	lg0 := grs[0].Count()
	if lg0 != 2 {
		t.Errorf("len(GroupBy[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.Key() != 3 {
		t.Errorf("GroupBy[0].Key = %v, want %v", gr0.Key(), 3)
	}
	got0 := gr0.Values()
	want0 := []string{"abc", "def"}
	if !reflect.DeepEqual(got0, want0) {
		t.Errorf("GroupBy[0].values = %v, want %v", got0, want0)
	}

	gr1 := grs[1]
	if gr1.Key() != 5 {
		t.Errorf("GroupBy[1].Key = %v, want %v", gr1.Key(), 5)
	}
	got1 := gr1.Values()
	want1 := []string{"hello", "there"}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("GroupBy[1].values = %v, want %v", got1, want1)
	}

	gr2 := grs[2]
	if gr2.Key() != 4 {
		t.Errorf("GroupBy[2].Key = %v, want %v", gr2.Key(), 4)
	}
	got2 := gr2.Values()
	want2 := []string{"four"}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("GroupBy[2].values = %v, want %v", got2, want2)
	}
}

func TestGroupBySel(t *testing.T) {
	en := []string{"abc", "hello", "def", "there", "four"}
	grs, _ := GroupBySel(en, func(el string) int { return len(el) }, func(el string) rune { return []rune(el)[0] }, nil)
	if len(grs) != 3 {
		t.Errorf("len(GroupBySel) = %v, want %v", len(grs), 3)
	}
	lg0 := grs[0].Count()
	if lg0 != 2 {
		t.Errorf("len(GroupBySel[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.Key() != 3 {
		t.Errorf("GroupBySel[0].Key = %v, want %v", gr0.Key(), 3)
	}
	got0 := gr0.Values()
	want0 := []rune{'a', 'd'}
	if !reflect.DeepEqual(got0, want0) {
		t.Errorf("GroupBySel[0].values = %v, want %v", got0, want0)
	}

	gr1 := grs[1]
	if gr1.Key() != 5 {
		t.Errorf("GroupBySel[1].Key = %v, want %v", gr1, 3)
	}
	got1 := gr1.Values()
	want1 := []rune{'h', 't'}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("GroupBySel[1].values = %v, want %v", got1, want1)
	}

	gr2 := grs[2]
	if gr2.Key() != 4 {
		t.Errorf("GroupBySel[2].Key = %v, want %v", gr2, 3)
	}
	got2 := gr2.Values()
	want2 := []rune{'f'}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("GroupBySel[2].values = %v, want %v", got2, want2)
	}
}

func TestGroupByRes(t *testing.T) {
	ss := []string{"abc", "hello", "def", "there", "four"}
	got, _ := GroupByRes(ss,
		func(s string) int { return len(s) },
		func(i int, ss []string) string {
			return fmt.Sprintf("%v:%v", i, strings.Join(ss, ";"))
		},
		nil)
	want := []string{"3:abc;def", "5:hello;there", "4:four"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupByRes = %v, want %v", got, want)
	}
}

func TestGroupBySelRes(t *testing.T) {
	en := []string{"abc", "hello", "def", "there", "four"}
	got, _ := GroupBySelRes(en,
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
		},
		nil)
	want := []string{"3:a;d", "5:h;t", "4:f"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBySelRes = %v, want %v", got, want)
	}
}
