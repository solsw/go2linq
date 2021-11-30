//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 6 - Repeat
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-6-repeat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat

// Repeat generates a sequence that contains one repeated value.
func Repeat[Result any](element Result, count int) (Enumerator[Result], error) {
	if count < 0 {
		return nil, ErrNegativeCount
	}
	i := 0
	return OnFunc[Result]{
			mvNxt: func() bool {
				if i >= count {
					return false
				}
				i++
				return true
			},
			crrnt: func() Result { return element },
			rst:   func() { i = 0 },
		},
		nil
}

// RepeatMust is like Repeat but panics in case of error.
func RepeatMust[Result any](element Result, count int) Enumerator[Result] {
	r, err := Repeat(element, count)
	if err != nil {
		panic(err)
	}
	return r
}
