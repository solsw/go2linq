//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 12 - DefaultIfEmpty
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-12-defaultifempty/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

// DefaultIfEmpty returns the elements of the specified sequence
// or the type parameter's default value in a singleton collection if the sequence is empty.
// DefaultIfEmpty panics if 'source' is nil.
func DefaultIfEmpty[Source any](source Enumerator[Source]) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	var s0 Source
	return DefaultIfEmptyDef(source, s0)
}

// DefaultIfEmptyErr is like DefaultIfEmpty but returns an error instead of panicking.
func DefaultIfEmptyErr[Source any](source Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return DefaultIfEmpty(source), nil
}

// DefaultIfEmptyDef returns the elements of the specified sequence
// or the specified value in a singleton collection if the sequence is empty.
// DefaultIfEmptyDef panics if 'source' is nil.
func DefaultIfEmptyDef[Source any](source Enumerator[Source], defaultValue Source) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
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
	}
}

// DefaultIfEmptyDefErr is like DefaultIfEmptyDef but returns an error instead of panicking.
func DefaultIfEmptyDefErr[Source any](source Enumerator[Source], defaultValue Source) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return DefaultIfEmptyDef(source, defaultValue), nil
}
