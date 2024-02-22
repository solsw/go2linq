package go2linq

import (
	"iter"

	"golang.org/x/exp/constraints"
)

func sumPrim[Source any, Result constraints.Integer | constraints.Float](source iter.Seq[Source],
	selector func(Source) Result) (Result, int) {
	var sum Result = 0
	count := 0
	for s := range source {
		sum += selector(s)
		count++
	}
	return sum, count
}

// [Sum] computes the sum of a sequence of [constraints.Integer] or [constraints.Float] values.
//
// [Sum]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sum
func Sum[Source constraints.Integer | constraints.Float](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	return SumSel(source, Identity[Source])
}

// [SumSel] computes the sum of a sequence of [constraints.Integer] or [constraints.Float] values
// that are obtained by invoking a transform function on each element of the input sequence.
//
// [SumSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sum
func SumSel[Source any, Result constraints.Integer | constraints.Float](source iter.Seq[Source],
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

// [Average] computes the average of a sequence of [constraints.Integer] or [constraints.Float] values.
//
// [Average]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average
func Average[Source constraints.Integer | constraints.Float](source iter.Seq[Source]) (float64, error) {
	if source == nil {
		return 0, ErrNilSource
	}
	return AverageSel(source, Identity[Source])
}

// [AverageSel] computes the average of a sequence of [constraints.Integer] or [constraints.Float]
// values that are obtained by invoking a transform function on each element of the input sequence.
//
// [AverageSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average
func AverageSel[Source any, Result constraints.Integer | constraints.Float](source iter.Seq[Source],
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
