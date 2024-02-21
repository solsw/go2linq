package go2linq

import (
	"iter"
)

// [ToMap] creates a [map] from a sequence according to a specified key selector function.
//
// [map]: https://go.dev/ref/spec#Map_types
// [ToMap]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.todictionary
func ToMap[Source any, Key comparable](source iter.Seq[Source], keySelector func(Source) Key) (map[Key]Source, error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	m := make(map[Key]Source)
	for s := range source {
		k := keySelector(s)
		if _, ok := m[k]; ok {
			return nil, ErrDuplicateKeys
		}
		m[k] = s
	}
	return m, nil
}

// [ToMapSel] creates a [map] from a sequence according to specified key selector and element selector functions.
//
// Since Go's map does not support custom equaler to determine equality of the keys,
// LINQ's key comparer is not implemented.
// Similar to the keys equality functionality may be achieved using appropriate key selector.
// Example of custom key selector that mimics case-insensitive equaler for string keys
// is presented in [TestCustomSelector_string_string_int].
//
// [map]: https://go.dev/ref/spec#Map_types
// [ToMapSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.todictionary
func ToMapSel[Source any, Key comparable, Element any](source iter.Seq[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (map[Key]Element, error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	m := make(map[Key]Element)
	for s := range source {
		k := keySelector(s)
		if _, ok := m[k]; ok {
			return nil, ErrDuplicateKeys
		}
		m[k] = elementSelector(s)
	}
	return m, nil
}
