package go2linq

import (
	"github.com/solsw/collate"
	"github.com/solsw/errorhelper"
)

// Reimplementing LINQ to Objects: Part 34 - SequenceEqual
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-34-sequenceequal/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

// [SequenceEqual] determines whether two sequences are equal by comparing the elements using [collate.DeepEqualer].
//
// [SequenceEqual]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func SequenceEqual[Source any](first, second Enumerable[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	return SequenceEqualEq(first, second, nil)
}

// SequenceEqualMust is like [SequenceEqual] but panics in case of error.
func SequenceEqualMust[Source any](first, second Enumerable[Source]) bool {
	return errorhelper.Must(SequenceEqual(first, second))
}

// [SequenceEqualEq] determines whether two sequences are equal by comparing their elements using a specified equaler.
// If 'equaler' is nil, [collate.DeepEqualer] is used.
//
// [SequenceEqualEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func SequenceEqualEq[Source any](first, second Enumerable[Source], equaler collate.Equaler[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	counter1, ok1 := first.(Counter)
	if ok1 {
		counter2, ok2 := second.(Counter)
		if ok2 && (counter1.Count() != counter2.Count()) {
			return false, nil
		}
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Source]{}
	}
	enr1 := first.GetEnumerator()
	enr2 := second.GetEnumerator()
	for enr1.MoveNext() {
		if !enr2.MoveNext() {
			return false, nil
		}
		if !equaler.Equal(enr1.Current(), enr2.Current()) {
			return false, nil
		}
	}
	if enr2.MoveNext() {
		return false, nil
	}
	return true, nil
}

// SequenceEqualEqMust is like [SequenceEqualEq] but panics in case of error.
func SequenceEqualEqMust[Source any](first, second Enumerable[Source], equaler collate.Equaler[Source]) bool {
	return errorhelper.Must(SequenceEqualEq(first, second, equaler))
}
