//go:build go1.18

package go2linq

// EnMap is an Enumerable implementation based on a map.
type EnMap[Key comparable, Element any] EnSlice[KeyElement[Key, Element]]

// NewEnMap creates a new EnMap with the specified map as contents.
func NewEnMap[Key comparable, Element any](m map[Key]Element) Enumerable[KeyElement[Key, Element]] {
	sl := make([]KeyElement[Key, Element], 0, len(m))
	for k, e := range m {
		sl = append(sl, KeyElement[Key, Element]{k, e})
	}
	return NewEnSlice(sl...)
}

func (en *EnMap[Key, Element]) enSlice() *EnSlice[KeyElement[Key, Element]] {
	return (*EnSlice[KeyElement[Key, Element]])(en)
}

// GetEnumerator implements the Enumerable interface.
func (en *EnMap[Key, Element]) GetEnumerator() Enumerator[KeyElement[Key, Element]] {
	return en.enSlice().GetEnumerator()
}

// Count implements the Counter interface.
func (en *EnMap[Key, Element]) Count() int {
	return en.enSlice().Count()
}

// Item implements the Itemer interface.
func (en *EnMap[Key, Element]) Item(i int) KeyElement[Key, Element] {
	return en.enSlice().Item(i)
}

// Slice implements the Slicer interface.
func (en *EnMap[Key, Element]) Slice() []KeyElement[Key, Element] {
	return en.enSlice().Slice()
}
