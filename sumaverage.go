//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 28 – Sum
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-28-sum/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

// Reimplementing LINQ to Objects: Part 30 – Average
// https://codeblog.jonskeet.uk/2011/01/10/reimplementing-linq-to-objects-part-30-average/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func sumIntCountPrim[Source any](source Enumerator[Source], selector func(Source) int) (int, int) {
	sum := 0
	count := 0
	for source.MoveNext() {
		sum += selector(source.Current())
		count++
	}
	return sum, count
}

func sumFloat64CountPrim[Source any](source Enumerator[Source], selector func(Source) float64) (float64, int) {
	sum := 0.0
	count := 0
	for source.MoveNext() {
		sum += selector(source.Current())
		count++
	}
	return sum, count
}

// SumInt computes the sum of a sequence of int values that are obtained
// by invoking a transform function on each element of the input sequence.
// SumInt panics if 'source' or 'selector' is nil.
func SumInt[Source any](source Enumerator[Source], selector func(Source) int) int {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	r, _ := sumIntCountPrim(source, selector)
	return r
}

// SumIntErr is like SumInt but returns an error instead of panicking.
func SumIntErr[Source any](source Enumerator[Source], selector func(Source) int) (res int, err error) {
	defer func() {
		catchErrPanic[int](recover(), &res, &err)
	}()
	return SumInt(source, selector), nil
}

// SumFloat64 computes the sum of a sequence of float64 values that are obtained
// by invoking a transform function on each element of the input sequence.
// SumFloat64 panics if 'source' or 'selector' is nil.
func SumFloat64[Source any](source Enumerator[Source], selector func(Source) float64) float64 {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	r, _ := sumFloat64CountPrim(source, selector)
	return r
}

// SumFloat64Err is like SumFloat64 but returns an error instead of panicking.
func SumFloat64Err[Source any](source Enumerator[Source], selector func(Source) float64) (res float64, err error) {
	defer func() {
		catchErrPanic[float64](recover(), &res, &err)
	}()
	return SumFloat64(source, selector), nil
}

// AverageInt computes the average of a sequence of int values that are obtained
// by invoking a transform function on each element of the input sequence.
// AverageInt panics if 'source' or 'selector' is nil or 'source' is empty.
func AverageInt[Source any](source Enumerator[Source], selector func(Source) int) float64 {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	sum, count := sumIntCountPrim(source, selector)
	if count == 0 {
		panic(ErrEmptySource)
	}
	return float64(sum) / float64(count)
}

// AverageIntErr is like AverageInt but returns an error instead of panicking.
func AverageIntErr[Source any](source Enumerator[Source], selector func(Source) int) (res float64, err error) {
	defer func() {
		catchErrPanic[float64](recover(), &res, &err)
	}()
	return AverageInt(source, selector), nil
}

// AverageFloat64 computes the average of float64 values that are obtained
// by invoking a transform function on each element of the input sequence.
// AverageFloat64 panics if 'source' or 'selector' is nil or 'source' is empty.
func AverageFloat64[Source any](source Enumerator[Source], selector func(Source) float64) float64 {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	sum, count := sumFloat64CountPrim(source, selector)
	if count == 0 {
		panic(ErrEmptySource)
	}
	return sum / float64(count)
}

// AverageFloat64Err is like AverageFloat64 but returns an error instead of panicking.
func AverageFloat64Err[Source any](source Enumerator[Source], selector func(Source) float64) (res float64, err error) {
	defer func() {
		catchErrPanic[float64](recover(), &res, &err)
	}()
	return AverageFloat64(source, selector), nil
}
