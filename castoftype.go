//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 33 – Cast and OfType
// https://codeblog.jonskeet.uk/2011/01/13/reimplementing-linq-to-objects-part-33-cast-and-oftype/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.cast
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.oftype

// Cast casts the elements of an Enumerator to the specified type.
func Cast[Source, Result any](source Enumerator[Source]) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OnFunc[Result]{
			mvNxt: func() bool { return source.MoveNext() },
			crrnt: func() Result {
				var i interface{} = source.Current()
				return i.(Result)
			},
			rst: func() { source.Reset() },
		},
		nil
}

// CastMust is like Cast but panics in case of error.
func CastMust[Source, Result any](source Enumerator[Source]) Enumerator[Result] {
	r, err := Cast[Source, Result](source)
	if err != nil {
		panic(err)
	}
	return r
}

// OfType filters the elements of an Enumerator based on a specified type.
func OfType[Source, Result any](source Enumerator[Source]) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
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
		},
		nil
}

// OfTypeMust is like OfType but panics in case of error.
func OfTypeMust[Source, Result any](source Enumerator[Source]) Enumerator[Result] {
	r, err := OfType[Source, Result](source)
	if err != nil {
		panic(err)
	}
	return r
}
