//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 32 – Contains
// https://codeblog.jonskeet.uk/2011/01/12/reimplementing-linq-to-objects-part-32-contains/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains

// Contains determines whether a sequence contains a specified element using DeepEqualer.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains)
func Contains[Source any](source Enumerable[Source], value Source) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	return ContainsEq(source, value, nil)
}

// ContainsMust is like Contains but panics in case of error.
func ContainsMust[Source any](source Enumerable[Source], value Source) bool {
	r, err := Contains(source, value)
	if err != nil {
		panic(err)
	}
	return r
}

// ContainsEq determines whether a sequence contains a specified element using a specified Equaler.
// If 'equaler' is nil DeepEqualer is used.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains)
func ContainsEq[Source any](source Enumerable[Source], value Source, equaler Equaler[Source]) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if equaler == nil {
		equaler = DeepEqualer[Source]{}
	}
	enr := source.GetEnumerator()
	for enr.MoveNext() {
		if equaler.Equal(value, enr.Current()) {
			return true, nil
		}
	}
	return false, nil
}

// ContainsEqMust is like ContainsEq but panics in case of error.
func ContainsEqMust[Source any](source Enumerable[Source], value Source, equaler Equaler[Source]) bool {
	r, err := ContainsEq(source, value, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
