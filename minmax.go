package go2linq

import (
	"cmp"

	"github.com/solsw/collate"
	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 29 – Min/Max
// https://codeblog.jonskeet.uk/2011/01/09/reimplementing-linq-to-objects-part-29-min-max/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.minby
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.maxby

// 'selector' projects each element of 'source'
// 'lesser' compares the projected values
// if 'min', function searches for minimum, otherwise - for maximum
// element of sequence, corresponding projected value and count of 'source' elements are returned
func minMaxPrim[Source, Result any](source Enumerable[Source],
	selector func(Source) Result, lesser collate.Lesser[Result], min bool) (Source, Result, int) {
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

// [Min] returns the minimum value in a sequence.
//
// [Min]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func Min[Source cmp.Ordered](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	return MinSel(source, Identity[Source])
}

// MinMust is like [Min] but panics in case of error.
func MinMust[Source cmp.Ordered](source Enumerable[Source]) Source {
	return errorhelper.Must(Min(source))
}

// [MinLs] returns the minimum value in a sequence using a specified lesser.
//
// [MinLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func MinLs[Source any](source Enumerable[Source], lesser collate.Lesser[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if lesser == nil {
		return generichelper.ZeroValue[Source](), ErrNilLesser
	}
	return MinSelLs(source, Identity[Source], lesser)
}

// MinLsMust is like [MinLs] but panics in case of error.
func MinLsMust[Source any](source Enumerable[Source], lesser collate.Lesser[Source]) Source {
	return errorhelper.Must(MinLs(source, lesser))
}

// [MinSel] invokes a transform function on each element of a sequence and returns the minimum resulting value.
//
// [MinSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func MinSel[Source any, Result cmp.Ordered](source Enumerable[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	return MinSelLs(source, selector, collate.Order[Result]{})
}

// MinSelMust is like [MinSel] but panics in case of error.
func MinSelMust[Source any, Result cmp.Ordered](source Enumerable[Source], selector func(Source) Result) Result {
	return errorhelper.Must(MinSel(source, selector))
}

// [MinSelLs] invokes a transform function on each element of a sequence
// and returns the minimum resulting value using a specified lesser.
//
// [MinSelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func MinSelLs[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser collate.Lesser[Result]) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	if lesser == nil {
		return generichelper.ZeroValue[Result](), ErrNilLesser
	}
	_, min, count := minMaxPrim(source, selector, lesser, true)
	if count == 0 {
		return generichelper.ZeroValue[Result](), ErrEmptySource
	}
	return min, nil
}

// MinSelLsMust is like [MinSelLs] but panics in case of error.
func MinSelLsMust[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser collate.Lesser[Result]) Result {
	return errorhelper.Must(MinSelLs(source, selector, lesser))
}

// [MinBySel] returns the value in a sequence that produces the minimum key according to a key selector function.
//
// [MinBySel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.minby
func MinBySel[Source any, Key cmp.Ordered](source Enumerable[Source], selector func(Source) Key) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	return MinBySelLs(source, selector, collate.Order[Key]{})
}

// MinBySelMust is like [MinBySel] but panics in case of error.
func MinBySelMust[Source any, Key cmp.Ordered](source Enumerable[Source], selector func(Source) Key) Source {
	return errorhelper.Must(MinBySel(source, selector))
}

// [MinBySelLs] returns the value in a sequence that produces the minimum key according to a key selector function and a key lesser.
//
// [MinBySelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.minby
func MinBySelLs[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser collate.Lesser[Key]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	if lesser == nil {
		return generichelper.ZeroValue[Source](), ErrNilLesser
	}
	min, _, count := minMaxPrim(source, selector, lesser, true)
	if count == 0 {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return min, nil
}

// MinBySelLsMust is like [MinBySelLs] but panics in case of error.
func MinBySelLsMust[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser collate.Lesser[Key]) Source {
	return errorhelper.Must(MinBySelLs(source, selector, lesser))
}

// [Max] returns the maximum value in a sequence.
//
// [Max]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func Max[Source cmp.Ordered](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	return MaxSel(source, Identity[Source])
}

// MaxMust is like [Max] but panics in case of error.
func MaxMust[Source cmp.Ordered](source Enumerable[Source]) Source {
	return errorhelper.Must(Max(source))
}

// [MaxLs] returns the maximum value in a sequence using a specified lesser.
//
// [MaxLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func MaxLs[Source any](source Enumerable[Source], lesser collate.Lesser[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if lesser == nil {
		return generichelper.ZeroValue[Source](), ErrNilLesser
	}
	return MaxSelLs(source, Identity[Source], lesser)
}

// MaxLsMust is like [MaxLs] but panics in case of error.
func MaxLsMust[Source any](source Enumerable[Source], lesser collate.Lesser[Source]) Source {
	return errorhelper.Must(MaxLs(source, lesser))
}

// [MaxSel] invokes a transform function on each element of a sequence and returns the maximum resulting value.
//
// [MaxSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func MaxSel[Source any, Result cmp.Ordered](source Enumerable[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	return MaxSelLs(source, selector, collate.Order[Result]{})
}

// MaxSelMust is like [MaxSel] but panics in case of error.
func MaxSelMust[Source any, Result cmp.Ordered](source Enumerable[Source], selector func(Source) Result) Result {
	return errorhelper.Must(MaxSel(source, selector))
}

// [MaxSelLs] invokes a transform function on each element of a sequence
// and returns the maximum resulting value using a specified lesser.
//
// [MaxSelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func MaxSelLs[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser collate.Lesser[Result]) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	if lesser == nil {
		return generichelper.ZeroValue[Result](), ErrNilLesser
	}
	_, max, count := minMaxPrim(source, selector, lesser, false)
	if count == 0 {
		return generichelper.ZeroValue[Result](), ErrEmptySource
	}
	return max, nil
}

// MaxSelLsMust is like [MaxSelLs] but panics in case of error.
func MaxSelLsMust[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser collate.Lesser[Result]) Result {
	return errorhelper.Must(MaxSelLs(source, selector, lesser))
}

// [MaxBySel] returns the value in a sequence that produces the maximum key according to a key selector function.
//
// [MaxBySel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.maxby
func MaxBySel[Source any, Key cmp.Ordered](source Enumerable[Source], selector func(Source) Key) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	return MaxBySelLs(source, selector, collate.Order[Key]{})
}

// MaxBySelMust is like [MaxBySel] but panics in case of error.
func MaxBySelMust[Source any, Key cmp.Ordered](source Enumerable[Source], selector func(Source) Key) Source {
	return errorhelper.Must(MaxBySel(source, selector))
}

// [MaxBySelLs] returns the value in a sequence that produces the maximum key according to a key selector function and a key lesser.
//
// [MaxBySelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.maxby
func MaxBySelLs[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser collate.Lesser[Key]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	if lesser == nil {
		return generichelper.ZeroValue[Source](), ErrNilLesser
	}
	max, _, count := minMaxPrim(source, selector, lesser, false)
	if count == 0 {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return max, nil
}

// MaxBySelLsMust is like [MaxBySelLs] but panics in case of error.
func MaxBySelLsMust[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser collate.Lesser[Key]) Source {
	return errorhelper.Must(MaxBySelLs(source, selector, lesser))
}
