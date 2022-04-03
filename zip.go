//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 35 – Zip
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-35-zip/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.zip

func enrZip[First, Second, Result any](first Enumerable[First], second Enumerable[Second],
	resultSelector func(First, Second) Result) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enr1 := first.GetEnumerator()
		enr2 := second.GetEnumerator()
		return enrFunc[Result]{
			mvNxt: func() bool {
				if enr1.MoveNext() && enr2.MoveNext() {
					return true
				}
				return false
			},
			crrnt: func() Result { return resultSelector(enr1.Current(), enr2.Current()) },
			rst:   func() { enr1.Reset(); enr2.Reset() },
		}
	}
}

// Zip applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.zip)
func Zip[First, Second, Result any](first Enumerable[First], second Enumerable[Second],
	resultSelector func(First, Second) Result) (Enumerable[Result], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if resultSelector == nil {
		return nil, ErrNilSelector
	}
	return OnFactory(enrZip(first, second, resultSelector)), nil
}

// ZipMust is like Zip but panics in case of an error.
func ZipMust[First, Second, Result any](first Enumerable[First], second Enumerable[Second],
	resultSelector func(First, Second) Result) Enumerable[Result] {
	r, err := Zip(first, second, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
