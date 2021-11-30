//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 12 - DefaultIfEmpty
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-12-defaultifempty/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

// DefaultIfEmpty returns the elements of the specified sequence
// or the type parameter's default value in a singleton collection if the sequence is empty.
func DefaultIfEmpty[Source any](source Enumerator[Source]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	var s0 Source
	return DefaultIfEmptyDef(source, s0)
}

// DefaultIfEmptyMust is like DefaultIfEmpty but panics in case of error.
func DefaultIfEmptyMust[Source any](source Enumerator[Source]) Enumerator[Source] {
	r, err := DefaultIfEmpty(source)
	if err != nil {
		panic(err)
	}
	return r
}

// DefaultIfEmptyDef returns the elements of the specified sequence
// or the specified value in a singleton collection if the sequence is empty.
func DefaultIfEmptyDef[Source any](source Enumerator[Source], defaultValue Source) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	first := true
	empty := false
	return OnFunc[Source]{
			mvNxt: func() bool {
				if first {
					first = false
					if !source.MoveNext() {
						empty = true
					}
					return true
				}
				if empty {
					return false
				}
				return source.MoveNext()
			},
			crrnt: func() Source {
				if empty {
					return defaultValue
				}
				return source.Current()
			},
			rst: func() { first = true; empty = false; source.Reset() },
		},
		nil
}

// DefaultIfEmptyDefMust is like DefaultIfEmptyDef but panics in case of error.
func DefaultIfEmptyDefMust[Source any](source Enumerator[Source], defaultValue Source) Enumerator[Source] {
	r, err := DefaultIfEmptyDef(source, defaultValue)
	if err != nil {
		panic(err)
	}
	return r
}
