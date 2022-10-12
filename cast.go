//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 33 – Cast and OfType
// https://codeblog.jonskeet.uk/2011/01/13/reimplementing-linq-to-objects-part-33-cast-and-oftype/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.cast

func factoryCast[Source, Result any](source Enumerable[Source]) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enr := source.GetEnumerator()
		return enrFunc[Result]{
			mvNxt: func() bool { return enr.MoveNext() },
			crrnt: func() Result {
				var i any = enr.Current()
				return i.(Result)
				// r, ok := i.(Result)
				// if !ok {
				// 	panic(ErrInvalidCast)
				// }
				// return r
			},
			rst: func() { enr.Reset() },
		}
	}
}

// Cast casts the elements of an Enumerable to the specified type.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.cast)
func Cast[Source, Result any](source Enumerable[Source]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OnFactory(factoryCast[Source, Result](source)), nil
}

// CastMust is like Cast but panics in case of error.
func CastMust[Source, Result any](source Enumerable[Source]) Enumerable[Result] {
	r, err := Cast[Source, Result](source)
	if err != nil {
		panic(err)
	}
	return r
}
