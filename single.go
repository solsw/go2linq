package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Single] returns the only element of a sequence and returns an error if there is not exactly one element in the sequence.
//
// [Single]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
func Single[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	next, stop := iter.Pull(source)
	defer stop()
	s, ok := next()
	if !ok {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	_, ok = next()
	if ok {
		return generichelper.ZeroValue[Source](), ErrMultipleElements
	}
	return s, nil
}

// [SinglePred] returns the only element of a sequence that satisfies a specified condition.
//
// [SinglePred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
func SinglePred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	empty := true
	found := false
	var r Source
	for s := range source {
		empty = false
		if predicate(s) {
			if found {
				return generichelper.ZeroValue[Source](), ErrMultipleMatch
			}
			found = true
			r = s
		}
	}
	if empty {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	if !found {
		return generichelper.ZeroValue[Source](), ErrNoMatch
	}
	return r, nil
}

// [SingleOrDefault] returns the only element of a sequence or a [zero value] if the sequence is empty.
//
// [SingleOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func SingleOrDefault[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	r, err := Single(source)
	if err != nil {
		if err == ErrMultipleElements {
			return generichelper.ZeroValue[Source](), ErrMultipleElements
		}
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}

// [SingleOrDefaultPred] returns the only element of a sequence that satisfies a specified condition
// or a [zero value] if no such element exists.
//
// [SingleOrDefaultPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func SingleOrDefaultPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	r, err := SinglePred(source, predicate)
	if err != nil {
		if err == ErrMultipleMatch {
			return generichelper.ZeroValue[Source](), ErrMultipleMatch
		}
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}
