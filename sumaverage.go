//go:build go1.18

package go2linq

import (
	"constraints"
	"math"
)

// Reimplementing LINQ to Objects: Part 28 – Sum
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-28-sum/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

// Reimplementing LINQ to Objects: Part 30 – Average
// https://codeblog.jonskeet.uk/2011/01/10/reimplementing-linq-to-objects-part-30-average/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func sumIntegerPrim[Source any, Result constraints.Integer](source Enumerator[Source], selector func(Source) Result) (Result, int) {
	var sum Result = 0
	count := 0
	for source.MoveNext() {
		sum += selector(source.Current())
		count++
	}
	return sum, count
}

func sumFloatPrim[Source any, Result constraints.Float](source Enumerator[Source], selector func(Source) Result) (Result, int) {
	var sum Result = 0
	count := 0
	for source.MoveNext() {
		sum += selector(source.Current())
		count++
	}
	return sum, count
}

// SumInteger computes the sum of a sequence of constraints.Integer values that are obtained
// by invoking a transform function on each element of the input sequence.
func SumInteger[Source any, Result constraints.Integer](source Enumerator[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	if selector == nil {
		return 0, ErrNilSelector
	}
	r, _ := sumIntegerPrim(source, selector)
	return r, nil
}

// SumIntegerMust is like SumInteger but panics in case of error.
func SumIntegerMust[Source any, Result constraints.Integer](source Enumerator[Source], selector func(Source) Result) Result {
	r, err := SumInteger(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// SumFloat computes the sum of a sequence of constraints.Float values that are obtained
// by invoking a transform function on each element of the input sequence.
func SumFloat[Source any, Result constraints.Float](source Enumerator[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	if selector == nil {
		return 0, ErrNilSelector
	}
	r, _ := sumFloatPrim(source, selector)
	return r, nil
}

// SumFloatMust is like SumFloat but panics in case of error.
func SumFloatMust[Source any, Result constraints.Float](source Enumerator[Source], selector func(Source) Result) Result {
	r, err := SumFloat(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// AverageInteger computes the average of a sequence of constraints.Integer values that are obtained
// by invoking a transform function on each element of the input sequence.
func AverageInteger[Source any, Result constraints.Integer](source Enumerator[Source], selector func(Source) Result) (float64, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	if selector == nil {
		return 0, ErrNilSelector
	}
	sum, count := sumIntegerPrim(source, selector)
	if count == 0 {
		return 0, ErrEmptySource
	}
	return (float64(sum) / float64(count)), nil
}

// AverageIntegerMust is like AverageInteger but panics in case of error.
func AverageIntegerMust[Source any, Result constraints.Integer](source Enumerator[Source], selector func(Source) Result) float64 {
	r, err := AverageInteger(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// AverageFloat computes the average of constraints.Float values that are obtained
// by invoking a transform function on each element of the input sequence.
func AverageFloat[Source any, Result constraints.Float](source Enumerator[Source], selector func(Source) Result) (float64, error) {
	if source == nil {
		return math.NaN(), ErrNilSource
	}
	if selector == nil {
		return math.NaN(), ErrNilSelector
	}
	sum, count := sumFloatPrim(source, selector)
	if count == 0 {
		return math.NaN(), ErrEmptySource
	}
	return (float64(sum) / float64(count)), nil
}

// AverageFloatMust is like AverageFloat but panics in case of error.
func AverageFloatMust[Source any, Result constraints.Float](source Enumerator[Source], selector func(Source) Result) float64 {
	r, err := AverageFloat(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}
