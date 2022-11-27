package slice

import (
	"github.com/solsw/go2linq/v2"
)

// GroupBy groups the elements of a slice according to a specified key selector function.
// The keys are compared using go2linq.DeepEqualer.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func GroupBy[Source, Key any](source []Source, keySelector func(Source) Key) ([]go2linq.Grouping[Key, Source], error) {
	return GroupBySelEq(source, keySelector, go2linq.Identity[Source], nil)
}

// GroupByMust is like GroupBy but panics in case of error.
func GroupByMust[Source, Key any](source []Source, keySelector func(Source) Key) []go2linq.Grouping[Key, Source] {
	r, err := GroupBy(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByEq groups the elements of a slice according to a specified key selector function
// and compares the keys using 'equaler'.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func GroupByEq[Source, Key any](source []Source,
	keySelector func(Source) Key, equaler go2linq.Equaler[Key]) ([]go2linq.Grouping[Key, Source], error) {
	return GroupBySelEq(source, keySelector, go2linq.Identity[Source], equaler)
}

// GroupByEqMust is like GroupByEq but panics in case of error.
func GroupByEqMust[Source, Key any](source []Source,
	keySelector func(Source) Key, equaler go2linq.Equaler[Key]) []go2linq.Grouping[Key, Source] {
	r, err := GroupByEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySel groups the elements of a slice according to a specified key selector function
// and projects the elements for each group using a specified function.
// The keys are compared using go2linq.DeepEqualer.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func GroupBySel[Source, Key, Element any](source []Source,
	keySelector func(Source) Key, elementSelector func(Source) Element) ([]go2linq.Grouping[Key, Element], error) {
	return GroupBySelEq(source, keySelector, elementSelector, nil)
}

// GroupBySelMust is like GroupBySel but panics in case of error.
func GroupBySelMust[Source, Key, Element any](source []Source,
	keySelector func(Source) Key, elementSelector func(Source) Element) []go2linq.Grouping[Key, Element] {
	r, err := GroupBySel(source, keySelector, elementSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelEq groups the elements of a slice according to a key selector function.
// The keys are compared using 'equaler' and each group's elements are projected using a specified function.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func GroupBySelEq[Source, Key, Element any](source []Source, keySelector func(Source) Key,
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
	lk := ToLookupSelEqMust(source, keySelector, elementSelector, equaler)
	gg := make([]go2linq.Grouping[Key, Element], 0, lk.Count())
	enr := lk.GetEnumerator()
	for enr.MoveNext() {
		gg = append(gg, enr.Current())
	}
	return gg, nil
}

// GroupBySelEqMust is like GroupBySelEq but panics in case of error.
func GroupBySelEqMust[Source, Key, Element any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, equaler go2linq.Equaler[Key]) []go2linq.Grouping[Key, Element] {
	r, err := GroupBySelEq(source, keySelector, elementSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByRes groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using go2linq.DeepEqualer.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func GroupByRes[Source, Key, Result any](source []Source,
	keySelector func(Source) Key, resultSelector func(Key, []Source) Result) ([]Result, error) {
	return GroupBySelResEq(source, keySelector, go2linq.Identity[Source], resultSelector, nil)
}

// GroupByResMust is like GroupByRes but panics in case of error.
func GroupByResMust[Source, Key, Result any](source []Source,
	keySelector func(Source) Key, resultSelector func(Key, []Source) Result) []Result {
	r, err := GroupByRes(source, keySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByResEq groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using 'equaler'.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func GroupByResEq[Source, Key, Result any](source []Source, keySelector func(Source) Key,
	resultSelector func(Key, []Source) Result, equaler go2linq.Equaler[Key]) ([]Result, error) {
	return GroupBySelResEq(source, keySelector, go2linq.Identity[Source], resultSelector, equaler)
}

// GroupByResEqMust is like GroupByResEq but panics in case of error.
func GroupByResEqMust[Source, Key, Result any](source []Source, keySelector func(Source) Key,
	resultSelector func(Key, []Source) Result, equaler go2linq.Equaler[Key]) []Result {
	r, err := GroupByResEq(source, keySelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelRes groups the elements of a slice according to a specified key selector function
// and creates a result value from each group and its key.
// The elements of each group are projected using 'resultSelector'.
// Key values are compared using go2linq.DeepEqualer.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func GroupBySelRes[Source, Key, Element, Result any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, []Element) Result) ([]Result, error) {
	return GroupBySelResEq(source, keySelector, elementSelector, resultSelector, nil)
}

// GroupBySelResMust is like GroupBySelRes but panics in case of error.
func GroupBySelResMust[Source, Key, Element, Result any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, []Element) Result) []Result {
	r, err := GroupBySelRes(source, keySelector, elementSelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelResEq groups the elements of a slice according to a specified key selector function
// and creates a result value from each group and its key.
// Key values are compared using 'equaler' and the elements of each group are projected using 'resultSelector'.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func GroupBySelResEq[Source, Key, Element, Result any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, []Element) Result, equaler go2linq.Equaler[Key]) ([]Result, error) {
	gg := GroupBySelEqMust(source, keySelector, elementSelector, equaler)
	return Select(gg, func(g go2linq.Grouping[Key, Element]) Result {
		return resultSelector(g.Key(), g.Values())
	})
}

// GroupBySelResEqMust is like GroupBySelResEq but panics in case of error.
func GroupBySelResEqMust[Source, Key, Element, Result any](source []Source, keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, []Element) Result, equaler go2linq.Equaler[Key]) []Result {
	r, err := GroupBySelResEq(source, keySelector, elementSelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
