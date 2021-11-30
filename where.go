//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 2 - "Where"
// https://codeblog.jonskeet.uk/2010/09/03/reimplementing-linq-to-objects-part-2-quot-where-quot/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

// Where filters a sequence of values based on a predicate.
func Where[Source any](source Enumerator[Source], predicate func(Source) bool) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	var c Source
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
					if predicate(c) {
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { source.Reset() },
		},
		nil
}

// WhereMust is like Where but panics in case of error.
func WhereMust[Source any](source Enumerator[Source], predicate func(Source) bool) Enumerator[Source] {
	r, err := Where(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// WhereIdx filters a sequence of values based on a predicate.
// Each element's index is used in the logic of the predicate function.
func WhereIdx[Source any](source Enumerator[Source], predicate func(Source, int) bool) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	var c Source
	i := -1 // position before the first element
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
					i++
					if predicate(c, i) {
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { i = -1; source.Reset() },
		},
		nil
}

// WhereIdxMust is like WhereIdx but panics in case of error.
func WhereIdxMust[Source any](source Enumerator[Source], predicate func(Source, int) bool) Enumerator[Source] {
	r, err := WhereIdx(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
