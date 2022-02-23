//go:build go1.18

package go2linq

import (
	"golang.org/x/exp/constraints"
)

// Reimplementing LINQ to Objects: Part 28 – Sum
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-28-sum/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

// Reimplementing LINQ to Objects: Part 30 – Average
// https://codeblog.jonskeet.uk/2011/01/10/reimplementing-linq-to-objects-part-30-average/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func sumPrim[Source any, Result constraints.Integer | constraints.Float](source Enumerable[Source],
	selector func(Source) Result) (Result, int) {
	enr := source.GetEnumerator()
	var sum Result = 0
	count := 0
	for enr.MoveNext() {
		sum += selector(enr.Current())
		count++
	}
	return sum, count
}

// Sum computes the sum of a sequence of constraints.Integer or constraints.Float values.
func Sum[Source constraints.Integer | constraints.Float](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	return SumSel(source, Identity[Source])
}

// SumMust is like Sum but panics in case of error.
func SumMust[Source constraints.Integer | constraints.Float](source Enumerable[Source]) Source {
	r, err := Sum(source)
	if err != nil {
		panic(err)
	}
	return r
}

// SumSel computes the sum of a sequence of constraints.Integer or constraints.Float values that are obtained
// by invoking a transform function on each element of the input sequence.
func SumSel[Source any, Result constraints.Integer | constraints.Float](source Enumerable[Source],
	selector func(Source) Result) (Result, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	if selector == nil {
		return 0, ErrNilSelector
	}
	r, _ := sumPrim(source, selector)
	return r, nil
}

// SumSelMust is like SumSel but panics in case of error.
func SumSelMust[Source any, Result constraints.Integer | constraints.Float](source Enumerable[Source],
	selector func(Source) Result) Result {
	r, err := SumSel(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// Average computes the average of a sequence of constraints.Integer or constraints.Float values.
func Average[Source constraints.Integer | constraints.Float](source Enumerable[Source]) (float64, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	return AverageSel(source, Identity[Source])
}

// AverageMust is like Average but panics in case of error.
func AverageMust[Source constraints.Integer | constraints.Float](source Enumerable[Source]) float64 {
	r, err := Average(source)
	if err != nil {
		panic(err)
	}
	return r
}

// AverageSel computes the average of a sequence of constraints.Integer or constraints.Float values that are obtained
// by invoking a transform function on each element of the input sequence.
func AverageSel[Source any, Result constraints.Integer | constraints.Float](source Enumerable[Source],
	selector func(Source) Result) (float64, error) {
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

// AverageSelMust is like AverageSel but panics in case of error.
func AverageSelMust[Source any, Result constraints.Integer | constraints.Float](source Enumerable[Source],
	selector func(Source) Result) float64 {
	r, err := AverageSel(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}
