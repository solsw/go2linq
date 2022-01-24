//go:build go1.18

package go2linq

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

// https://docs.microsoft.com/dotnet/api/system.collections.generic.ienumerable-1

// Enumerable exposes the enumerator, which supports a simple iteration over a collection of a specified type.
type Enumerable[T any] interface {
	// GetEnumerator returns an enumerator that iterates through the collection.
	GetEnumerator() Enumerator[T]
}

// enmrbl implements the Enumerable interface.
type enmrbl[T any] struct {
	// must return a new instance of Enumerator on each call
	getEnr func() Enumerator[T]
}

// GetEnumerator implements the Enumerable.GetEnumerator interface method.
func (en enmrbl[T]) GetEnumerator() Enumerator[T] {
	if en.getEnr == nil {
		return enrFunc[T]{}
	}
	return en.getEnr()
}

// EnOnFactory creates a new Enumerable based on the provided Enumerator factory.
func EnOnFactory[T any](factory func() Enumerator[T]) Enumerable[T] {
	return enmrbl[T]{
		getEnr: factory,
	}
}

// // EnOnSlice creates a new Enumerable based on the provided slice.
// func EnOnSlice[T any](slice ...T) Enumerable[T] {
// 	return EnOnFactory[T](
// 		func() Enumerator[T] {
// 			return newEnrSlice[T](slice...)
// 		},
// 	)
// }

// EnOnMap creates a new Enumerable based on the provided map.
func EnOnMap[Key comparable, Element any](m map[Key]Element) Enumerable[KeyElement[Key, Element]] {
	r := make([]KeyElement[Key, Element], 0, len(m))
	for k, e := range m {
		r = append(r, KeyElement[Key, Element]{k, e})
	}
	return NewEnSlice(r...)
}

// EnOnChan creates a new Enumerable based on the provided channel.
func EnOnChan[T any](ch <-chan T) Enumerable[T] {
	return EnOnFactory[T](
		func() Enumerator[T] {
			return newEnrChan[T](ch)
		},
	)
}

// EnToSlice creates a slice from an Enumerable. EnToSlice returns nil if 'en' is nil.
func EnToSlice[T any](en Enumerable[T]) []T {
	if en == nil {
		return nil
	}
	return enrToSlice(en.GetEnumerator())
}

// EnToSliceErr is like EnToSlice but:
//
// - if the underlying EnToSlice panics with an error, the error is recovered and returned;
//
// - if the underlying EnToSlice panics with a string, the string is recovered, wrapped into an error and returned.
func EnToSliceErr[T any](en Enumerable[T]) (res []T, err error) {
	defer func() {
		catchErrStr[[]T](recover(), &res, &err)
	}()
	return EnToSlice[T](en), nil
}

// EnToString returns string representation of a sequence.
func EnToString[T any](en Enumerable[T]) string {
	if s, ok := en.(fmt.Stringer); ok {
		return s.String()
	}
	return enrToString(en.GetEnumerator())
}

// EnToStringEn converts a sequence to Enumerable[string].
func EnToStringEn[T any](en Enumerable[T]) Enumerable[string] {
	return EnOnFactory(
		func() Enumerator[string] {
			return enrToStringEnr(en.GetEnumerator())
		},
	)
}

// EnToStrings returns a sequence contents as a slice of strings.
func EnToStrings[T any](en Enumerable[T]) []string {
	return enrToStrings(en.GetEnumerator())
}

// ForEachEn sequentially performs the specified action on each element of the sequence starting from the current.
// 'ctx' may be used to cancel the operation in progress.
func ForEachEn[T any](ctx context.Context, en Enumerable[T], action func(context.Context, T) error) error {
	if en == nil {
		return ErrNilSource
	}
	if action == nil {
		return ErrNilAction
	}
	enr := en.GetEnumerator()
	for enr.MoveNext() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := action(ctx, enr.Current()); err != nil {
				return err
			}
		}
	}
	return nil
}

// ForEachEnConcurrent concurrently performs the specified action on each element of the sequence starting from the current.
// 'ctx' may be used to cancel the operation in progress.
func ForEachEnConcurrent[T any](ctx context.Context, en Enumerable[T], action func(context.Context, T) error) error {
	if en == nil {
		return ErrNilSource
	}
	if action == nil {
		return ErrNilAction
	}
	enr := en.GetEnumerator()
	g := new(errgroup.Group)
	for enr.MoveNext() {
		c := enr.Current()
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if err := action(ctx, c); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return g.Wait()
}
