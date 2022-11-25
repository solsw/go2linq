package go2linq

import (
	"fmt"
	"reflect"
	"strings"
)

// Enumerator supports a simple iteration over a generic sequence. T - the type of objects to enumerate
// (https://docs.microsoft.com/dotnet/api/system.collections.generic.ienumerator-1).
type Enumerator[T any] interface {

	// MoveNext advances the enumerator to the next element of the sequence
	// (https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.movenext).
	MoveNext() bool

	// Current returns the element in the sequence at the current position of the enumerator
	// (https://docs.microsoft.com/dotnet/api/system.collections.generic.ienumerator-1.current).
	Current() T

	// Reset sets the enumerator to its initial position, which is before the first element in the sequence
	// (https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.reset,
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.reset#remarks,
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator#remarks).
	Reset()
}

// enrToSlice creates a slice from Enumerator (returns nil if 'enr' is nil).
func enrToSlice[T any](enr Enumerator[T]) []T {
	if enr == nil {
		return nil
	}
	if slicer, ok := enr.(Slicer[T]); ok {
		return slicer.Slice()
	}
	r := []T{}
	for enr.MoveNext() {
		r = append(r, enr.Current())
	}
	return r
}

func enrToStringFmt[T any](enr Enumerator[T], sep, lrim, rrim, ledge, redge string) string {
	if stringer, ok := enr.(fmt.Stringer); ok {
		return stringer.String()
	}
	isStringer := typeIsStringer[T]()
	var b strings.Builder
	for enr.MoveNext() {
		if b.Len() > 0 {
			b.WriteString(sep)
		}
		b.WriteString(lrim + asStringPrim(enr.Current(), isStringer) + rrim)
	}
	return ledge + b.String() + redge
}

// enrToStringEnr converts the sequence to Enumerator[string].
func enrToStringEnr[T any](enr Enumerator[T]) Enumerator[string] {
	isStringer := typeIsStringer[T]()
	return enrFunc[string]{
		mvNxt: func() bool { return enr.MoveNext() },
		crrnt: func() string { return asStringPrim(enr.Current(), isStringer) },
		rst:   func() { enr.Reset() },
	}
}

// enrToStrings returns the sequence contents as a slice of strings.
func enrToStrings[T any](enr Enumerator[T]) []string {
	return enrToSlice(enrToStringEnr(enr))
}

// CloneEmpty creates a new empty Enumerator of the same type as 'enr'.
func CloneEmpty[T any](enr Enumerator[T]) Enumerator[T] {
	// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
	t := reflect.TypeOf(enr)
	var v reflect.Value
	if t.Kind() == reflect.Ptr {
		v = reflect.New(t.Elem())
	} else {
		v = reflect.Zero(t)
	}
	i := v.Interface()
	return i.(Enumerator[T])
}
