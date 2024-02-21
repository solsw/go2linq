package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [First] returns the first element of a sequence.
//
// [First]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.first
func First[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	for s := range source {
		return s, nil
	}
	return generichelper.ZeroValue[Source](), ErrEmptySource
}

// [FirstPred] returns the first element in a sequence that satisfies a specified condition.
//
// [FirstPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.first
func FirstPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	empty := true
	for s := range source {
		empty = false
		if predicate(s) {
			return s, nil
		}
	}
	if empty {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return generichelper.ZeroValue[Source](), ErrNoMatch
}

// [FirstOrDefault] returns the first element of a sequence, or a [zero value] if the sequence contains no elements.
//
// [FirstOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func FirstOrDefault[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	r, err := First(source)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}

// [FirstOrDefaultPred] returns the first element of the sequence that satisfies a condition
// or a [zero value] if no such element is found.
//
// [FirstOrDefaultPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func FirstOrDefaultPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	r, err := FirstPred(source, predicate)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}
