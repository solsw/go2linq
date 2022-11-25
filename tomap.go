package go2linq

// Reimplementing LINQ to Objects: Part 25 – ToDictionary
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-todictionary/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary

// ToMap creates a map from an Enumerable according to a specified key selector function.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary)
func ToMap[Source any, Key comparable](source Enumerable[Source], keySelector func(Source) Key) (map[Key]Source, error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	enr := source.GetEnumerator()
	r := make(map[Key]Source)
	for enr.MoveNext() {
		c := enr.Current()
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

// ToMapMust is like ToMap but panics in case of error.
func ToMapMust[Source any, Key comparable](source Enumerable[Source], keySelector func(Source) Key) map[Key]Source {
	r, err := ToMap(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ToMapSel creates a map from an Enumerable according to specified key selector and element selector functions.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary)
//
// Since Go's map does not support custom equaler to determine equality of the keys,
// LINQ's key comparer is not implemented.
// Similar to the keys equality functionality may be achieved using appropriate key selector.
// Example of custom key selector that mimics case-insensitive equaler for string keys
// is presented in Test_CustomSelector_string_string_int.
func ToMapSel[Source any, Key comparable, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (map[Key]Element, error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	enr := source.GetEnumerator()
	r := make(map[Key]Element)
	for enr.MoveNext() {
		c := enr.Current()
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

// ToMapSelMust is like ToMapSel but panics in case of error.
func ToMapSelMust[Source any, Key comparable, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) map[Key]Element {
	r, err := ToMapSel(source, keySelector, elementSelector)
	if err != nil {
		panic(err)
	}
	return r
}
