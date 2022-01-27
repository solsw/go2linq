//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 17 – Except
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-17-except/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

// Except produces the set difference of two sequences using DeepEqual to compare values.
// 'second' is enumerated on the first MoveNext call.
// Order of elements in the result corresponds to the order of elements in 'first'.
func Except[Source any](first, second Enumerable[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return ExceptEq(first, second, nil)
}

// ExceptMust is like Except but panics in case of error.
func ExceptMust[Source any](first, second Enumerable[Source]) Enumerable[Source] {
	r, err := Except(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptEq produces the set difference of two sequences using the specified Equaler to compare values.
// If 'equaler' is nil DeepEqual is used. 'second' is enumerated on the first MoveNext call.
// Order of elements in the result corresponds to the order of elements in 'first'.
func ExceptEq[Source any](first, second Enumerable[Source], equaler Equaler[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if equaler == nil {
		equaler = DeepEqual[Source]{}
	}
	return ExceptByEq(first, second, Identity[Source], equaler)
}

// ExceptEqMust is like ExceptEq but panics in case of error.
func ExceptEqMust[Source any](first, second Enumerable[Source], equaler Equaler[Source]) Enumerable[Source] {
	r, err := ExceptEq(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptCmp produces the set difference of two sequences using a specified Comparer to compare values.
// (See DistinctCmp function.) 'second' is enumerated on the first MoveNext call.
// Order of elements in the result corresponds to the order of elements in 'first'.
func ExceptCmp[Source any](first, second Enumerable[Source], comparer Comparer[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return ExceptByCmp(first, second, Identity[Source], comparer)
}

// ExceptCmpMust is like ExceptCmp but panics in case of error.
func ExceptCmpMust[Source any](first, second Enumerable[Source], comparer Comparer[Source]) Enumerable[Source] {
	r, err := ExceptCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
