//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 8 - Concat
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-8-concat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat

// Concat concatenates two sequences.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ConcatSelf instead.
// Concat panics if 'first' or 'second' is nil.
func Concat[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	from1 := true
	return OnFunc[Source]{
		MvNxt: func() bool {
			if from1 && first.MoveNext() {
				return true
			}
			from1 = false
			return second.MoveNext()
		},
		Crrnt: func() Source {
			if from1 {
				return first.Current()
			}
			return second.Current()
		},
		Rst: func() { 
			first.Reset()
			if !from1 {
				from1 = true
				second.Reset()
			}
		},
	}
}

// ConcatErr is like Concat but returns an error instead of panicking.
func ConcatErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Concat(first, second), nil
}

// ConcatSelf concatenates two sequences.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// ConcatSelf panics if 'first' or 'second' is nil.
func ConcatSelf[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return Concat(first, NewOnSlice(sl2...))
}

// ConcatSelfErr is like ConcatSelf but returns an error instead of panicking.
func ConcatSelfErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return ConcatSelf(first, second), nil
}
