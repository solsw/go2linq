//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 4 - Range
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-4-range/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.range

func enrRange(start, count int) func() Enumerator[int] {
	return func() Enumerator[int] {
		i := 0
		return enrFunc[int]{
			mvNxt: func() bool {
				if i < count {
					i++
					return true
				}
				return false
			},
			crrnt: func() int { return start + i - 1 },
			rst:   func() { i = 0 },
		}
	}
}

// Range generates a sequence of ints within a specified range.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.range)
func Range(start, count int) (Enumerable[int], error) {
	if count < 0 {
		return nil, ErrNegativeCount
	}
	return OnFactory(enrRange(start, count)), nil
}

// RangeMust is like Range but panics in case of error.
func RangeMust(start, count int) Enumerable[int] {
	r, err := Range(start, count)
	if err != nil {
		panic(err)
	}
	return r
}
