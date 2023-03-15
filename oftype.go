package go2linq

// Reimplementing LINQ to Objects: Part 33 – Cast and OfType
// https://codeblog.jonskeet.uk/2011/01/13/reimplementing-linq-to-objects-part-33-cast-and-oftype/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.oftype

func factoryOfType[Source, Result any](source Enumerable[Source]) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enr := source.GetEnumerator()
		var r Result
		return enrFunc[Result]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					var i any = enr.Current()
					var ok bool
					r, ok = i.(Result)
					if ok {
						return true
					}
				}
				return false
			},
			crrnt: func() Result { return r },
			rst:   func() { enr.Reset() },
		}
	}
}

// OfType filters the elements of an Enumerable based on a specified type.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.oftype)
func OfType[Source, Result any](source Enumerable[Source]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OnFactory(factoryOfType[Source, Result](source)), nil
}

// OfTypeMust is like [OfType] but panics in case of error.
func OfTypeMust[Source, Result any](source Enumerable[Source]) Enumerable[Result] {
	r, err := OfType[Source, Result](source)
	if err != nil {
		panic(err)
	}
	return r
}
