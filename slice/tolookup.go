package slice

import (
	"github.com/solsw/go2linq/v2"
)

// ToLookup creates a go2linq.Lookup from a slice according to a specified key selector function.
// If 'source' is nil or empty, zero value of go2linq.Lookup is returned.
// go2linq.DeepEqualer is used to compare keys.
func ToLookup[Source, Key any](source []Source, keySelector func(Source) Key) (*go2linq.Lookup[Key, Source], error) {
	return ToLookupSelEq(source, keySelector, go2linq.Identity[Source], nil)
}

// ToLookupMust is like ToLookup but panics in case of error.
func ToLookupMust[Source, Key any](source []Source, keySelector func(Source) Key) *go2linq.Lookup[Key, Source] {
	r, err := ToLookup(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupEq creates a go2linq.Lookup from a slice according to a specified key selector function and a key equaler.
// If 'source' is nil or empty, zero value of go2linq.Lookup is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func ToLookupEq[Source, Key any](source []Source, keySelector func(Source) Key, equaler go2linq.Equaler[Key]) (*go2linq.Lookup[Key, Source], error) {
	return ToLookupSelEq(source, keySelector, go2linq.Identity[Source], equaler)
}

// ToLookupEqMust is like ToLookupEq but panics in case of error.
func ToLookupEqMust[Source, Key any](source []Source, keySelector func(Source) Key, equaler go2linq.Equaler[Key]) *go2linq.Lookup[Key, Source] {
	r, err := ToLookupEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupSel creates a go2linq.Lookup from a slice
// according to specified key selector and element selector functions.
// If 'source' is nil or empty, zero value of go2linq.Lookup is returned.
// go2linq.DeepEqualer is used to compare keys.
func ToLookupSel[Source, Key, Element any](source []Source,
	keySelector func(Source) Key, elementSelector func(Source) Element) (*go2linq.Lookup[Key, Element], error) {
	return ToLookupSelEq(source, keySelector, elementSelector, nil)
}

// ToLookupSelMust is like ToLookupSel but panics in case of error.
func ToLookupSelMust[Source, Key, Element any](source []Source,
	keySelector func(Source) Key, elementSelector func(Source) Element) *go2linq.Lookup[Key, Element] {
	r, err := ToLookupSel(source, keySelector, elementSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ToLookupSelEq creates a go2linq.Lookup from a slice according to a specified key selector function,
// an element selector function and a key equaler.
// If 'source' is nil or empty, zero value of go2linq.Lookup is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func ToLookupSelEq[Source, Key, Element any](source []Source,
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler go2linq.Equaler[Key]) (*go2linq.Lookup[Key, Element], error) {
	if len(source) == 0 {
		return &go2linq.Lookup[Key, Element]{}, nil
	}
	r, err := go2linq.ToLookupSelEq(go2linq.NewEnSlice(source...), keySelector, elementSelector, equaler)
	if err != nil {
		return &go2linq.Lookup[Key, Element]{}, err
	}
	return r, nil
}

// ToLookupSelEqMust is like ToLookupSelEq but panics in case of error.
func ToLookupSelEqMust[Source, Key, Element any](source []Source,
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler go2linq.Equaler[Key]) *go2linq.Lookup[Key, Element] {
	r, err := ToLookupSelEq(source, keySelector, elementSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
