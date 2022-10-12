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

// GetEnumerator implements the Enumerable interface.
func (en *EnMap[Key, Element]) GetEnumerator() Enumerator[KeyElement[Key, Element]] {
	return (*EnSlice[KeyElement[Key, Element]])(en).GetEnumerator()
}

// Count implements the Counter interface.
func (en *EnMap[Key, Element]) Count() int {
	return len(*en)
}
