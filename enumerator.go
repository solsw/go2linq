//go:build go1.18

package go2linq

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	//	"golang.org/x/sync/errgroup"
)

// Enumerator supports a simple iteration over a generic sequence
// (https://docs.microsoft.com/dotnet/api/system.collections.generic.ienumerator-1).
// T - the type of objects to enumerate.
type Enumerator[T any] interface {

	// MoveNext advances the enumerator to the next element of the sequence.
	MoveNext() bool

	// Current returns the element in the sequence at the current position of the enumerator.
	Current() T

	// Reset sets the enumerator to its initial position, which is before the first element in the sequence
	// (https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.reset#remarks,
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator#remarks).
	Reset()
}

// Slice creates a slice from an Enumerator.
// Slice returns nil if 'en' is nil.
func Slice[T any](en Enumerator[T]) []T {
	if en == nil {
		return nil
	}
	if slicer, ok := en.(Slicer[T]); ok {
		return slicer.Slice()
	}
	var r []T
	for en.MoveNext() {
		r = append(r, en.Current())
	}
	return r
}

// SliceErr is like Slice but if the underlying Slice panics with error, this error is recovered and returned.
func SliceErr[T any](en Enumerator[T]) (res []T, err error) {
	defer func() {
		catchErrPanic[[]T](recover(), &res, &err)
	}()
	return Slice[T](en), nil
}

func asStringPrim[T any](t T, isStringer bool) string {
	if isStringer {
		return any(t).(fmt.Stringer).String()
	}
	return fmt.Sprint(t)
}

func typeIsStringer[T any]() bool {
	var t0 T
	var i any = t0
	_, isStringer := i.(fmt.Stringer)
	return isStringer
}

// StringFmt prints formatted sequence:
// elements, each surrounded by 'leftRim' and 'rightRim', are separated by 'separator'.
// If element implements fmt.Stringer it is used to convert element to string,
// otherwise fmt.Sprint is used.
func StringFmt[T any](en Enumerator[T], separator, leftRim, rightRim string) string {
	isStringer := typeIsStringer[T]()
	var b strings.Builder
	for en.MoveNext() {
		if b.Len() > 0 {
			b.WriteString(separator)
		}
		b.WriteString(leftRim + asStringPrim(en.Current(), isStringer) + rightRim)
	}
	return b.String()
}

// String mimics the fmt.Stringer interface.
func String[T any](en Enumerator[T]) string {
	if s, ok := en.(fmt.Stringer); ok {
		return s.String()
	}
	//	return fmt.Sprint(Slice(en))
	return "[" + StringFmt(en, " ", "", "") + "]"
}

// ToStrings converts the sequence to Enumerator[string].
func ToStrings[T any](en Enumerator[T]) Enumerator[string] {
	isStringer := typeIsStringer[T]()
	return OnFunc[string]{
		mvNxt: func() bool { return en.MoveNext() },
		crrnt: func() string { return asStringPrim(en.Current(), isStringer) },
		rst:   func() { en.Reset() },
	}
}

// Strings returns the sequence contents as a slice of strings.
func Strings[T any](en Enumerator[T]) []string {
	return Slice(ToStrings(en))
}

// CloneEmpty creates a new empty Enumerator of the same type as 'en'.
func CloneEmpty[T any](en Enumerator[T]) Enumerator[T] {
	// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
	t := reflect.TypeOf(en)
	var v reflect.Value
	if t.Kind() == reflect.Ptr {
		v = reflect.New(t.Elem())
	} else {
		v = reflect.Zero(t)
	}
	i := v.Interface()
	return i.(Enumerator[T])
}

// ForEach sequentially performs the specified action on each element of the sequence starting from the current.
// 'ctx' may be used to cancel the operation in progress.
func ForEach[T any](ctx context.Context, en Enumerator[T], action Action[T]) error {
	if en == nil {
		return ErrNilSource
	}
	if action == nil {
		return ErrNilAction
	}
	for en.MoveNext() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := action(ctx, en.Current()); err != nil {
				return err
			}
		}
	}
	return nil
}

/*
// ForEachConcurrent concurrently (using errgroup.Group.Go) performs the specified action
// on each element of the sequence starting from the current.
// 'ctx' may be used to cancel the operation in progress.
func ForEachConcurrent[T any](ctx context.Context, en Enumerator[T], action Action[T]) error {
	if en == nil {
		return ErrNilSource
	}
	if action == nil {
		return ErrNilAction
	}
	g := new(errgroup.Group)
	for en.MoveNext() {
		c := en.Current()
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
*/
