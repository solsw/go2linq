//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 32 – Contains
// https://codeblog.jonskeet.uk/2011/01/12/reimplementing-linq-to-objects-part-32-contains/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains

// Contains determines whether a sequence contains a specified element by using reflect.DeepEqual.
// Contains panics if 'source' is nil.
func Contains[Source any](source Enumerator[Source], value Source) bool {
	if source == nil {
		panic(ErrNilSource)
	}
	return ContainsEq(source, value, nil)
}

// ContainsErr is like Contains but returns an error instead of panicking.
func ContainsErr[Source any](source Enumerator[Source], value Source) (res bool, err error) {
	defer func() {
		catchErrPanic[bool](recover(), &res, &err)
	}()
	return Contains(source, value), nil
}

// ContainsEq determines whether a sequence contains a specified element by using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used.
func ContainsEq[Source any](source Enumerator[Source], value Source, eq Equaler[Source]) bool {
	if source == nil {
		panic(ErrNilSource)
	}
	if eq == nil {
		eq = EqualerFunc[Source](DeepEqual[Source])
	}
	for source.MoveNext() {
		if eq.Equal(value, source.Current()) {
			return true
		}
	}
	return false
}

// ContainsEqErr is like ContainsEq but returns an error instead of panicking.
func ContainsEqErr[Source any](source Enumerator[Source], value Source, eq Equaler[Source]) (res bool, err error) {
	defer func() {
		catchErrPanic[bool](recover(), &res, &err)
	}()
	return ContainsEq(source, value, eq), nil
}
