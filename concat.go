package go2linq

// Reimplementing LINQ to Objects: Part 8 - Concat
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-8-concat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat

func factoryConcat[Source any](first, second Enumerable[Source]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr1 := first.GetEnumerator()
		enr2 := second.GetEnumerator()
		from1 := true
		return enrFunc[Source]{
			mvNxt: func() bool {
				if from1 && enr1.MoveNext() {
					return true
				}
				from1 = false
				return enr2.MoveNext()
			},
			crrnt: func() Source {
				if from1 {
					return enr1.Current()
				}
				return enr2.Current()
			},
			rst: func() {
				enr1.Reset()
				if !from1 {
					from1 = true
					enr2.Reset()
				}
			},
		}
	}
}

// Concat concatenates two sequences.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat)
func Concat[Source any](first, second Enumerable[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return OnFactory(factoryConcat(first, second)), nil
}

// ConcatMust is like [Concat] but panics in case of error.
func ConcatMust[Source any](first, second Enumerable[Source]) Enumerable[Source] {
	r, err := Concat(first, second)
	if err != nil {
		panic(err)
	}
	return r
}
