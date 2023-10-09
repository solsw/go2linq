package go2linq

import (
	"github.com/solsw/generichelper"
)

// EnMap is an [Enumerable] implementation based on a [map].
//
// [map]: https://go.dev/ref/spec#Map_types
type EnMap[Key comparable, Element any] EnSlice[generichelper.Tuple2[Key, Element]]

// NewEnMap creates a new [EnMap] with a specified map as contents.
func NewEnMap[Key comparable, Element any](m map[Key]Element) *EnMap[Key, Element] {
	sl := make([]generichelper.Tuple2[Key, Element], 0, len(m))
	for k, e := range m {
		sl = append(sl, generichelper.Tuple2[Key, Element]{Item1: k, Item2: e})
	}
	en := EnMap[Key, Element](EnSlice[generichelper.Tuple2[Key, Element]](sl))
	return &en
}

// GetEnumerator implements the [Enumerable] interface.
func (en *EnMap[Key, Element]) GetEnumerator() Enumerator[generichelper.Tuple2[Key, Element]] {
	return (*EnSlice[generichelper.Tuple2[Key, Element]])(en).GetEnumerator()
}

// Count implements the [Counter] interface.
func (en *EnMap[Key, Element]) Count() int {
	return len(*en)
}
