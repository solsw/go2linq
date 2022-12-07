package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Where filters a slice of Source based on a predicate.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func Where[Source any](source []Source, predicate func(Source) bool) ([]Source, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Source{}, nil
	}
	en, err := go2linq.Where(go2linq.NewEnSlice(source...), predicate)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// WhereMust is like Where but panics in case of error.
func WhereMust[Source any](source []Source, predicate func(Source) bool) []Source {
	r, err := Where(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// WhereIdx filters a slice of Source based on a predicate.
// Each element's index is used in the logic of the predicate function.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func WhereIdx[Source any](source []Source, predicate func(Source, int) bool) ([]Source, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Source{}, nil
	}
	en, err := go2linq.WhereIdx(go2linq.NewEnSlice(source...), predicate)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// WhereIdxMust is like WhereIdx but panics in case of error.
func WhereIdxMust[Source any](source []Source, predicate func(Source, int) bool) []Source {
	r, err := WhereIdx(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
