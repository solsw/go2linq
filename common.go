//go:build go1.18

package go2linq

import (
	"reflect"
)

type (
	// Counter is the interface that wraps the Count method.
	Counter interface {
		// Count returns the number of elements in a sequence.
		Count() int
	}

	// Itemer is the interface that wraps the Item method.
	Itemer[T any] interface {
		// Item returns the element of a sequence at the specified index.
		Item(int) T
	}

	// Slicer is the interface that wraps the Slice method.
	Slicer[T any] interface {
		// Slice returns the sequence contents as a slice.
		Slice() []T
	}
)

// ZeroValue returns T's zero value.
func ZeroValue[T any]() T {
	var t0 T
	return t0
}

// Identity is a selector that projects the element into itself.
func Identity[T any](el T) T {
	return el
}

// DeepEqual is an equaler that is a generic wrapper for reflect.DeepEqual.
func DeepEqual[T any](x, y T) bool {
	return reflect.DeepEqual(x, y)
}
