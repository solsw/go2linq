//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 2 - "Where"
// https://codeblog.jonskeet.uk/2010/09/03/reimplementing-linq-to-objects-part-2-quot-where-quot/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

// Where filters a sequence of values based on a predicate.
// Where panics if 'source' or 'predicate' is nil.
func Where[Source any](source Enumerator[Source], predicate func(Source) bool) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	var c Source
	return OnFunc[Source]{
		MvNxt: func() bool {
			for source.MoveNext() {
				c = source.Current()
				if predicate(c) {
					return true
				}
			}
			return false
		},
		Crrnt: func() Source { return c },
		Rst: func() { source.Reset() },
	}
}

// WhereErr is like Where but returns an error instead of panicking.
func WhereErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Where(source, predicate), nil
}

// WhereIdx filters a sequence of values based on a predicate.
// Each element's index is used in the logic of the predicate function.
// WhereIdx panics if 'source' or 'predicate' is nil.
func WhereIdx[Source any](source Enumerator[Source], predicate func(Source, int) bool) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	var c Source
	i := -1 // position before the first element
	return OnFunc[Source]{
		MvNxt: func() bool {
			for source.MoveNext() {
				c = source.Current()
				i++
				if predicate(c, i) {
					return true
				}
			}
			return false
		},
		Crrnt: func() Source { return c },
		Rst: func() { i = -1; source.Reset() },
	}
}

// WhereIdxErr is like WhereIdx but returns an error instead of panicking.
func WhereIdxErr[Source any](source Enumerator[Source], predicate func(Source, int) bool) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return WhereIdx(source, predicate), nil
}
