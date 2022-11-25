package go2linq

import (
	"fmt"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq/Grouping.cs
// https://docs.microsoft.com/dotnet/api/system.linq.igrouping-2

// Grouping represents a collection of objects that have a common key.
// (https://docs.microsoft.com/dotnet/api/system.linq.igrouping-2)
type Grouping[Key, Element any] struct {
	key    Key
	values []Element
}

// Key gets the key of the Grouping.
// (https://docs.microsoft.com/dotnet/api/system.linq.igrouping-2.key)
func (gr *Grouping[Key, Element]) Key() Key {
	return gr.key
}

// GetEnumerator returns an enumerator that iterates through the Grouping's collection.
// GetEnumerator implements the Enumerable interface.
func (gr *Grouping[Key, Element]) GetEnumerator() Enumerator[Element] {
	return newEnrSlice(gr.values...)
}

// Slice returns an slice containing the Grouping's collection.
func (gr *Grouping[Key, Element]) Slice() []Element {
	return gr.values
}

// String implements the fmt.Stringer interface.
func (gr *Grouping[Key, Element]) String() string {
	return fmt.Sprintf("%v: %v", gr.key, gr.values)
}
