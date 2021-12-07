//go:build go1.18

package go2linq

import (
	"sync"
)

// Reimplementing LINQ to Objects: Part 27 - Reverse
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-27-reverse/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.reverse

// Reverse inverts the order of the elements in a sequence.
func Reverse[Source any](source Enumerator[Source]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	var once sync.Once
	var sl []Source
	var i int
	return OnFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() { sl = Slice(source); i = len(sl) })
				if i > 0 {
					i--
					return true
				}
				return false
			},
			crrnt: func() Source { return sl[i] },
			rst:   func() { i = len(sl) },
		},
		nil
}

// ReverseMust is like Reverse but panics in case of error.
func ReverseMust[Source any](source Enumerator[Source]) Enumerator[Source] {
	r, err := Reverse(source)
	if err != nil {
		panic(err)
	}
	return r
}
