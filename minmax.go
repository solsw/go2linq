package go2linq

import (
	"cmp"
	"iter"

	"github.com/solsw/generichelper"
)

// 'selector' projects each element of 'source'
// 'less' compares the projected values
// if 'min', function searches for minimum, otherwise - for maximum
// returns: element of sequence, corresponding projected value, count of 'source' elements
func minMaxPrim[Source, Result any](source iter.Seq[Source], selector func(Source) Result,
	less func(Result, Result) bool, min bool) (Source, Result, int) {
	count := 0
	first := true
	var rs Source
	var rr Result
	for s := range source {
		count++
		if first {
			first = false
			rs = s
			rr = selector(s)
			continue
		}
		r := selector(s)
		if (min && less(r, rr)) || (!min && less(rr, r)) {
			rs = s
			rr = r
		}
	}
	return rs, rr, count
}

// [Min] returns the minimum value in a sequence.
//
// [Min]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func Min[Source cmp.Ordered](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	return MinSel(source, Identity[Source])
}

// [MinLs] returns the minimum value in a sequence using a specified less.
//
// [MinLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func MinLs[Source any](source iter.Seq[Source], less func(Source, Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if less == nil {
		return generichelper.ZeroValue[Source](), ErrNilLess
	}
	return MinSelLs(source, Identity[Source], less)
}

// [MinSel] invokes a transform function on each element of a sequence and returns the minimum resulting value.
//
// [MinSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func MinSel[Source any, Result cmp.Ordered](source iter.Seq[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	return MinSelLs(source, selector, cmp.Less[Result])
}

// [MinSelLs] invokes a transform function on each element of a sequence
// and returns the minimum resulting value using a specified less.
//
// [MinSelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func MinSelLs[Source, Result any](source iter.Seq[Source], selector func(Source) Result, less func(Result, Result) bool) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	if less == nil {
		return generichelper.ZeroValue[Result](), ErrNilLess
	}
	_, min, count := minMaxPrim(source, selector, less, true)
	if count == 0 {
		return generichelper.ZeroValue[Result](), ErrEmptySource
	}
	return min, nil
}

// [MinBySel] returns the value in a sequence that produces the minimum key according to a key selector function.
//
// [MinBySel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.minby
func MinBySel[Source any, Key cmp.Ordered](source iter.Seq[Source], selector func(Source) Key) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	return MinBySelLs(source, selector, cmp.Less[Key])
}

// [MinBySelLs] returns the value in a sequence that produces the minimum key according to a key selector function and a key less.
//
// [MinBySelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.minby
func MinBySelLs[Source, Key any](source iter.Seq[Source], selector func(Source) Key, less func(Key, Key) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	if less == nil {
		return generichelper.ZeroValue[Source](), ErrNilLess
	}
	min, _, count := minMaxPrim(source, selector, less, true)
	if count == 0 {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return min, nil
}

// [Max] returns the maximum value in a sequence.
//
// [Max]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func Max[Source cmp.Ordered](source iter.Seq[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	return MaxSel(source, Identity[Source])
}

// [MaxLs] returns the maximum value in a sequence using a specified less.
//
// [MaxLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func MaxLs[Source any](source iter.Seq[Source], less func(Source, Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if less == nil {
		return generichelper.ZeroValue[Source](), ErrNilLess
	}
	return MaxSelLs(source, Identity[Source], less)
}

// [MaxSel] invokes a transform function on each element of a sequence and returns the maximum resulting value.
//
// [MaxSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func MaxSel[Source any, Result cmp.Ordered](source iter.Seq[Source], selector func(Source) Result) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	return MaxSelLs(source, selector, cmp.Less[Result])
}

// [MaxSelLs] invokes a transform function on each element of a sequence
// and returns the maximum resulting value using a specified less.
//
// [MaxSelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func MaxSelLs[Source, Result any](source iter.Seq[Source], selector func(Source) Result, less func(Result, Result) bool) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	if less == nil {
		return generichelper.ZeroValue[Result](), ErrNilLess
	}
	_, max, count := minMaxPrim(source, selector, less, false)
	if count == 0 {
		return generichelper.ZeroValue[Result](), ErrEmptySource
	}
	return max, nil
}

// [MaxBySel] returns the value in a sequence that produces the maximum key according to a key selector function.
//
// [MaxBySel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.maxby
func MaxBySel[Source any, Key cmp.Ordered](source iter.Seq[Source], selector func(Source) Key) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	return MaxBySelLs(source, selector, cmp.Less[Key])
}

// [MaxBySelLs] returns the value in a sequence that produces the maximum key according to a key selector function and a key less.
//
// [MaxBySelLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.maxby
func MaxBySelLs[Source, Key any](source iter.Seq[Source], selector func(Source) Key, less func(Key, Key) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return generichelper.ZeroValue[Source](), ErrNilSelector
	}
	if less == nil {
		return generichelper.ZeroValue[Source](), ErrNilLess
	}
	max, _, count := minMaxPrim(source, selector, less, false)
	if count == 0 {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return max, nil
}
