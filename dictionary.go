//go:build go1.18

package go2linq

// Dictionary represents map[Key]Value.
type Dictionary[Key comparable, Value any] map[Key]Value

// GetEnumerator returns an enumerator that iterates through a Dictionary.
func (d Dictionary[Key, Value]) GetEnumerator() Enumerator[KeyValue[Key, Value]] {
	r := make([]KeyValue[Key, Value], 0, len(d))
	for k, v := range d {
		r = append(r, KeyValue[Key, Value]{k, v})
	}
	return NewOnSlice(r...)
}

// AsDictionary converts a sequence of KeyValues to Dictionary.
func AsDictionary[Key comparable, Value any](en Enumerator[KeyValue[Key, Value]]) Dictionary[Key, Value] {
	r := make(Dictionary[Key, Value])
	for en.MoveNext() {
		kv := en.Current()
		r[kv.key] = kv.value
	}
	return r
}
