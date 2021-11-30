//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 4 - Range
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-4-range/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.range

// Range generates a sequence of ints within a specified range.
func Range(start, count int) (Enumerator[int], error) {
	if count < 0 {
		return nil, ErrNegativeCount
	}
	i := 0
	return OnFunc[int]{
			mvNxt: func() bool {
				if i < count {
					i++
					return true
				}
				return false
			},
			crrnt: func() int { return start + i - 1 },
			rst:   func() { i = 0 },
		},
		nil
}

// RangeMust is like Range but panics in case of error.
func RangeMust(start, count int) Enumerator[int] {
	r, err := Range(start, count)
	if err != nil {
		panic(err)
	}
	return r
}
