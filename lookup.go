//go:build go1.18

package go2linq

import (
	"reflect"
	"strings"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq/Lookup.cs
// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2

// Lookup represents a collection of keys each mapped to one or more values.
// (https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2)
type Lookup[Key, Element any] struct {
	grgr []*Grouping[Key, Element]
	// keyEq is an equaler for grgr's keys
	keyEq Equaler[Key]
}

// newLookupEq creates new empty Lookup with the provided keys equaler
func newLookupEq[Key, Element any](keq Equaler[Key]) *Lookup[Key, Element] {
	return &Lookup[Key, Element]{keyEq: keq}
}

// newLookup creates new empty Lookup using DeepEqualer as key equaler
func newLookup[Key, Element any]() *Lookup[Key, Element] {
	var keq Equaler[Key] = DeepEqualer[Key]{}
	return newLookupEq[Key, Element](keq)
}

func (lk *Lookup[Key, Element]) keyIndex(key Key) int {
	for i, g := range lk.grgr {
		if lk.keyEq.Equal(g.key, key) {
			return i
		}
	}
	return -1
}

// add adds element 'el' with specified 'key' to 'lk'
func (lk *Lookup[Key, Element]) add(key Key, el Element) {
	i := lk.keyIndex(key)
	if i >= 0 {
		lk.grgr[i].values = append(lk.grgr[i].values, el)
	} else {
		gr := Grouping[Key, Element]{key: key, values: []Element{el}}
		lk.grgr = append(lk.grgr, &gr)
	}
}

// Count gets the number of key/value collection pairs in the Lookup.
// (https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2.count)
func (lk *Lookup[Key, Element]) Count() int {
	return len(lk.grgr)
}

// ItemSlice returns a slice containing values.
func (lk *Lookup[Key, Element]) ItemSlice(key Key) []Element {
	i := lk.keyIndex(key)
	if i < 0 {
		return []Element{}
	}
	return lk.grgr[i].values
}

// Item gets the collection of values indexed by the specified key.
// (https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2.item)
func (lk *Lookup[Key, Element]) Item(key Key) Enumerable[Element] {
	return NewEnSlice(lk.ItemSlice(key)...)
}

// Contains determines whether a specified key is in the Lookup.
// (https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2.contains)
func (lk *Lookup[Key, Element]) Contains(key Key) bool {
	return lk.keyIndex(key) >= 0
}

// GetEnumerator returns a generic enumerator that iterates through the Lookup.
// (https://docs.microsoft.com/dotnet/api/system.linq.lookup-2.getenumerator)
func (lk *Lookup[Key, Element]) GetEnumerator() Enumerator[*Grouping[Key, Element]] {
	return newEnrSlice(lk.grgr...)
}

// Slice returns a slice containing the Lookup's contents.
func (lk *Lookup[Key, Element]) Slice() []*Grouping[Key, Element] {
	return lk.grgr
}

// EqualTo determines whether the current Lookup is equal to the specified Lookup.
// Keys equality comparers do not participate in equality verification,
// since non-nil funcs are always not deeply equal.
func (lk *Lookup[Key, Element]) EqualTo(lk2 *Lookup[Key, Element]) bool {
	if lk.Count() != lk2.Count() {
		return false
	}
	for i, g := range lk.grgr {
		g2 := lk2.grgr[i]
		if !lk.keyEq.Equal(g.key, g2.key) || !reflect.DeepEqual(g.values, g2.values) {
			return false
		}
	}
	return true
}

// String implements the fmt.Stringer interface.
func (lk *Lookup[Key, Element]) String() string {
	var b strings.Builder
	for _, gr := range lk.grgr {
		if b.Len() > 0 {
			// b.WriteString(oshelper.NewLine)
			b.WriteString("\n")
		}
		b.WriteString(gr.String())
	}
	return b.String()
}
