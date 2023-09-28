package go2linq

import (
	"github.com/solsw/errorhelper"
)

// Reimplementing LINQ to Objects: Part 9 - SelectMany
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-9-selectmany/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

func factorySelectMany[Source, Result any](source Enumerable[Source], selector func(Source) Enumerable[Result]) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enrSource := source.GetEnumerator()
		enrTmp := Empty[Result]().GetEnumerator()
		return enrFunc[Result]{
			mvNxt: func() bool {
				for {
					if enrTmp.MoveNext() {
						return true
					}
					if !enrSource.MoveNext() {
						return false
					}
					enrTmp = selector(enrSource.Current()).GetEnumerator()
				}
			},
			// crrnt: enrTmp.Current, // yields wrong results
			crrnt: func() Result { return enrTmp.Current() },
			rst: func() {
				enrTmp = Empty[Result]().GetEnumerator()
				enrSource.Reset()
			},
		}
	}
}

// [SelectMany] projects each element of a sequence to an [Enumerable]
// and flattens the resulting sequences into one sequence.
//
// [SelectMany]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectMany[Source, Result any](source Enumerable[Source], selector func(Source) Enumerable[Result]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return OnFactory(factorySelectMany(source, selector)), nil
}

// SelectManyMust is like [SelectMany] but panics in case of error.
func SelectManyMust[Source, Result any](source Enumerable[Source], selector func(Source) Enumerable[Result]) Enumerable[Result] {
	return errorhelper.Must(SelectMany(source, selector))
}

func factorySelectManyIdx[Source, Result any](source Enumerable[Source], selector func(Source, int) Enumerable[Result]) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enrSource := source.GetEnumerator()
		enrTmp := Empty[Result]().GetEnumerator()
		i := -1
		return enrFunc[Result]{
			mvNxt: func() bool {
				for {
					if enrTmp.MoveNext() {
						return true
					}
					if !enrSource.MoveNext() {
						return false
					}
					i++
					enrTmp = selector(enrSource.Current(), i).GetEnumerator()
				}
			},
			crrnt: func() Result { return enrTmp.Current() },
			rst: func() {
				i = -1
				enrTmp = Empty[Result]().GetEnumerator()
				enrSource.Reset()
			},
		}
	}
}

// [SelectManyIdx] projects each element of a sequence and its index to an [Enumerable]
// and flattens the resulting sequences into one sequence.
//
// [SelectManyIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectManyIdx[Source, Result any](source Enumerable[Source], selector func(Source, int) Enumerable[Result]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return OnFactory(factorySelectManyIdx(source, selector)), nil
}

// SelectManyIdxMust is like [SelectManyIdx] but panics in case of error.
func SelectManyIdxMust[Source, Result any](source Enumerable[Source], selector func(Source, int) Enumerable[Result]) Enumerable[Result] {
	return errorhelper.Must(SelectManyIdx(source, selector))
}

func factorySelectManyColl[Source, Collection, Result any](source Enumerable[Source],
	collectionSelector func(Source) Enumerable[Collection], resultSelector func(Source, Collection) Result) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enrSource := source.GetEnumerator()
		enrTmp := Empty[Collection]().GetEnumerator()
		var e1 Source
		return enrFunc[Result]{
			mvNxt: func() bool {
				for {
					if enrTmp.MoveNext() {
						return true
					}
					if !enrSource.MoveNext() {
						return false
					}
					e1 = enrSource.Current()
					enrTmp = collectionSelector(e1).GetEnumerator()
				}
			},
			crrnt: func() Result { return resultSelector(e1, enrTmp.Current()) },
			rst: func() {
				enrTmp = Empty[Collection]().GetEnumerator()
				enrSource.Reset()
			},
		}
	}
}

// [SelectManyColl] projects each element of a sequence to an [Enumerable],
// flattens the resulting sequences into one sequence
// and invokes a result selector function on each element therein.
//
// [SelectManyColl]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectManyColl[Source, Collection, Result any](source Enumerable[Source],
	collectionSelector func(Source) Enumerable[Collection], resultSelector func(Source, Collection) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if collectionSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return OnFactory(factorySelectManyColl(source, collectionSelector, resultSelector)), nil
}

// SelectManyCollMust is like [SelectManyColl] but panics in case of error.
func SelectManyCollMust[Source, Collection, Result any](source Enumerable[Source],
	collectionSelector func(Source) Enumerable[Collection], resultSelector func(Source, Collection) Result) Enumerable[Result] {
	return errorhelper.Must(SelectManyColl(source, collectionSelector, resultSelector))
}

func factorySelectManyCollIdx[Source, Collection, Result any](source Enumerable[Source],
	collectionSelector func(Source, int) Enumerable[Collection], resultSelector func(Source, Collection) Result) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enrSource := source.GetEnumerator()
		enrTmp := Empty[Collection]().GetEnumerator()
		var e1 Source
		i := -1
		return enrFunc[Result]{
			mvNxt: func() bool {
				for {
					if enrTmp.MoveNext() {
						return true
					}
					if !enrSource.MoveNext() {
						return false
					}
					e1 = enrSource.Current()
					i++
					enrTmp = collectionSelector(e1, i).GetEnumerator()
				}
			},
			crrnt: func() Result { return resultSelector(e1, enrTmp.Current()) },
			rst: func() {
				i = -1
				enrTmp = Empty[Collection]().GetEnumerator()
				enrSource.Reset()
			},
		}
	}
}

// [SelectManyCollIdx] projects each element of a sequence and its index to an [Enumerable],
// flattens the resulting sequences into one sequence
// and invokes a result selector function on each element therein.
//
// [SelectManyCollIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectManyCollIdx[Source, Collection, Result any](source Enumerable[Source],
	collectionSelector func(Source, int) Enumerable[Collection], resultSelector func(Source, Collection) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if collectionSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return OnFactory(factorySelectManyCollIdx(source, collectionSelector, resultSelector)), nil
}

// SelectManyCollIdxMust is like [SelectManyCollIdx] but panics in case of error.
func SelectManyCollIdxMust[Source, Collection, Result any](source Enumerable[Source],
	collectionSelector func(Source, int) Enumerable[Collection], resultSelector func(Source, Collection) Result) Enumerable[Result] {
	return errorhelper.Must(SelectManyCollIdx(source, collectionSelector, resultSelector))
}
