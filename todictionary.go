//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 25 – ToDictionary
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-todictionary/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary

// ToDictionary creates a Dictionary from an Enumerator according to a specified key selector function.
// ToDictionary panics if 'source' or 'keySelector' is nil.
func ToDictionary[Source any, Key comparable](source Enumerator[Source], keySelector func(Source) Key) Dictionary[Key, Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	r := make(Dictionary[Key, Source])
	for source.MoveNext() {
		c := source.Current()
		k := keySelector(c)
		/*
			if k == nil {
			  panic(ErrNilKey)
			}
		*/
		if _, ok := r[k]; ok {
			panic(ErrDuplicateKeys)
		}
		r[k] = c
	}
	return r
}

// ToDictionaryErr is like ToDictionary but returns an error instead of panicking.
func ToDictionaryErr[Source any, Key comparable](source Enumerator[Source], keySelector func(Source) Key) (res Dictionary[Key, Source], err error) {
	defer func() {
		catchErrPanic[Dictionary[Key, Source]](recover(), &res, &err)
	}()
	return ToDictionary(source, keySelector), nil
}

// ToDictionarySel creates a Dictionary from an Enumerator according to specified key selector and element selector functions.
// ToDictionarySel panics if 'source' or 'keySelector' or 'elementSelector' is nil.
//
// Since Dictionary is implemented as a map[Key]Element and since Go's map does not support custom
// equality comparer to determine equality of keys, hence LINQ's key comparer is not implemented.
// Similar (to key comparer) functionality may be achieved using appropriate key selector.
// Example of custom key selector that mimics case-insensitive equality comparer for string keys
// is presented in Test_CustomSelector_string_string_int.
func ToDictionarySel[Source any, Key comparable, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) Dictionary[Key, Element] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil {
		panic(ErrNilSelector)
	}
	r := make(Dictionary[Key, Element])
	for source.MoveNext() {
		c := source.Current()
		k := keySelector(c)
		/*
			if k == nil {
			  panic(ErrNilKey)
			}
		*/
		if _, ok := r[k]; ok {
			panic(ErrDuplicateKeys)
		}
		r[k] = elementSelector(c)
	}
	return r
}

// ToDictionarySelErr is like ToDictionarySel but returns an error instead of panicking.
func ToDictionarySelErr[Source any, Key comparable, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (res Dictionary[Key, Element], err error) {
	defer func() {
		catchErrPanic[Dictionary[Key, Element]](recover(), &res, &err)
	}()
	return ToDictionarySel(source, keySelector, elementSelector), nil
}
