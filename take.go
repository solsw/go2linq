package go2linq

import (
	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 23 - Take/Skip/TakeWhile/SkipWhile
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-part-23-take-skip-takewhile-skipwhile/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.take
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takelast
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile

func factoryTake[Source any](source Enumerable[Source], count int) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		i := 0
		return enrFunc[Source]{
			mvNxt: func() bool {
				if i < count && enr.MoveNext() {
					i++
					return true
				}
				return false
			},
			crrnt: func() Source { return enr.Current() },
			rst:   func() { i = 0; enr.Reset() },
		}
	}
}

// [Take] returns a specified number of contiguous elements from the start of a sequence.
//
// [Take]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.take
func Take[Source any](source Enumerable[Source], count int) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return Empty[Source](), nil
	}
	return OnFactory(factoryTake(source, count)), nil
}

// TakeMust is like [Take] but panics in case of error.
func TakeMust[Source any](source Enumerable[Source], count int) Enumerable[Source] {
	return generichelper.Must(Take(source, count))
}

// [TakeLast] returns a new [Enumerable] that contains the last 'count' elements from 'source'.
//
// [TakeLast]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takelast
func TakeLast[Source any](source Enumerable[Source], count int) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return Empty[Source](), nil
	}
	sl := ToSliceMust(source)
	return NewEnSlice(sl[len(sl)-count:]...), nil
}

// TakeLastMust is like [TakeLast] but panics in case of error.
func TakeLastMust[Source any](source Enumerable[Source], count int) Enumerable[Source] {
	return generichelper.Must(TakeLast(source, count))
}

func factoryTakeWhile[Source any](source Enumerable[Source], predicate func(Source) bool) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		enough := false
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				if enough {
					return false
				}
				if enr.MoveNext() {
					c = enr.Current()
					if predicate(c) {
						return true
					}
				}
				enough = true
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { enough = false; enr.Reset() },
		}
	}
}

// [TakeWhile] returns elements from a sequence as long as a specified condition is true.
//
// [TakeWhile]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func TakeWhile[Source any](source Enumerable[Source], predicate func(Source) bool) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return OnFactory(factoryTakeWhile(source, predicate)), nil
}

// TakeWhileMust is like [TakeWhile] but panics in case of error.
func TakeWhileMust[Source any](source Enumerable[Source], predicate func(Source) bool) Enumerable[Source] {
	return generichelper.Must(TakeWhile(source, predicate))
}

func factoryTakeWhileIdx[Source any](source Enumerable[Source], predicate func(Source, int) bool) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		enough := false
		var c Source
		i := -1 // position before the first element
		return enrFunc[Source]{
			mvNxt: func() bool {
				if enough {
					return false
				}
				if enr.MoveNext() {
					c = enr.Current()
					i++
					if predicate(c, i) {
						return true
					}
				}
				enough = true
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { enough = false; i = -1; enr.Reset() },
		}
	}
}

// [TakeWhileIdx] returns elements from a sequence as long as a specified condition is true.
// The element's index is used in the logic of the predicate function.
//
// [TakeWhileIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func TakeWhileIdx[Source any](source Enumerable[Source], predicate func(Source, int) bool) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return OnFactory(factoryTakeWhileIdx(source, predicate)), nil
}

// TakeWhileIdxMust is like [TakeWhileIdx] but panics in case of error.
func TakeWhileIdxMust[Source any](source Enumerable[Source], predicate func(Source, int) bool) Enumerable[Source] {
	return generichelper.Must(TakeWhileIdx(source, predicate))
}
