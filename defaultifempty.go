package go2linq

import (
	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 12 - DefaultIfEmpty
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-12-defaultifempty/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

// [DefaultIfEmpty] returns the elements of a specified sequence
// or the type parameter's [zero value] in a singleton collection if the sequence is empty.
//
// [DefaultIfEmpty]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
// [zero value]: https://go.dev/ref/spec#The_zero_value
func DefaultIfEmpty[Source any](source Enumerable[Source]) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return DefaultIfEmptyDef(source, generichelper.ZeroValue[Source]())
}

// DefaultIfEmptyMust is like [DefaultIfEmpty] but panics in case of error.
func DefaultIfEmptyMust[Source any](source Enumerable[Source]) Enumerable[Source] {
	return errorhelper.Must(DefaultIfEmpty(source))
}

func factoryDefaultIfEmptyDef[Source any](source Enumerable[Source], defaultValue Source) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		first := true
		empty := false
		return enrFunc[Source]{
			mvNxt: func() bool {
				if first {
					first = false
					if !enr.MoveNext() {
						empty = true
					}
					return true
				}
				if empty {
					return false
				}
				return enr.MoveNext()
			},
			crrnt: func() Source {
				if empty {
					return defaultValue
				}
				return enr.Current()
			},
			rst: func() {
				first = true
				empty = false
				enr.Reset()
			},
		}
	}
}

// [DefaultIfEmptyDef] returns the elements of a specified sequence
// or a specified value in a singleton collection if the sequence is empty.
//
// [DefaultIfEmptyDef]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func DefaultIfEmptyDef[Source any](source Enumerable[Source], defaultValue Source) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OnFactory(factoryDefaultIfEmptyDef(source, defaultValue)), nil
}

// DefaultIfEmptyDefMust is like [DefaultIfEmptyDef] but panics in case of error.
func DefaultIfEmptyDefMust[Source any](source Enumerable[Source], defaultValue Source) Enumerable[Source] {
	return errorhelper.Must(DefaultIfEmptyDef(source, defaultValue))
}
