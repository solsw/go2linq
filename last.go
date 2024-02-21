package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Last] returns the last element of a sequence.
//
// [Last]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func Last[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	empty := true
	var res Source
	for s := range source {
		empty = false
		res = s
	}
	if empty {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return res, nil
}

// [LastPred] returns the last element of a sequence that satisfies a specified condition.
//
// [LastPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func LastPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	empty := true
	found := false
	var res Source
	for s := range source {
		empty = false
		if predicate(s) {
			found = true
			res = s
		}
	}
	if empty {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	if !found {
		return generichelper.ZeroValue[Source](), ErrNoMatch
	}
	return res, nil
}

// [LastOrDefault] returns the last element of a sequence or a [zero value] if the sequence contains no elements.
//
// [LastOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func LastOrDefault[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	res, err := Last(source)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return res, nil
}

// [LastOrDefaultPred] returns the last element of a sequence that satisfies a condition
// or a [zero value] if no such element is found.
//
// [LastOrDefaultPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func LastOrDefaultPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	res, err := LastPred(source, predicate)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return res, nil
}
