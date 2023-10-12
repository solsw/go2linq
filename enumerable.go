package go2linq

import (
	"context"
	"fmt"

	"github.com/solsw/generichelper"
	"golang.org/x/sync/errgroup"
)

// https://learn.microsoft.com/dotnet/api/system.collections.generic.ienumerable-1

// [Enumerable] exposes the enumerator, which supports a simple iteration over a collection of a specified type.
//
// [Enumerable]: https://learn.microsoft.com/dotnet/api/system.collections.generic.ienumerable-1
type Enumerable[T any] interface {

	// [GetEnumerator] returns [Enumerator] that iterates through the collection.
	//
	// [GetEnumerator]: https://learn.microsoft.com/dotnet/api/system.collections.generic.ienumerable-1.getenumerator
	GetEnumerator() Enumerator[T]
}

// enmrbl implements the [Enumerable] interface.
type enmrbl[T any] struct {
	// must return a new instance of [Enumerator] on each call
	getEnr func() Enumerator[T]
}

// GetEnumerator implements the [Enumerable.GetEnumerator] method.
func (en enmrbl[T]) GetEnumerator() Enumerator[T] {
	if en.getEnr == nil {
		return enrFunc[T]{}
	}
	return en.getEnr()
}

// OnFactory creates a new [Enumerable] based on the provided [Enumerator] factory.
func OnFactory[T any](factory func() Enumerator[T]) Enumerable[T] {
	return enmrbl[T]{getEnr: factory}
}

// OnMap creates a new [Enumerable] based on the provided [map].
// Deprecated: OnMap is for backward compatibility. Use [NewEnMapEn] instead.
//
// [map]: https://go.dev/ref/spec#Map_types
func OnMap[Key comparable, Element any](m map[Key]Element) Enumerable[generichelper.Tuple2[Key, Element]] {
	return NewEnMap(m)
}

// OnChan creates a new [Enumerable] based on the provided [channel].
// Deprecated: OnChan is for backward compatibility. Use [NewEnChanEn] instead.
//
// [channel]: https://go.dev/ref/spec#Channel_types
func OnChan[T any](ch <-chan T) Enumerable[T] {
	return NewEnChan(ch)
}

// ToStringFmt returns string representation of a sequence:
//   - if 'en' is nil, empty string is returned;
//   - if 'en' or underlying Enumerator implements fmt.Stringer, it is used;
//   - if 'T' implements fmt.Stringer, it is used to convert each element to string;
//   - 'sep' is inserted between elements;
//   - 'lrim' and 'rrim' surround each element;
//   - 'ledge' and 'redge' surround the whole string.
func ToStringFmt[T any](en Enumerable[T], sep, lrim, rrim, ledge, redge string) string {
	if en == nil {
		return ""
	}
	if stringer, ok := en.(fmt.Stringer); ok {
		return stringer.String()
	}
	return enrToStringFmt(en.GetEnumerator(), sep, lrim, rrim, ledge, redge)
}

// ToStringDef returns string representation of a sequence using default formatting.
// If 'en' is nil, empty string is returned.
func ToStringDef[T any](en Enumerable[T]) string {
	if en == nil {
		return ""
	}
	return ToStringFmt(en, " ", "", "", "[", "]")
}

// ToEnString converts a sequence to [Enumerable[string]].
// If 'en' is nil, nil is returned.
func ToEnString[T any](en Enumerable[T]) Enumerable[string] {
	if en == nil {
		return nil
	}
	return OnFactory(
		func() Enumerator[string] {
			return enrToStringEnr(en.GetEnumerator())
		},
	)
}

// ToStrings returns a sequence contents as a slice of strings.
// If 'en' is nil, nil is returned.
func ToStrings[T any](en Enumerable[T]) []string {
	if en == nil {
		return nil
	}
	return enrToStrings(en.GetEnumerator())
}

// ForEach sequentially performs a specified 'action' on each element of the sequence starting from the current.
// If 'ctx' is canceled or 'action' returns non-nil error,
// the operation is canceled and corresponding error is returned.
func ForEach[T any](ctx context.Context, en Enumerable[T], action func(T) error) error {
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
			if err := action(enr.Current()); err != nil {
				return err
			}
		}
	}
	return nil
}

// ForEachConcurrent concurrently performs a specified 'action' on each element of the sequence starting from the current.
// If 'ctx' is canceled or 'action' returns non-nil error,
// the operation is canceled and corresponding error is returned.
func ForEachConcurrent[T any](ctx context.Context, en Enumerable[T], action func(T) error) error {
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
				if err := action(c); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return g.Wait()
}
