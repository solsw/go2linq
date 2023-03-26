package go2linq

type (
	// Counter is the interface that wraps the Count method.
	Counter interface {
		// Count returns the number of elements contained in the sequence.
		Count() int
	}

	// Itemer is the interface that wraps the Item method.
	Itemer[T any] interface {
		// Item returns the element at the specified index in the sequence.
		Item(int) T
	}

	// Slicer is the interface that wraps the Slice method.
	Slicer[T any] interface {
		// Slice returns the sequence contents as a slice.
		Slice() []T
	}
)

// Identity is a selector that projects the element into itself.
func Identity[T any](el T) T {
	return el
}
