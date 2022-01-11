//go:build go1.18

package go2linq

import (
	"constraints"
)

// Reimplementing LINQ to Objects: Part 28 – Sum
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-28-sum/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

// Reimplementing LINQ to Objects: Part 30 – Average
// https://codeblog.jonskeet.uk/2011/01/10/reimplementing-linq-to-objects-part-30-average/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func sumPrim[Source any, Result constraints.Integer | constraints.Float](source Enumerator[Source], selector func(Source) Result) (Result, int) {
	var sum Result = 0
	count := 0
	for source.MoveNext() {
		sum += selector(source.Current())
		count++
	}
	return sum, count
}

// Sum computes the sum of a sequence of constraints.Integer values that are obtained
// by invoking a transform function on each element of the input sequence.
func Sum[Source any, Result constraints.Integer | constraints.Float](source Enumerator[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	if selector == nil {
		return 0, ErrNilSelector
	}
	r, _ := sumPrim(source, selector)
	return r, nil
}

// SumMust is like Sum but panics in case of error.
func SumMust[Source any, Result constraints.Integer | constraints.Float](source Enumerator[Source], selector func(Source) Result) Result {
	r, err := Sum(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// Average computes the average of a sequence of constraints.Integer values that are obtained
// by invoking a transform function on each element of the input sequence.
func Average[Source any, Result constraints.Integer | constraints.Float](source Enumerator[Source], selector func(Source) Result) (float64, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	if selector == nil {
		return 0, ErrNilSelector
	}
	sum, count := sumPrim(source, selector)
	if count == 0 {
		return 0, ErrEmptySource
	}
	return (float64(sum) / float64(count)), nil
}

// AverageMust is like Average but panics in case of error.
func AverageMust[Source any, Result constraints.Integer | constraints.Float](source Enumerator[Source], selector func(Source) Result) float64 {
	r, err := Average(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}
