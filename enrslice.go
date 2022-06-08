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

func (enr *enrSlice[T]) item(i int) T {
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.current#remarks
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.list-1.item#exceptions
	if !(0 <= i && i < len(enr.elel)) {
		return ZeroValue[T]()
	}
	return enr.elel[i]
}

// MoveNext implements the Enumerator.MoveNext method.
func (enr *enrSlice[T]) MoveNext() bool {
	if enr.indx >= len(enr.elel) {
		return false
	}
	enr.indx++
	return true
}

// Current implements the Enumerator.Current method.
func (enr *enrSlice[T]) Current() T {
	return enr.item(enr.indx - 1)
}

// Reset implements the Enumerator.Reset method.
func (enr *enrSlice[T]) Reset() {
	enr.indx = 0
}
