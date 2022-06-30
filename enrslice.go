//go:build go1.18

package go2linq

// enrSlice is an Enumerator implementation based on a slice of T.
type enrSlice[T any] struct {
	// indx-1 - index of the current element in elel
	//     ^^ because initially enumerator is positioned before the first element in the collection
	indx int
	// elements of the enrSlice instance
	elel []T
}

// newEnrSlice creates a new enrSlice with the specified contents.
func newEnrSlice[T any](ee ...T) *enrSlice[T] {
	var enr enrSlice[T]
	enr.elel = make([]T, len(ee))
	if len(ee) > 0 {
		copy(enr.elel, ee)
	}
	return &enr
}

// enrSlice_moveNext is used for testing
func enrSlice_moveNext[T any](enr *enrSlice[T]) bool {
	if enr.indx >= len(enr.elel) {
		return false
	}
	enr.indx++
	return true
}

// MoveNext implements the Enumerator.MoveNext method.
func (enr *enrSlice[T]) MoveNext() bool {
	return enrSlice_moveNext(enr)
}

// enrSlice_current is used for testing
func enrSlice_current[T any](enr *enrSlice[T]) T {
	return enr.Item(enr.indx - 1)
}

// Current implements the Enumerator.Current method.
func (enr *enrSlice[T]) Current() T {
	return enrSlice_current(enr)
}

// Reset implements the Enumerator.Reset method.
func (enr *enrSlice[T]) Reset() {
	enr.indx = 0
}

// Count implements the Counter interface.
func (enr *enrSlice[T]) Count() int {
	return len(enr.elel)
}

// Item implements the Itemer interface.
func (enr *enrSlice[T]) Item(i int) T {
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.current#remarks
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.list-1.item#exceptions
	// if i < 0 {
	// 	panic("Enumeration has not started. Call MoveNext.")
	// }
	// if i >= len(enr.elel) {
	// 	panic("Enumeration already finished.")
	// }

	if !(0 <= i && i < len(enr.elel)) {
		return ZeroValue[T]()
	}

	return enr.elel[i]
}

// Slice implements the Slicer interface.
func (enr *enrSlice[T]) Slice() []T {
	return enr.elel
}
