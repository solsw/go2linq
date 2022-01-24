//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 18 – ToLookup
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-18-tolookup/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup

// ToLookup creates a Lookup from an Enumerable according to a specified key selector function.
// DeepEqual is used to compare keys. 'source' is enumerated immediately.
func ToLookup[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) (*Lookup[Key, Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ToLookupEq(source, keySelector, nil)
}

// ToLookupMust is like ToLookup but panics in case of error.
func ToLookupMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) *Lookup[Key, Source] {
	r, err := ToLookup(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupEq creates a Lookup from an Enumerable according to a specified key selector function and a key equaler.
// If 'equaler' is nil DeepEqual is used. 'source' is enumerated immediately.
func ToLookupEq[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, equaler Equaler[Key]) (*Lookup[Key, Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = DeepEqual[Key]{}
	}
	enr := source.GetEnumerator()
	lk := newLookupEq[Key, Source](equaler)
	for enr.MoveNext() {
		c := enr.Current()
		k := keySelector(c)
		lk.add(k, c)
	}
	return lk, nil
}

// ToLookupEqMust is like ToLookupEq but panics in case of error.
func ToLookupEqMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, equaler Equaler[Key]) *Lookup[Key, Source] {
	r, err := ToLookupEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupSel creates a Lookup from an Enumerable according to specified key selector and element selector functions.
// DeepEqual is used to compare keys. 'source' is enumerated immediately.
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

// ToLookupSelMust is like ToLookupSel but panics in case of error.
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
// If 'equaler' is nil DeepEqual is used. 'source' is enumerated immediately.
func ToLookupSelEq[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler Equaler[Key]) (*Lookup[Key, Element], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = DeepEqual[Key]{}
	}
	enr := source.GetEnumerator()
	lk := newLookupEq[Key, Element](equaler)
	for enr.MoveNext() {
		c := enr.Current()
		k := keySelector(c)
		lk.add(k, elementSelector(c))
	}
	return lk, nil
}

// ToLookupSelEqMust is like ToLookupSelEq but panics in case of error.
func ToLookupSelEqMust[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler Equaler[Key]) *Lookup[Key, Element] {
	r, err := ToLookupSelEq(source, keySelector, elementSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
