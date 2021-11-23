//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 33 – Cast and OfType
// https://codeblog.jonskeet.uk/2011/01/13/reimplementing-linq-to-objects-part-33-cast-and-oftype/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.cast
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.oftype

// Cast casts the elements of an Enumerator to the specified type.
// Cast panics if 'source' is nil or an element in the sequence cannot be cast to type 'Result'.
func Cast[Source, Result any](source Enumerator[Source]) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	return OnFunc[Result]{
		mvNxt: func() bool { return source.MoveNext() },
		crrnt: func() Result {
			var i interface{} = source.Current()
			r, ok := i.(Result)
			if !ok {
				panic(ErrInvalidCast)
			}
			return r
		},
		rst: func() { source.Reset() },
	}
}

// CastErr is like Cast but returns an error instead of panicking.
func CastErr[Source, Result any](source Enumerator[Source]) (res Enumerator[Result], err error) {
	defer func() {
		catchErrPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return Cast[Source, Result](source), nil
}

// OfType filters the elements of an Enumerator based on a specified type.
// OfType panics if 'source' is nil.
func OfType[Source, Result any](source Enumerator[Source]) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	var r Result
	return OnFunc[Result]{
		mvNxt: func() bool {
			for source.MoveNext() {
				var i interface{} = source.Current()
				var ok bool
				r, ok = i.(Result)
				if ok {
					return true
				}
			}
			return false
		},
		crrnt: func() Result { return r },
		rst:   func() { source.Reset() },
	}
}

// OfTypeErr is like OfType but returns an error instead of panicking.
func OfTypeErr[Source, Result any](source Enumerator[Source]) (res Enumerator[Result], err error) {
	defer func() {
		catchErrPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return OfType[Source, Result](source), nil
}
