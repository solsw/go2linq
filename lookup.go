package go2linq

import (
	"iter"
	"reflect"
	"slices"
)

// [Lookup] represents a collection of keys each mapped to one or more values.
//
// [Lookup]: https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2
type Lookup[Key, Element any] struct {
	groupings []Grouping[Key, Element]
	// KeyEqual is an equaler for groupings' keys.
	KeyEqual func(Key, Key) bool
}

func (lk *Lookup[Key, Element]) keyIndex(key Key) int {
	return slices.IndexFunc(lk.groupings,
		func(g Grouping[Key, Element]) bool { return lk.KeyEqual(g.key, key) })
}

// Add adds element 'el' with specified 'key' to 'lk'
func (lk *Lookup[Key, Element]) Add(key Key, el Element) {
	i := lk.keyIndex(key)
	if i >= 0 {
		lk.groupings[i].values = append(lk.groupings[i].values, el)
	} else {
		gr := Grouping[Key, Element]{key: key, values: []Element{el}}
		lk.groupings = append(lk.groupings, gr)
	}
}

// [Count] gets the number of key/value collection pairs in the Lookup.
//
// [Count]: https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2.count
func (lk *Lookup[Key, Element]) Count() int {
	return len(lk.groupings)
}

// [Contains] determines whether a specified key is in the [Lookup].
//
// [Contains]: https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2.contains
func (lk *Lookup[Key, Element]) Contains(key Key) bool {
	return lk.keyIndex(key) >= 0
}

// itemSlice returns a slice containing values.
func (lk *Lookup[Key, Element]) itemSlice(key Key) []Element {
	i := lk.keyIndex(key)
	if i < 0 {
		return []Element{}
	}
	return lk.groupings[i].values
}

// [Item] gets the collection of values indexed by a specified key.
//
// [Item]: https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2.item
func (lk *Lookup[Key, Element]) Item(key Key) iter.Seq[Element] {
	return SliceAll(lk.itemSlice(key))
}

// EqualTo determines whether the current Lookup is equal to a specified Lookup.
// Keys equality comparers do not participate in equality verification,
// since non-nil funcs are always not deeply equal.
func (lk *Lookup[Key, Element]) EqualTo(lk2 *Lookup[Key, Element]) bool {
	if lk.Count() != lk2.Count() {
		return false
	}
	for i, g := range lk.groupings {
		g2 := lk2.groupings[i]
		if !lk.KeyEqual(g.key, g2.key) || !reflect.DeepEqual(g.values, g2.values) {
			return false
		}
	}
	return true
}

// [ApplyResultSelector] applies a transform function to each key and its associated values and returns the results.
//
// [ApplyResultSelector]: https://learn.microsoft.com/dotnet/api/system.linq.lookup-2.applyresultselector
func ApplyResultSelector[Key, Element, Result any](lookup *Lookup[Key, Element],
	resultSelector func(Key, iter.Seq[Element]) Result) (iter.Seq[Result], error) {
	if lookup == nil {
		return nil, ErrNilSource
	}
	if resultSelector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			for _, g := range lookup.groupings {
				if !yield(resultSelector(g.Key(), g.Values())) {
					return
				}
			}
		},
		nil
}
