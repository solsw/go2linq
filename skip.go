package go2linq

// Reimplementing LINQ to Objects: Part 23 - Take/Skip/TakeWhile/SkipWhile
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-part-23-take-skip-takewhile-skipwhile/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skiplast
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

func factorySkip[Source any](source Enumerable[Source], count int) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		i := 1
		return enrFunc[Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					if i > count {
						return true
					}
					i++
				}
				return false
			},
			crrnt: func() Source { return enr.Current() },
			rst:   func() { i = 1; enr.Reset() },
		}
	}
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip)
func Skip[Source any](source Enumerable[Source], count int) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return source, nil
	}
	return OnFactory(factorySkip(source, count)), nil
}

// SkipMust is like Skip but panics in case of error.
func SkipMust[Source any](source Enumerable[Source], count int) Enumerable[Source] {
	r, err := Skip(source, count)
	if err != nil {
		panic(err)
	}
	return r
}

// SkipLast returns a new enumerable collection that contains the elements from 'source'
// with the last 'count' elements of the source collection omitted.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skiplast)
func SkipLast[Source any](source Enumerable[Source], count int) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return source, nil
	}
	sl := ToSliceMust(source)
	return NewEnSlice(sl[:len(sl)-count]...), nil
}

// SkipLastMust is like SkipLast but panics in case of error.
func SkipLastMust[Source any](source Enumerable[Source], count int) Enumerable[Source] {
	r, err := SkipLast(source, count)
	if err != nil {
		panic(err)
	}
	return r
}

func factorySkipWhile[Source any](source Enumerable[Source], predicate func(Source) bool) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		var c Source
		remaining := false
		return enrFunc[Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					c = enr.Current()
					if remaining {
						return true
					}
					if !predicate(c) {
						remaining = true
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { remaining = false; enr.Reset() },
		}
	}
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile)
func SkipWhile[Source any](source Enumerable[Source], predicate func(Source) bool) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return OnFactory(factorySkipWhile(source, predicate)), nil
}

// SkipWhileMust is like SkipWhile but panics in case of error.
func SkipWhileMust[Source any](source Enumerable[Source], predicate func(Source) bool) Enumerable[Source] {
	r, err := SkipWhile(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

func factorySkipWhileIdx[Source any](source Enumerable[Source], predicate func(Source, int) bool) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		remaining := false
		var c Source
		i := -1 // position before the first element
		return enrFunc[Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					c = enr.Current()
					if remaining {
						return true
					}
					i++
					if !predicate(c, i) {
						remaining = true
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { remaining = false; i = -1; enr.Reset() },
		}
	}
}

// SkipWhileIdx bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// The element's index is used in the logic of the predicate function.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile)
func SkipWhileIdx[Source any](source Enumerable[Source], predicate func(Source, int) bool) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return OnFactory(factorySkipWhileIdx(source, predicate)), nil
}

// SkipWhileIdxMust is like SkipWhileIdx but panics in case of error.
func SkipWhileIdxMust[Source any](source Enumerable[Source], predicate func(Source, int) bool) Enumerable[Source] {
	r, err := SkipWhileIdx(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
