//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 6 - Repeat
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-6-repeat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat

func factoryRepeat[Result any](element Result, count int) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		i := 0
		return enrFunc[Result]{
			mvNxt: func() bool {
				if i >= count {
					return false
				}
				i++
				return true
			},
			crrnt: func() Result { return element },
			rst:   func() { i = 0 },
		}
	}
}

// Repeat generates a sequence that contains one repeated value.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat)
func Repeat[Result any](element Result, count int) (Enumerable[Result], error) {
	if count < 0 {
		return nil, ErrNegativeCount
	}
	return OnFactory(factoryRepeat(element, count)), nil
}

// RepeatMust is like Repeat but panics in case of error.
func RepeatMust[Result any](element Result, count int) Enumerable[Result] {
	r, err := Repeat(element, count)
	if err != nil {
		panic(err)
	}
	return r
}
