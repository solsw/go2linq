//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 25 – ToDictionary
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-todictionary/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary

// ToDictionary creates a Dictionary from an Enumerator according to a specified key selector function.
func ToDictionary[Source any, Key comparable](source Enumerator[Source], keySelector func(Source) Key) (Dictionary[Key, Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	r := make(Dictionary[Key, Source])
	for source.MoveNext() {
		c := source.Current()
		k := keySelector(c)
		// invalid operation: cannot compare k == nil (mismatched types Key and untyped nil)
		// if k == nil {
		// 	return nil, ErrNilKey
		// }
		if _, ok := r[k]; ok {
			return nil, ErrDuplicateKeys
		}
		r[k] = c
	}
	return r, nil
}

// ToDictionaryMust is like ToDictionary but panics in case of error.
func ToDictionaryMust[Source any, Key comparable](source Enumerator[Source], keySelector func(Source) Key) Dictionary[Key, Source] {
	r, err := ToDictionary(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ToDictionarySel creates a Dictionary from an Enumerator according to specified key selector and element selector functions.
//
// Since Dictionary is implemented as a map[Key]Element and since Go's map does not support custom
// equality comparer to determine equality of keys, hence LINQ's key comparer is not implemented.
// Similar (to key comparer) functionality may be achieved using appropriate key selector.
// Example of custom key selector that mimics case-insensitive equality comparer for string keys
// is presented in Test_CustomSelector_string_string_int.
func ToDictionarySel[Source any, Key comparable, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (Dictionary[Key, Element], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	r := make(Dictionary[Key, Element])
	for source.MoveNext() {
		c := source.Current()
		k := keySelector(c)
		// invalid operation: cannot compare k == nil (mismatched types Key and untyped nil)
		// if k == nil {
		//   panic(ErrNilKey)
		// }
		if _, ok := r[k]; ok {
			return nil, ErrDuplicateKeys
		}
		r[k] = elementSelector(c)
	}
	return r, nil
}

// ToDictionarySelMust is like ToDictionarySel but panics in case of error.
func ToDictionarySelMust[Source any, Key comparable, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) Dictionary[Key, Element] {
	r, err := ToDictionarySel(source, keySelector, elementSelector)
	if err != nil {
		panic(err)
	}
	return r
}
