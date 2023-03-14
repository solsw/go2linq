package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v2"
)

// ToLookup creates a go2linq.Lookup from a slice according to a specified key selector function and a key equaler.
// If 'equaler' is nil go2linq.collate.DeepEqualer is used.
// If 'source' is nil or empty, zero value of go2linq.Lookup is returned.
func ToLookup[Source, Key any](source []Source, keySelector func(Source) Key, equaler collate.Equaler[Key]) (*go2linq.Lookup[Key, Source], error) {
	return ToLookupSel(source, keySelector, go2linq.Identity[Source], equaler)
}

// ToLookupSel creates a go2linq.Lookup from a slice according to a specified key selector function,
// an element selector function and a key equaler.
// If 'equaler' is nil go2linq.collate.DeepEqualer is used.
// If 'source' is nil or empty, zero value of go2linq.Lookup is returned.
func ToLookupSel[Source, Key, Element any](source []Source,
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler collate.Equaler[Key]) (*go2linq.Lookup[Key, Element], error) {
	if len(source) == 0 {
		return &go2linq.Lookup[Key, Element]{}, nil
	}
	r, err := go2linq.ToLookupSelEq(go2linq.NewEnSlice(source...), keySelector, elementSelector, equaler)
	if err != nil {
		return &go2linq.Lookup[Key, Element]{}, err
	}
	return r, nil
}
