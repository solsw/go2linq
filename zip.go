//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 35 – Zip
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-35-zip/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.zip

// Zip applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ZipSelf instead.
// Zip panics if 'first' or 'second' or 'resultSelector' is nil.
func Zip[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) Enumerator[Result] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if resultSelector == nil {
		panic(ErrNilSelector)
	}
	return OnFunc[Result]{
		MvNxt: func() bool {
			if first.MoveNext() && second.MoveNext() {
				return true
			}
			return false
		},
		Crrnt: func() Result { return resultSelector(first.Current(), second.Current()) },
		Rst:   func() { first.Reset(); second.Reset() },
	}
}

// ZipErr is like Zip but returns an error instead of panicking.
func ZipErr[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return Zip(first, second, resultSelector), nil
}

// ZipSelf applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// (See Test_ZipSelf for examples.)
// ZipSelf panics if 'first' or 'second' or 'resultSelector' is nil.
func ZipSelf[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) Enumerator[Result] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if resultSelector == nil {
		panic(ErrNilSelector)
	}
	sl2 := Slice(second)
	first.Reset()
	return Zip(first, NewOnSlice(sl2...), resultSelector)
}

// ZipSelfErr is like ZipSelf but returns an error instead of panicking.
func ZipSelfErr[First, Second, Result any](first Enumerator[First], second Enumerator[Second],
	resultSelector func(First, Second) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return ZipSelf(first, second, resultSelector), nil
}
