//go:build go1.18

package go2linq

import (
	"sync"
)

// Reimplementing LINQ to Objects: Part 27 - Reverse
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-27-reverse/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.reverse

// Reverse inverts the order of the elements in a sequence.
// Reverse panics if 'source' is nil.
func Reverse[Source any](source Enumerator[Source]) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	var once sync.Once
	var sl []Source
	var i int
	return OnFunc[Source]{
		MvNxt: func() bool {
			once.Do(func() { sl = Slice(source); i = len(sl) })
			if i > 0 {
				i--
				return true
			}
			return false
		},
		Crrnt: func() Source { return sl[i] },
		Rst:   func() { i = len(sl) },
	}
}

// ReverseErr is like Reverse but returns an error instead of panicking.
func ReverseErr[Source any](source Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Reverse(source), nil
}
