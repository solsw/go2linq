//go:build go1.18

package go2linq

import (
	"sort"
	"sync"
)

// Reimplementing LINQ to Objects: Part 26a – IOrderedEnumerable
// https://codeblog.jonskeet.uk/2011/01/04/reimplementing-linq-to-objects-part-26a-iorderedenumerable/
// https://docs.microsoft.com/dotnet/api/system.linq.iorderedenumerable-1

// OrderedEnumerable represents a sorted sequence.
type OrderedEnumerable[Element any] struct {
	en Enumerable[Element]
	ls Lesser[Element]
}

// GetEnumerator converts OrderedEnumerable to sorted sequence using sort.SliceStable for sorting.
func (oe *OrderedEnumerable[Element]) GetEnumerator() Enumerator[Element] {
	var once sync.Once
	var elel []Element
	idx := 0
	return enrFunc[Element]{
		mvNxt: func() bool {
			once.Do(func() {
				elel = EnToSlice(oe.en)
				sort.SliceStable(elel, func(i, j int) bool {
					return oe.ls.Less(elel[i], elel[j])
				})
			})
			if idx >= len(elel) {
				return false
			}
			idx++
			return true
		},
		crrnt: func() Element { return elel[idx-1] },
		rst:   func() { idx = 0 },
	}
}
