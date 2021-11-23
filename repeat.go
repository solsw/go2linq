//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 6 - Repeat
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-6-repeat/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat

// Repeat generates a sequence that contains one repeated value.
// Repeat panics if 'count' is less than 0.
func Repeat[Result any](element Result, count int) Enumerator[Result] {
	if count < 0 {
		panic(ErrNegativeCount)
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
	}
}

// RepeatErr is like Repeat but returns an error instead of panicking.
func RepeatErr[Result any](element Result, count int) (res Enumerator[Result], err error) {
	defer func() {
		catchErrPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return Repeat(element, count), nil
}
