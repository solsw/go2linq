//go:build go1.18

package go2linq

import (
	"math"
)

// Reimplementing LINQ to Objects: Part 4 - Range
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-4-range/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.range

// Range generates a sequence of ints within a specified range.
// Range panics if 'count' is less than 0 or 'start' + 'count' - 1 is larger than math.MaxInt32.
func Range(start, count int) Enumerator[int] {
	if count < 0 {
		panic(ErrNegativeCount)
	}
	if int64(start)+int64(count)-int64(1) > math.MaxInt32 {
		panic(ErrStartCount)
	}
	i := 0
	return OnFunc[int]{
		MvNxt: func() bool {
			if i < count {
				i++
				return true
			}
			return false
		},
		Crrnt: func() int { return start + i - 1 },
		Rst:   func() { i = 0 },
	}
}

// RangeErr is like Range but returns an error instead of panicking.
func RangeErr(start, count int) (res Enumerator[int], err error) {
	defer func() {
		catchPanic[Enumerator[int]](recover(), &res, &err)
	}()
	return Range(start, count), nil
}
