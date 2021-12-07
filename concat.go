//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 8 - Concat
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-8-concat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat

// Concat concatenates two sequences.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ConcatSelf instead.
func Concat[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	from1 := true
	return OnFunc[Source]{
			mvNxt: func() bool {
				if from1 && first.MoveNext() {
					return true
				}
				from1 = false
				return second.MoveNext()
			},
			crrnt: func() Source {
				if from1 {
					return first.Current()
				}
				return second.Current()
			},
			rst: func() {
				first.Reset()
				if !from1 {
					from1 = true
					second.Reset()
				}
			},
		},
		nil
}

// ConcatMust is like Concat but panics in case of error.
func ConcatMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := Concat(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// ConcatSelf concatenates two sequences.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func ConcatSelf[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return Concat(first, NewOnSliceEn(sl2...))
}

// ConcatSelfMust is like ConcatSelf but panics in case of error.
func ConcatSelfMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := ConcatSelf(first, second)
	if err != nil {
		panic(err)
	}
	return r
}
