//go:build go1.18

package go2linq

// KeyElement represents key, element pair of map.
type KeyElement[Key, Element any] struct {
	key     Key
	element Element
}

// Key returns the key of the KeyElement.
func (ke *KeyElement[Key, Element]) Key() Key {
	return ke.key
}

// Element returns the element of the KeyElement.
func (ke *KeyElement[Key, Element]) Element() Element {
	return ke.element
}
