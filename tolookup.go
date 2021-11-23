//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 18 – ToLookup
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-18-tolookup/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolookup

// ToLookup creates a Lookup from an Enumerator according to a specified key selector function.
// reflect.DeepEqual is used to compare keys. 'source' is enumerated immediately.
// ToLookup panics if 'source' or 'keySelector' is nil.
func ToLookup[Source, Key any](source Enumerator[Source], keySelector func(Source) Key) *Lookup[Key, Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	return ToLookupEq(source, keySelector, nil)
}

// ToLookupErr is like ToLookup but returns an error instead of panicking.
func ToLookupErr[Source, Key any](source Enumerator[Source], keySelector func(Source) Key) (res *Lookup[Key, Source], err error) {
	defer func() {
		catchErrPanic[*Lookup[Key, Source]](recover(), &res, &err)
	}()
	return ToLookup(source, keySelector), nil
}

// ToLookupEq creates a Lookup from an Enumerator according to a specified key selector function and key equality comparer.
// If 'eq' is nil reflect.DeepEqual is used. 'source' is enumerated immediately.
// ToLookupEq panics if 'source' or 'keySelector' is nil.
func ToLookupEq[Source, Key any](source Enumerator[Source], keySelector func(Source) Key, eq Equaler[Key]) *Lookup[Key, Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if eq == nil {
		eq = EqualerFunc[Key](DeepEqual[Key])
	}
	lk := newLookupEq[Key, Source](eq)
	for source.MoveNext() {
		c := source.Current()
		k := keySelector(c)
		lk.add(k, c)
	}
	return lk
}

// ToLookupEqErr is like ToLookupEq but returns an error instead of panicking.
func ToLookupEqErr[Source, Key any](source Enumerator[Source], keySelector func(Source) Key, eq Equaler[Key]) (res *Lookup[Key, Source], err error) {
	defer func() {
		catchErrPanic[*Lookup[Key, Source]](recover(), &res, &err)
	}()
	return ToLookupEq(source, keySelector, eq), nil
}

// ToLookupSel creates a Lookup from an Enumerator according to specified key selector and element selector functions.
// reflect.DeepEqual is used to compare keys. 'source' is enumerated immediately.
// ToLookupSel panics if 'source' or 'keySelector' or 'elementSelector' is nil.
func ToLookupSel[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) *Lookup[Key, Element] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil {
		panic(ErrNilSelector)
	}
	return ToLookupSelEq(source, keySelector, elementSelector, nil)
}

// ToLookupSelEqErr is like ToLookupSel but returns an error instead of panicking.
func ToLookupSelErr[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (res *Lookup[Key, Element], err error) {
	defer func() {
		catchErrPanic[*Lookup[Key, Element]](recover(), &res, &err)
	}()
	return ToLookupSel(source, keySelector, elementSelector), nil
}

// ToLookupSelEq creates a Lookup from an Enumerator according to a specified key selector function,
// an element selector function and key equality comparer.
// If 'eq' is nil reflect.DeepEqual is used. 'source' is enumerated immediately.
// ToLookupSelEq panics if 'source' or 'keySelector' or 'elementSelector' is nil.
func ToLookupSelEq[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, eq Equaler[Key]) *Lookup[Key, Element] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil {
		panic(ErrNilSelector)
	}
	if eq == nil {
		eq = EqualerFunc[Key](DeepEqual[Key])
	}
	lk := newLookupEq[Key, Element](eq)
	for source.MoveNext() {
		c := source.Current()
		k := keySelector(c)
		lk.add(k, elementSelector(c))
	}
	return lk
}

// ToLookupSelEqErr is like ToLookupSelEq but returns an error instead of panicking.
func ToLookupSelEqErr[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, eq Equaler[Key]) (res *Lookup[Key, Element], err error) {
	defer func() {
		catchErrPanic[*Lookup[Key, Element]](recover(), &res, &err)
	}()
	return ToLookupSelEq(source, keySelector, elementSelector, eq), nil
}
