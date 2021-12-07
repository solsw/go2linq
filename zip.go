//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 35 – Zip
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-35-zip/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.zip

// Zip applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ZipSelf instead.
func Zip[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) (Enumerator[Result], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if resultSelector == nil {
		return nil, ErrNilSelector
	}
	return OnFunc[Result]{
			mvNxt: func() bool {
				if first.MoveNext() && second.MoveNext() {
					return true
				}
				return false
			},
			crrnt: func() Result { return resultSelector(first.Current(), second.Current()) },
			rst:   func() { first.Reset(); second.Reset() },
		},
		nil
}

// ZipMust is like Zip but panics in case of error.
func ZipMust[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) Enumerator[Result] {
	r, err := Zip(first, second, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ZipSelf applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// (See Test_ZipSelf for examples.)
func ZipSelf[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) (Enumerator[Result], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if resultSelector == nil {
		return nil, ErrNilSelector
	}
	sl2 := Slice(second)
	first.Reset()
	return Zip(first, NewOnSliceEn(sl2...), resultSelector)
}

// ZipSelfMust is like ZipSelf but panics in case of error.
func ZipSelfMust[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) Enumerator[Result] {
	r, err := ZipSelf(first, second, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
