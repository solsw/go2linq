package slice

import (
	"github.com/solsw/go2linq/v2"
)

// SelectMany projects each element of a slice to a slice and flattens the resulting slices into one slice.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func SelectMany[Source, Result any](source []Source, selector func(Source) []Result) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.SelectMany(go2linq.NewEnSlice(source...),
		func(s Source) go2linq.Enumerable[Result] { return go2linq.NewEnSlice(selector(s)...) },
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// SelectManyMust is like SelectMany but panics in case of error.
func SelectManyMust[Source, Result any](source []Source, selector func(Source) []Result) []Result {
	r, err := SelectMany(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyIdx projects each element of a slice and its index to a slice
// and flattens the resulting slices into one slice.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func SelectManyIdx[Source, Result any](source []Source, selector func(Source, int) []Result) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.SelectManyIdx(go2linq.NewEnSlice(source...),
		func(s Source, idx int) go2linq.Enumerable[Result] { return go2linq.NewEnSlice(selector(s, idx)...) },
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// SelectManyIdxMust is like SelectManyIdx but panics in case of error.
func SelectManyIdxMust[Source, Result any](source []Source, selector func(Source, int) []Result) []Result {
	r, err := SelectManyIdx(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyColl projects each element of a slice to a slice,
// flattens the resulting slices into one slice,
// and invokes a result selector function on each element therein.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func SelectManyColl[Source, Collection, Result any](source []Source,
	collectionSelector func(Source) []Collection, resultSelector func(Source, Collection) Result) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.SelectManyColl(go2linq.NewEnSlice(source...),
		func(s Source) go2linq.Enumerable[Collection] { return go2linq.NewEnSlice(collectionSelector(s)...) },
		resultSelector,
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// SelectManyCollMust is like SelectManyColl but panics in case of error.
func SelectManyCollMust[Source, Collection, Result any](source []Source,
	collectionSelector func(Source) []Collection, resultSelector func(Source, Collection) Result) []Result {
	r, err := SelectManyColl(source, collectionSelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyCollIdx projects each element of a slice and its index to a slice,
// flattens the resulting slices into one slice,
// and invokes a result selector function on each element therein.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func SelectManyCollIdx[Source, Collection, Result any](source []Source,
	collectionSelector func(Source, int) []Collection, resultSelector func(Source, Collection) Result) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.SelectManyCollIdx(go2linq.NewEnSlice(source...),
		func(s Source, idx int) go2linq.Enumerable[Collection] {
			return go2linq.NewEnSlice(collectionSelector(s, idx)...)
		},
		resultSelector,
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// SelectManyCollIdxMust is like SelectManyCollIdx but panics in case of error.
func SelectManyCollIdxMust[Source, Collection, Result any](source []Source,
	collectionSelector func(Source, int) []Collection, resultSelector func(Source, Collection) Result) []Result {
	r, err := SelectManyCollIdx(source, collectionSelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
