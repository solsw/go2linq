package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [ToLookup] creates a [Lookup] from a sequence according to a specified key selector function.
// [generichelper.DeepEqual] is used to compare keys. 'source' is enumerated immediately.
//
// [ToLookup]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.tolookup
func ToLookup[Source, Key any](source iter.Seq[Source], keySelector func(Source) Key) (*Lookup[Key, Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ToLookupSelEq(source, keySelector, Identity[Source], generichelper.DeepEqual[Key])
}

// [ToLookupEq] creates a [Lookup] from a sequence according to a specified key selector function and a key equaler.
// 'source' is enumerated immediately.
//
// [ToLookupEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.tolookup
func ToLookupEq[Source, Key any](source iter.Seq[Source],
	keySelector func(Source) Key, equal func(Key, Key) bool) (*Lookup[Key, Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	return ToLookupSelEq(source, keySelector, Identity[Source], equal)
}

// [ToLookupSel] creates a [Lookup] from a according to specified key selector and element selector functions.
// 'source' is enumerated immediately.
//
// [ToLookupSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.tolookup
func ToLookupSel[Source, Key, Element any](source iter.Seq[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (*Lookup[Key, Element], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	return ToLookupSelEq(source, keySelector, elementSelector, generichelper.DeepEqual[Key])
}

// [ToLookupSelEq] creates a [Lookup] from a sequence according to
// a specified key selector function, an element selector function and a key equaler.
// 'source' is enumerated immediately.
//
// [ToLookupSelEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.tolookup
func ToLookupSelEq[Source, Key, Element any](source iter.Seq[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, equal func(Key, Key) bool) (*Lookup[Key, Element], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	lk := &Lookup[Key, Element]{groupings: []Grouping[Key, Element]{}, KeyEqual: equal}
	for s := range source {
		k := keySelector(s)
		lk.Add(k, elementSelector(s))
	}
	return lk, nil
}
