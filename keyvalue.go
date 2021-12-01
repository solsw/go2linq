//go:build go1.18

package go2linq

// KeyValue represents key, value pair of Dictionary.
type KeyValue[Key, Value any] struct {
	key   Key
	value Value
}

// Key returns the key of the KeyValue.
func (kv *KeyValue[Key, Value]) Key() Key {
	return kv.key
}

// Value returns the value of the KeyValue.
func (kv *KeyValue[Key, Value]) Value() Value {
	return kv.value
}
