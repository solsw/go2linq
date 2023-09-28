package go2linq

import (
	"github.com/solsw/errorhelper"
	"golang.org/x/exp/constraints"
)

// Reimplementing LINQ to Objects: Part 28 – Sum
// https://codeblog.jonskeet.uk/2011/01/08/reimplementing-linq-to-objects-part-28-sum/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sum

// Reimplementing LINQ to Objects: Part 30 – Average
// https://codeblog.jonskeet.uk/2011/01/10/reimplementing-linq-to-objects-part-30-average/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average

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

// [Sum] computes the sum of a sequence of [constraints.Integer] or [constraints.Float] values.
//
// [Sum]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sum
func Sum[Source constraints.Integer | constraints.Float](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	return SumSel(source, Identity[Source])
}

// SumMust is like [Sum] but panics in case of error.
func SumMust[Source constraints.Integer | constraints.Float](source Enumerable[Source]) Source {
	return errorhelper.Must(Sum(source))
}

// [SumSel] computes the sum of a sequence of [constraints.Integer] or [constraints.Float] values
// that are obtained by invoking a transform function on each element of the input sequence.
//
// [SumSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sum
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

// SumSelMust is like [SumSel] but panics in case of error.
func SumSelMust[Source any, Result constraints.Integer | constraints.Float](source Enumerable[Source],
	selector func(Source) Result) Result {
	return errorhelper.Must(SumSel(source, selector))
}

// [Average] computes the average of a sequence of [constraints.Integer] or [constraints.Float] values.
//
// [Average]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average
func Average[Source constraints.Integer | constraints.Float](source Enumerable[Source]) (float64, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	return AverageSel(source, Identity[Source])
}

// AverageMust is like [Average] but panics in case of error.
func AverageMust[Source constraints.Integer | constraints.Float](source Enumerable[Source]) float64 {
	return errorhelper.Must(Average(source))
}

// [AverageSel] computes the average of a sequence of [constraints.Integer] or [constraints.Float]
// values that are obtained by invoking a transform function on each element of the input sequence.
//
// [AverageSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average
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

// AverageSelMust is like [AverageSel] but panics in case of error.
func AverageSelMust[Source any, Result constraints.Integer | constraints.Float](source Enumerable[Source],
	selector func(Source) Result) float64 {
	return errorhelper.Must(AverageSel(source, selector))
}
