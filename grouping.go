//go:build go1.18

package go2linq

import (
	"fmt"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq/Grouping.cs
// https://docs.microsoft.com/dotnet/api/system.linq.igrouping-2

// Grouping represents a collection of objects that have a common key.
type Grouping[Key, Element any] struct {
	key    Key
	values []Element
}

// Key gets the key of the Grouping.
func (gr *Grouping[Key, Element]) Key() Key {
	// https://docs.microsoft.com/dotnet/api/system.linq.igrouping-2.key
	return gr.key
}

// GetEnumerator returns an enumerator that iterates through a collection.
func (gr *Grouping[Key, Element]) GetEnumerator() Enumerator[Element] {
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerable.getenumerator
	return NewOnSlice(gr.values...)
}

// String implements the fmt.Stringer interface.
func (gr *Grouping[Key, Element]) String() string {
	return fmt.Sprintf("%v: %v", gr.key, gr.values)
}
