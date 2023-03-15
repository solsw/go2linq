package go2linq

import (
	"github.com/solsw/collate"
)

// Reimplementing LINQ to Objects: Part 18 – ToLookup
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-18-tolookup/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup

// ToLookup creates a Lookup from an Enumerable according to a specified key selector function.
// collate.DeepEqualer is used to compare keys.
// 'source' is enumerated immediately.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup)
func ToLookup[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) (*Lookup[Key, Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ToLookupSelEq(source, keySelector, Identity[Source], nil)
}

// ToLookupMust is like [ToLookup] but panics in case of error.
func ToLookupMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) *Lookup[Key, Source] {
	r, err := ToLookup(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupEq creates a Lookup from an Enumerable according to a specified key selector function and a key equaler.
// If 'equaler' is nil collate.DeepEqualer is used.
// 'source' is enumerated immediately.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup)
func ToLookupEq[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, equaler collate.Equaler[Key]) (*Lookup[Key, Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Key]{}
	}
	return ToLookupSelEq(source, keySelector, Identity[Source], equaler)
}

// ToLookupEqMust is like [ToLookupEq] but panics in case of error.
func ToLookupEqMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, equaler collate.Equaler[Key]) *Lookup[Key, Source] {
	r, err := ToLookupEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupSel creates a Lookup from an Enumerable according to specified key selector and element selector functions.
// collate.DeepEqualer is used to compare keys.
// 'source' is enumerated immediately.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup)
func ToLookupSel[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (*Lookup[Key, Element], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	return ToLookupSelEq(source, keySelector, elementSelector, nil)
}

// ToLookupSelMust is like [ToLookupSel] but panics in case of error.
func ToLookupSelMust[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) *Lookup[Key, Element] {
	r, err := ToLookupSel(source, keySelector, elementSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupSelEq creates a Lookup from an Enumerable according to a specified key selector function,
// an element selector function and a key equaler.
// If 'equaler' is nil collate.DeepEqualer is used.
// 'source' is enumerated immediately.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup)
func ToLookupSelEq[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler collate.Equaler[Key]) (*Lookup[Key, Element], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Key]{}
	}
	enr := source.GetEnumerator()
	lk := &Lookup[Key, Element]{groupings: []Grouping[Key, Element]{}, KeyEq: equaler}
	for enr.MoveNext() {
		c := enr.Current()
		k := keySelector(c)
		lk.Add(k, elementSelector(c))
	}
	return lk, nil
}

// ToLookupSelEqMust is like [ToLookupSelEq] but panics in case of error.
func ToLookupSelEqMust[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler collate.Equaler[Key]) *Lookup[Key, Element] {
	r, err := ToLookupSelEq(source, keySelector, elementSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
