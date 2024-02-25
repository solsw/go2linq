package go2linq

import "iter"

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

// Values returns a sequence of values in the Grouping.
func (gr *Grouping[Key, Element]) Values() iter.Seq[Element] {
	return SliceAll(gr.values)
}
