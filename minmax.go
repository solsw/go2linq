//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 29 – Min/Max
// https://codeblog.jonskeet.uk/2011/01/09/reimplementing-linq-to-objects-part-29-min-max/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.minby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.maxby

// 'selector' projects each element of 'source'
// 'lesser' compares the projected values
// if 'min', function searches for minimum, otherwise - for maximum
// corresponding projected value (with count of 'source' elements) is returned
func minMaxResPrim[Source, Result any](source Enumerable[Source],
	selector func(Source) Result, lesser Lesser[Result], min bool) (Result, int) {
	enr := source.GetEnumerator()
	count := 0
	first := true
	var rs Result
	for enr.MoveNext() {
		count++
		if first {
			first = false
			rs = selector(enr.Current())
			continue
		}
		s := selector(enr.Current())
		if (min && lesser.Less(s, rs)) || (!min && lesser.Less(rs, s)) {
			rs = s
		}
	}
	return rs, count
}

// 'selector' projects each element of 'source'
// 'lesser' compares the projected values
// if 'min', function searches for minimum, otherwise - for maximum
// element of sequence which produces corresponding projected value (with count of 'source' elements) is returned
func minMaxByPrim[Source, Result any](source Enumerable[Source],
	selector func(Source) Result, lesser Lesser[Result], min bool) (Source, int) {
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
	return re, count
}

// Min invokes a transform function on each element of a sequence and returns the minimum resulting value.
//
// To get the minimum element of the sequence itself pass Identity as 'selector'.
func Min[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) (Result, error) {
	if source == nil {
		return ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Result](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Result](), ErrNilLesser
	}
	min, count := minMaxResPrim(source, selector, lesser, true)
	if count == 0 {
		return ZeroValue[Result](), ErrEmptySource
	}
	return min, nil
}

// MinMust is like Min but panics in case of error.
func MinMust[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) Result {
	r, err := Min(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MinBy returns the minimum value in a generic sequence according to a specified key selector function and key lesser.
func MinBy[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Source](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Source](), ErrNilLesser
	}
	min, count := minMaxByPrim(source, selector, lesser, true)
	if count == 0 {
		return ZeroValue[Source](), ErrEmptySource
	}
	return min, nil
}

// MinByMust is like MinBy but panics in case of error.
func MinByMust[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) Source {
	r, err := MinBy(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// Max invokes a transform function on each element of a sequence and returns the maximum resulting value.
//
// To get the maximum element of the sequence itself pass Identity as 'selector'.
func Max[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) (Result, error) {
	if source == nil {
		return ZeroValue[Result](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Result](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Result](), ErrNilLesser
	}
	max, count := minMaxResPrim(source, selector, lesser, false)
	if count == 0 {
		return ZeroValue[Result](), ErrEmptySource
	}
	return max, nil
}

// MaxMust is like Max but panics in case of error.
func MaxMust[Source, Result any](source Enumerable[Source], selector func(Source) Result, lesser Lesser[Result]) Result {
	r, err := Max(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxBy returns the maximum value in a generic sequence according to a specified key selector function and key lesser.
func MaxBy[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if selector == nil {
		return ZeroValue[Source](), ErrNilSelector
	}
	if lesser == nil {
		return ZeroValue[Source](), ErrNilLesser
	}
	max, count := minMaxByPrim(source, selector, lesser, false)
	if count == 0 {
		return ZeroValue[Source](), ErrEmptySource
	}
	return max, nil
}

// MaxByMust is like MaxBy but panics in case of error.
func MaxByMust[Source, Key any](source Enumerable[Source], selector func(Source) Key, lesser Lesser[Key]) Source {
	r, err := MaxBy(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}
