package go2linq

import (
	"errors"
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [Single] returns the only element of a sequence and returns an error if there is not exactly one element in the sequence.
//
// [Single]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
func Single[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilSource)
	}
	next, stop := iter.Pull(source)
	defer stop()
	s, ok := next()
	if !ok {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrEmptySource)
	}
	_, ok = next()
	if ok {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrMultipleElements)
	}
	return s, nil
}

// [SinglePred] returns the only element of a sequence that satisfies a specified condition.
//
// [SinglePred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
func SinglePred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilSource)
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilPredicate)
	}
	empty := true
	found := false
	var r Source
	for s := range source {
		empty = false
		if predicate(s) {
			if found {
				return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrMultipleMatch)
			}
			found = true
			r = s
		}
	}
	if empty {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrEmptySource)
	}
	if !found {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNoMatch)
	}
	return r, nil
}

// [SingleOrDefault] returns the only element of a [sequence]
// or a specified default value if the [sequence] is empty.
//
// [SingleOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [sequence]: https://pkg.go.dev/iter#Seq
func SingleOrDefault[Source any](source iter.Seq[Source], defaultValue Source) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilSource)
	}
	r, err := Single(source)
	if err != nil {
		if errors.Is(err, ErrMultipleElements) {
			return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrMultipleElements)
		}
		// the sequence is empty, return the default value
		return defaultValue, nil
	}
	return r, nil
}

// [SingleOrZero] returns the only element of a [sequence] or a [zero value] if the [sequence] is empty.
//
// [SingleOrZero]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [sequence]: https://pkg.go.dev/iter#Seq
// [zero value]: https://go.dev/ref/spec#The_zero_value
func SingleOrZero[Source any](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilSource)
	}
	return SingleOrDefault(source, generichelper.ZeroValue[Source]())
}

// [SingleOrDefaultPred] returns the only element of a [sequence] that satisfies a specified condition
// or a specified default value if no such element exists.
//
// [SingleOrDefaultPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [sequence]: https://pkg.go.dev/iter#Seq
func SingleOrDefaultPred[Source any](source iter.Seq[Source], predicate func(Source) bool, defaultValue Source) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilSource)
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilPredicate)
	}
	r, err := SinglePred(source, predicate)
	if err != nil {
		if errors.Is(err, ErrMultipleMatch) {
			return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrMultipleMatch)
		}
		// sequence is empty or no match - return default value
		return defaultValue, nil
	}
	return r, nil
}

// [SingleOrZeroPred] returns the only element of a [sequence] that satisfies a specified condition
// or a [zero value] if no such element exists.
//
// [SingleOrZeroPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [sequence]: https://pkg.go.dev/iter#Seq
// [zero value]: https://go.dev/ref/spec#The_zero_value
func SingleOrZeroPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilSource)
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), errorhelper.CallerError(ErrNilPredicate)
	}
	return SingleOrDefaultPred(source, predicate, generichelper.ZeroValue[Source]())
}
