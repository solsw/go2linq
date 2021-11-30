//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 23 - Take/Skip/TakeWhile/SkipWhile
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-part-23-take-skip-takewhile-skipwhile/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.take
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takelast
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skiplast
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

// Take returns a specified number of contiguous elements from the start of a sequence.
func Take[Source any](source Enumerator[Source], count int) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return Empty[Source](), nil
	}
	i := 0
	return OnFunc[Source]{
			mvNxt: func() bool {
				if i < count && source.MoveNext() {
					i++
					return true
				}
				return false
			},
			crrnt: func() Source { return source.Current() },
			rst:   func() { i = 0; source.Reset() },
		},
		nil
}

// TakeMust is like Take but panics in case of error.
func TakeMust[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	r, err := Take(source, count)
	if err != nil {
		panic(err)
	}
	return r
}

// TakeLast returns a new enumerable collection that contains the last 'count' elements from 'source'.
func TakeLast[Source any](source Enumerator[Source], count int) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return Empty[Source](), nil
	}
	sl := Slice(source)
	return NewOnSlice(sl[len(sl)-count:]...), nil
}

// TakeLastMust is like TakeLast but panics in case of error.
func TakeLastMust[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	r, err := TakeLast(source, count)
	if err != nil {
		panic(err)
	}
	return r
}

// TakeWhile returns elements from a sequence as long as a specified condition is true.
func TakeWhile[Source any](source Enumerator[Source], predicate func(Source) bool) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	enough := false
	var c Source
	return OnFunc[Source]{
			mvNxt: func() bool {
				if enough {
					return false
				}
				if source.MoveNext() {
					c = source.Current()
					if predicate(c) {
						return true
					}
				}
				enough = true
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { enough = false; source.Reset() },
		},
		nil
}

// TakeWhileMust is like TakeWhile but panics in case of error.
func TakeWhileMust[Source any](source Enumerator[Source], predicate func(Source) bool) Enumerator[Source] {
	r, err := TakeWhile(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// TakeWhileIdx returns elements from a sequence as long as a specified condition is true.
// The element's index is used in the logic of the predicate function.
func TakeWhileIdx[Source any](source Enumerator[Source], predicate func(Source, int) bool) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	enough := false
	var c Source
	i := -1 // position before the first element
	return OnFunc[Source]{
			mvNxt: func() bool {
				if enough {
					return false
				}
				if source.MoveNext() {
					c = source.Current()
					i++
					if predicate(c, i) {
						return true
					}
				}
				enough = true
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { enough = false; i = -1; source.Reset() },
		},
		nil
}

// TakeWhileIdxMust is like TakeWhileIdx but panics in case of error.
func TakeWhileIdxMust[Source any](source Enumerator[Source], predicate func(Source, int) bool) Enumerator[Source] {
	r, err := TakeWhileIdx(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
func Skip[Source any](source Enumerator[Source], count int) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return source, nil
	}
	i := 1
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					if i > count {
						return true
					}
					i++
				}
				return false
			},
			crrnt: func() Source { return source.Current() },
			rst:   func() { i = 1; source.Reset() },
		},
		nil
}

// SkipMust is like Skip but panics in case of error.
func SkipMust[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	r, err := Skip(source, count)
	if err != nil {
		panic(err)
	}
	return r
}

// SkipLast returns a new enumerable collection that contains the elements from 'source'
// with the last 'count' elements of the source collection omitted.
func SkipLast[Source any](source Enumerator[Source], count int) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return source, nil
	}
	sl := Slice(source)
	return NewOnSlice(sl[:len(sl)-count]...), nil
}

// SkipLastMust is like SkipLast but panics in case of error.
func SkipLastMust[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	r, err := SkipLast(source, count)
	if err != nil {
		panic(err)
	}
	return r
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
func SkipWhile[Source any](source Enumerator[Source], predicate func(Source) bool) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	var c Source
	remaining := false
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
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
			rst:   func() { remaining = false; source.Reset() },
		},
		nil
}

// SkipWhileMust is like SkipWhile but panics in case of error.
func SkipWhileMust[Source any](source Enumerator[Source], predicate func(Source) bool) Enumerator[Source] {
	r, err := SkipWhile(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// SkipWhileIdx bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// The element's index is used in the logic of the predicate function.
func SkipWhileIdx[Source any](source Enumerator[Source], predicate func(Source, int) bool) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	remaining := false
	var c Source
	i := -1 // position before the first element
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
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
			rst:   func() { remaining = false; i = -1; source.Reset() },
		},
		nil
}

// SkipWhileIdxMust is like SkipWhileIdx but panics in case of error.
func SkipWhileIdxMust[Source any](source Enumerator[Source], predicate func(Source, int) bool) Enumerator[Source] {
	r, err := SkipWhileIdx(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
