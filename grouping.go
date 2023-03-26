package go2linq

import (
	"fmt"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq/Grouping.cs
// https://learn.microsoft.com/dotnet/api/system.linq.igrouping-2

// [Grouping] represents a collection of objects that have a common key.
//
// [Grouping]: https://learn.microsoft.com/dotnet/api/system.linq.igrouping-2
type Grouping[Key, Element any] struct {
	key    Key
	values []Element
}

// [Key] gets the key of the [Grouping].
//
// [Key]: https://learn.microsoft.com/dotnet/api/system.linq.igrouping-2.key
func (gr *Grouping[Key, Element]) Key() Key {
	return gr.key
}

// Values returns the values of the [Grouping].
func (gr *Grouping[Key, Element]) Values() []Element {
	return gr.values
}

// Count returns the number of elements in the [Grouping].
func (gr *Grouping[Key, Element]) Count() int {
	return len(gr.values)
}

// GetEnumerator returns an [Enumerator] that iterates through the [Grouping]'s collection.
//
// GetEnumerator implements the [Enumerable] interface.
func (gr *Grouping[Key, Element]) GetEnumerator() Enumerator[Element] {
	return newEnrSlice(gr.values...)
}

// String implements the [fmt.Stringer] interface.
func (gr *Grouping[Key, Element]) String() string {
	return fmt.Sprintf("%v: %v", gr.key, gr.values)
}
