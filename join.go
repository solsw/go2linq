package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Join] correlates the elements of two sequences based on matching keys.
// [generichelper.DeepEqual] is used to compare keys.
// 'inner' is enumerated on the first iteration over the result.
//
// [Join]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.join
func Join[Outer, Inner, Key, Result any](outer iter.Seq[Outer], inner iter.Seq[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) (iter.Seq[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return JoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, generichelper.DeepEqual[Key])
}

// [JoinEq] correlates the elements of two sequences based on matching keys.
// 'equal' is used to compare keys.
// 'inner' is enumerated on the first iteration over the result.
//
// Similar to the keys equality functionality may be achieved using appropriate key selectors.
// See [TestJoinEqMust_CustomComparer] test for usage of case insensitive string keys.
//
// [JoinEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.join
func JoinEq[Outer, Inner, Key, Result any](outer iter.Seq[Outer], inner iter.Seq[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, equal func(Key, Key) bool) (iter.Seq[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	return func(yield func(Result) bool) {
		},
		nil
}
