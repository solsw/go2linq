package slice

import (
	"github.com/solsw/go2linq/v2"
)

// GroupBy groups the elements of a slice according to a specified key selector function
// and compares the keys using 'equaler'.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func GroupBy[Source, Key any](source []Source,
	keySelector func(Source) Key, equaler go2linq.Equaler[Key]) ([]go2linq.Grouping[Key, Source], error) {
	return GroupBySel(source, keySelector, go2linq.Identity[Source], equaler)
}

// GroupByMust is like GroupBy but panics in case of error.
func GroupByMust[Source, Key any](source []Source,
	keySelector func(Source) Key, equaler go2linq.Equaler[Key]) []go2linq.Grouping[Key, Source] {
	r, err := GroupBy(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySel groups the elements of a slice according to a key selector function.
// The keys are compared using 'equaler' and each group's elements are projected using a specified function.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func GroupBySel[Source, Key, Element any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, equaler go2linq.Equaler[Key]) ([]go2linq.Grouping[Key, Element], error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []go2linq.Grouping[Key, Element]{}, nil
	}
	if equaler == nil {
		equaler = go2linq.DeepEqualer[Key]{}
	}
	lk := ToLookupSelMust(source, keySelector, elementSelector, equaler)
	gg := make([]go2linq.Grouping[Key, Element], 0, lk.Count())
	enr := lk.GetEnumerator()
	for enr.MoveNext() {
		gg = append(gg, enr.Current())
	}
	return gg, nil
}

// GroupBySelMust is like GroupBySel but panics in case of error.
func GroupBySelMust[Source, Key, Element any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, equaler go2linq.Equaler[Key]) []go2linq.Grouping[Key, Element] {
	r, err := GroupBySel(source, keySelector, elementSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByRes groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using 'equaler'.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func GroupByRes[Source, Key, Result any](source []Source, keySelector func(Source) Key,
	resultSelector func(Key, []Source) Result, equaler go2linq.Equaler[Key]) ([]Result, error) {
	return GroupBySelRes(source, keySelector, go2linq.Identity[Source], resultSelector, equaler)
}

// GroupByResMust is like GroupByRes but panics in case of error.
func GroupByResMust[Source, Key, Result any](source []Source, keySelector func(Source) Key,
	resultSelector func(Key, []Source) Result, equaler go2linq.Equaler[Key]) []Result {
	r, err := GroupByRes(source, keySelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelRes groups the elements of a slice according to a specified key selector function
// and creates a result value from each group and its key.
// Key values are compared using 'equaler' and the elements of each group are projected using 'resultSelector'.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func GroupBySelRes[Source, Key, Element, Result any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, []Element) Result, equaler go2linq.Equaler[Key]) ([]Result, error) {
	gg := GroupBySelMust(source, keySelector, elementSelector, equaler)
	return Select(gg, func(g go2linq.Grouping[Key, Element]) Result {
		return resultSelector(g.Key(), g.Values())
	})
}

// GroupBySelResMust is like GroupBySelRes but panics in case of error.
func GroupBySelResMust[Source, Key, Element, Result any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, []Element) Result, equaler go2linq.Equaler[Key]) []Result {
	r, err := GroupBySelRes(source, keySelector, elementSelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
