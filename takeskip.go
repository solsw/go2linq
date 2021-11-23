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
// Take panics if 'source' is nil.
func Take[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if count <= 0 {
		return Empty[Source]()
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
	}
}

// TakeErr is like Take but returns an error instead of panicking.
func TakeErr[Source any](source Enumerator[Source], count int) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Take(source, count), nil
}

// TakeLast returns a new enumerable collection that contains the last 'count' elements from 'source'.
// TakeLast panics if 'source' is nil.
func TakeLast[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if count <= 0 {
		return Empty[Source]()
	}
	sl := Slice(source)
	return NewOnSlice(sl[len(sl)-count:]...)
}

// TakeLastErr is like TakeLast but returns an error instead of panicking.
func TakeLastErr[Source any](source Enumerator[Source], count int) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return TakeLast(source, count), nil
}

// TakeWhile returns elements from a sequence as long as a specified condition is true.
// TakeWhile panics if 'source' or 'predicate' is nil.
func TakeWhile[Source any](source Enumerator[Source], predicate func(Source) bool) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
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
	}
}

// TakeWhileErr is like TakeWhile but returns an error instead of panicking.
func TakeWhileErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return TakeWhile(source, predicate), nil
}

// TakeWhileIdx returns elements from a sequence as long as a specified condition is true.
// The element's index is used in the logic of the predicate function.
// TakeWhileIdx panics if 'source' or 'predicate' is nil.
func TakeWhileIdx[Source any](source Enumerator[Source], predicate func(Source, int) bool) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
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
	}
}

// TakeWhileIdxErr is like TakeWhileIdx but returns an error instead of panicking.
func TakeWhileIdxErr[Source any](source Enumerator[Source], predicate func(Source, int) bool) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return TakeWhileIdx(source, predicate), nil
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
// Skip panics if 'source' is nil.
func Skip[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if count <= 0 {
		return source
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
	}
}

// SkipErr is like Skip but returns an error instead of panicking.
func SkipErr[Source any](source Enumerator[Source], count int) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Skip(source, count), nil
}

// SkipLast returns a new enumerable collection that contains the elements from 'source'
// with the last 'count' elements of the source collection omitted.
// SkipLast panics if 'source' is nil.
func SkipLast[Source any](source Enumerator[Source], count int) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if count <= 0 {
		return source
	}
	sl := Slice(source)
	return NewOnSlice(sl[:len(sl)-count]...)
}

// SkipLastErr is like SkipLast but returns an error instead of panicking.
func SkipLastErr[Source any](source Enumerator[Source], count int) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return SkipLast(source, count), nil
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// SkipWhile panics if 'source' or 'predicate' is nil.
func SkipWhile[Source any](source Enumerator[Source], predicate func(Source) bool) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
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
	}
}

// SkipWhileErr is like SkipWhile but returns an error instead of panicking.
func SkipWhileErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return SkipWhile(source, predicate), nil
}

// SkipWhileIdx bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// The element's index is used in the logic of the predicate function.
// SkipWhileIdx panics if 'source' or 'predicate' is nil.
func SkipWhileIdx[Source any](source Enumerator[Source], predicate func(Source, int) bool) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
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
	}
}

// SkipWhileIdxErr is like SkipWhileIdx but returns an error instead of panicking.
func SkipWhileIdxErr[Source any](source Enumerator[Source], predicate func(Source, int) bool) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return SkipWhileIdx(source, predicate), nil
}
