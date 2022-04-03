//go:build go1.18

package go2linq

import (
	"golang.org/x/exp/constraints"
)

// Reimplementing LINQ to Objects: Part 29 – Min/Max
// https://codeblog.jonskeet.uk/2011/01/09/reimplementing-linq-to-objects-part-29-min-max/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.minby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.maxby

// 'selector' projects each element of 'source'
// 'lesser' compares the projected values
// if 'min', function searches for minimum, otherwise - for maximum
// element of sequence, corresponding projected value and count of 'source' elements are returned
func minMaxPrim[Source, Result any](source Enumerable[Source],
	selector func(Source) Result, lesser Lesser[Result], min bool) (Source, Result, int) {
	enr := source.GetEnumerator()
	count := 0
	first := true
	var re Source
	var rs Result
	for enr.MoveNext() {
		count++
		if first {
			first = false
			re = enr.Current()
			rs = selector(re)
			continue
		}
		e := enr.Current()
		s := selector(e)
		if (min && lesser.Less(s, rs)) || (!min && lesser.Less(rs, s)) {
			re = e
			rs = s
		}
	}
	return re, rs, count
}

// Min returns the minimum value in a sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min)
func Min[Source constraints.Ordered](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	return MinSel(source, Identity[Source])
}

// MinMust is like Min but panics in case of an error.
func MinMust[Source constraints.Ordered](source Enumerable[Source]) Source {
	r, err := Min(source)
	if err != nil {
		panic(err)
	}
	return r
}

// MinLs returns the minimum value in a sequence using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min)
func MinLs[Source any](source Enumerable[Source], lesser Lesser[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if lesser == nil {
		return ZeroValue[Source](), ErrNilLesser
	}
	return MinSelLs(source, Identity[Source], lesser)
}

// MinLsMust is like MinLs but panics in case of an error.
func MinLsMust[Source any](source Enumerable[Source], lesser Lesser[Source]) Source {
	r, err := MinLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MinSel invokes a transform function on each element of a sequence and returns the minimum resulting value.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min)
func MinSel[Source any, Result constraints.Ordered](source Enumerable[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Result](), ErrNilSelector
	}
	return MinSelLs(source, selector, Lesser[Result](Order[Result]{}))
}

// MinSelMust is like MinSel but panics in case of an error.
func MinSelMust[Source any, Result constraints.Ordered](source Enumerable[Source], selector func(Source) Result) Result {
	r, err := MinSel(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// MinSelLs invokes a transform function on each element of a sequence
// and returns the minimum resulting value using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min)
func MinSelLs[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) (Result, error) {
	if source == nil {
		return ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Result](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Result](), ErrNilLesser
	}
	_, min, count := minMaxPrim(source, selector, lesser, true)
	if count == 0 {
		return ZeroValue[Result](), ErrEmptySource
	}
	return min, nil
}

// MinSelLsMust is like MinSelLs but panics in case of an error.
func MinSelLsMust[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) Result {
	r, err := MinSelLs(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MinBySel returns the value in a sequence that produces the minimum key according to a key selector function.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.minby)
func MinBySel[Source any, Key constraints.Ordered](source Enumerable[Source], selector func(Source) Key) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Source](), ErrNilSelector
	}
	return MinBySelLs(source, selector, Lesser[Key](Order[Key]{}))
}

// MinBySelMust is like MinBySel but panics in case of an error.
func MinBySelMust[Source any, Key constraints.Ordered](source Enumerable[Source], selector func(Source) Key) Source {
	r, err := MinBySel(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// MinBySelLs returns the value in a sequence that produces the minimum key according to a key selector function and a key lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.minby)
func MinBySelLs[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Source](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Source](), ErrNilLesser
	}
	min, _, count := minMaxPrim(source, selector, lesser, true)
	if count == 0 {
		return ZeroValue[Source](), ErrEmptySource
	}
	return min, nil
}

// MinBySelLsMust is like MinBySelLs but panics in case of an error.
func MinBySelLsMust[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) Source {
	r, err := MinBySelLs(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// Max returns the maximum value in a sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max)
func Max[Source constraints.Ordered](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	return MaxSel(source, Identity[Source])
}

// MaxMust is like Max but panics in case of an error.
func MaxMust[Source constraints.Ordered](source Enumerable[Source]) Source {
	r, err := Max(source)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxLs returns the maximum value in a sequence using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max)
func MaxLs[Source any](source Enumerable[Source], lesser Lesser[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if lesser == nil {
		return ZeroValue[Source](), ErrNilLesser
	}
	return MaxSelLs(source, Identity[Source], lesser)
}

// MaxLsMust is like MaxLs but panics in case of an error.
func MaxLsMust[Source any](source Enumerable[Source], lesser Lesser[Source]) Source {
	r, err := MaxLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxSel invokes a transform function on each element of a sequence and returns the maximum resulting value.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max)
func MaxSel[Source any, Result constraints.Ordered](source Enumerable[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Result](), ErrNilSelector
	}
	return MaxSelLs(source, selector, Lesser[Result](Order[Result]{}))
}

// MaxSelMust is like MaxSel but panics in case of an error.
func MaxSelMust[Source any, Result constraints.Ordered](source Enumerable[Source], selector func(Source) Result) Result {
	r, err := MaxSel(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxSelLs invokes a transform function on each element of a sequence
// and returns the maximum resulting value using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max)
func MaxSelLs[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) (Result, error) {
	if source == nil {
		return ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Result](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Result](), ErrNilLesser
	}
	_, max, count := minMaxPrim(source, selector, lesser, false)
	if count == 0 {
		return ZeroValue[Result](), ErrEmptySource
	}
	return max, nil
}

// MaxSelLsMust is like MaxSelLs but panics in case of an error.
func MaxSelLsMust[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) Result {
	r, err := MaxSelLs(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxBySel returns the value in a sequence that produces the maximum key according to a key selector function.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.maxby)
func MaxBySel[Source any, Key constraints.Ordered](source Enumerable[Source], selector func(Source) Key) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Source](), ErrNilSelector
	}
	return MaxBySelLs(source, selector, Lesser[Key](Order[Key]{}))
}

// MaxBySelMust is like MaxBySel but panics in case of an error.
func MaxBySelMust[Source any, Key constraints.Ordered](source Enumerable[Source], selector func(Source) Key) Source {
	r, err := MaxBySel(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxBySelLs returns the value in a sequence that produces the maximum key according to a key selector function and a key lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.maxby)
func MaxBySelLs[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Source](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Source](), ErrNilLesser
	}
	max, _, count := minMaxPrim(source, selector, lesser, false)
	if count == 0 {
		return ZeroValue[Source](), ErrEmptySource
	}
	return max, nil
}

// MaxBySelLsMust is like MaxBySelLs but panics in case of an error.
func MaxBySelLsMust[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) Source {
	r, err := MaxBySelLs(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}
