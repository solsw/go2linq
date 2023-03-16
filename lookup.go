package go2linq

import (
	"reflect"
	"strings"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq/Lookup.cs
// https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2

// Lookup represents a collection of keys each mapped to one or more values.
// (https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2)
type Lookup[Key, Element any] struct {
	groupings []Grouping[Key, Element]
	// KeyEq is an equaler for groupings' keys
	KeyEq collate.Equaler[Key]
}

func (lk *Lookup[Key, Element]) keyIndex(key Key) int {
	for i, g := range lk.groupings {
		if lk.KeyEq.Equal(g.key, key) {
			return i
		}
	}
	return -1
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

// Count gets the number of key/value collection pairs in the Lookup.
// (https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2.count)
func (lk *Lookup[Key, Element]) Count() int {
	return len(lk.groupings)
}

// ItemSlice returns a slice containing values.
func (lk *Lookup[Key, Element]) ItemSlice(key Key) []Element {
	i := lk.keyIndex(key)
	if i < 0 {
		return []Element{}
	}
	return lk.groupings[i].values
}

// Item gets the collection of values indexed by the specified key.
// (https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2.item)
func (lk *Lookup[Key, Element]) Item(key Key) Enumerable[Element] {
	return NewEnSlice(lk.ItemSlice(key)...)
}

// Contains determines whether a specified key is in the Lookup.
// (https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2.contains)
func (lk *Lookup[Key, Element]) Contains(key Key) bool {
	return lk.keyIndex(key) >= 0
}

// GetEnumerator returns an enumerator that iterates through the Lookup.
// GetEnumerator implements the Enumerable interface.
// (https://learn.microsoft.com/dotnet/api/system.linq.lookup-2.getenumerator)
func (lk *Lookup[Key, Element]) GetEnumerator() Enumerator[Grouping[Key, Element]] {
	return newEnrSlice(lk.groupings...)
}

// Slice returns a slice containing the Lookup's contents.
func (lk *Lookup[Key, Element]) Slice() []Grouping[Key, Element] {
	return lk.groupings
}

// EqualTo determines whether the current Lookup is equal to the specified Lookup.
// Keys equality comparers do not participate in equality verification,
// since non-nil funcs are always not deeply equal.
func (lk *Lookup[Key, Element]) EqualTo(lk2 *Lookup[Key, Element]) bool {
	if lk.Count() != lk2.Count() {
		return false
	}
	for i, g := range lk.groupings {
		g2 := lk2.groupings[i]
		if !lk.KeyEq.Equal(g.key, g2.key) || !reflect.DeepEqual(g.values, g2.values) {
			return false
		}
	}
	return true
}

// String implements the fmt.Stringer interface.
func (lk *Lookup[Key, Element]) String() string {
	var b strings.Builder
	for _, gr := range lk.groupings {
		if b.Len() > 0 {
			b.WriteString("\n")
		}
		b.WriteString(gr.String())
	}
	return b.String()
}
