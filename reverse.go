package go2linq

import (
	"sync"

	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 27 - Reverse
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-27-reverse/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.reverse

func factoryReverse[Source any](source Enumerable[Source]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		var once sync.Once
		var sl []Source
		var i int
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() { sl = ToSliceMust(source); i = len(sl) })
				if i > 0 {
					i--
					return true
				}
				return false
			},
			crrnt: func() Source { return sl[i] },
			rst:   func() { i = len(sl) },
		}
	}
}

// [Reverse] inverts the order of the elements in a sequence.
//
// [Reverse]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.reverse
func Reverse[Source any](source Enumerable[Source]) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OnFactory(factoryReverse(source)), nil
}

// ReverseMust is like [Reverse] but panics in case of error.
func ReverseMust[Source any](source Enumerable[Source]) Enumerable[Source] {
	return generichelper.Must(Reverse(source))
}
