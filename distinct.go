//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 14 - Distinct
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-14-distinct/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

// Distinct returns distinct elements from a sequence using reflect.DeepEqual to compare values.
func Distinct[Source any](source Enumerator[Source]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return DistinctEq(source, nil)
}

// DistinctMust is like Distinct but panics in case of error.
func DistinctMust[Source any](source Enumerator[Source]) Enumerator[Source] {
	r, err := Distinct(source)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctEq returns distinct elements from a sequence using a specified Equaler to compare values.
// If 'equaler' is nil reflect.DeepEqual is used.
func DistinctEq[Source any](source Enumerator[Source], equaler Equaler[Source]) (Enumerator[Source], error) {
	return DistinctByEq(source, Identity[Source], equaler)
}

// DistinctEqMust is like DistinctEq but panics in case of error.
func DistinctEqMust[Source any](source Enumerator[Source], equaler Equaler[Source]) Enumerator[Source] {
	r, err := DistinctEq(source, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctCmp returns distinct elements from a sequence using a specified Comparer to compare values.
//
// Sorted slice of already seen elements is internally built.
// Sorted slice allows to use binary search to determine whether the element was seen or not.
// This may give performance gain when processing large sequences (though this is a subject for benchmarking).
func DistinctCmp[Source any](source Enumerator[Source], comparer Comparer[Source]) (Enumerator[Source], error) {
	return DistinctByCmp(source, Identity[Source], comparer)
}

// DistinctCmpMust is like DistinctCmp but panics in case of error.
func DistinctCmpMust[Source any](source Enumerator[Source], comparer Comparer[Source]) Enumerator[Source] {
	r, err := DistinctCmp(source, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
