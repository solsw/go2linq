//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 29 – Min/Max
// https://codeblog.jonskeet.uk/2011/01/09/reimplementing-linq-to-objects-part-29-min-max/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max

// 'selector' projects each element of 'source'
// 'lesser' compares the projected values
// if 'min', function searches for minimum, otherwise - for maximum
// corresponding projected value (with count of 'source' elements) is returned
func minMaxResPrim[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result], min bool) (Result, int) {
	count := 0
	first := true
	var rs Result
	for source.MoveNext() {
		count++
		if first {
			first = false
			rs = selector(source.Current())
			continue
		}
		s := selector(source.Current())
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
func minMaxElPrim[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result], min bool) (Source, int) {
	count := 0
	first := true
	var re Source
	var rs Result
	for source.MoveNext() {
		count++
		if first {
			first = false
			re = source.Current()
			rs = selector(re)
			continue
		}
		e := source.Current()
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
func Min[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) (Result, error) {
	if source == nil {
		var r0 Result
		return r0, ErrNilSource
	}
	if selector == nil {
		var r0 Result
		return r0, ErrNilSelector
	}
	if lesser == nil {
		var r0 Result
		return r0, ErrNilLesser
	}
	min, count := minMaxResPrim(source, selector, lesser, true)
	if count == 0 {
		var r0 Result
		return r0, ErrEmptySource
	}
	return min, nil
}

// MinMust is like Min but panics in case of error.
func MinMust[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) Result {
	r, err := Min(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MinEl invokes a transform function on each element of a sequence
// and returns the element which produces the minimum resulting value.
func MinEl[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if selector == nil {
		var s0 Source
		return s0, ErrNilSelector
	}
	if lesser == nil {
		var s0 Source
		return s0, ErrNilLesser
	}
	min, count := minMaxElPrim(source, selector, lesser, true)
	if count == 0 {
		var s0 Source
		return s0, ErrEmptySource
	}
	return min, nil
}

// MinElMust is like MinEl but panics in case of error.
func MinElMust[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) Source {
	r, err := MinEl(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// Max invokes a transform function on each element of a sequence and returns the maximum resulting value.
//
// To get the maximum element of the sequence itself pass Identity as 'selector'.
func Max[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) (Result, error) {
	if source == nil {
		var r0 Result
		return r0, ErrNilSource
	}
	if selector == nil {
		var r0 Result
		return r0, ErrNilSelector
	}
	if lesser == nil {
		var r0 Result
		return r0, ErrNilLesser
	}
	max, count := minMaxResPrim(source, selector, lesser, false)
	if count == 0 {
		var r0 Result
		return r0, ErrEmptySource
	}
	return max, nil
}

// MaxMust is like Max but panics in case of error.
func MaxMust[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) Result {
	r, err := Max(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// MaxEl invokes a transform function on each element of a sequence
// and returns the element which produces the maximum resulting value.
func MaxEl[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if selector == nil {
		var s0 Source
		return s0, ErrNilSelector
	}
	if lesser == nil {
		var s0 Source
		return s0, ErrNilLesser
	}
	max, count := minMaxElPrim(source, selector, lesser, false)
	if count == 0 {
		var s0 Source
		return s0, ErrEmptySource
	}
	return max, nil
}

// MaxElMust is like MaxEl but panics in case of error.
func MaxElMust[Source, Result any](source Enumerator[Source], selector func(Source) Result, lesser Lesser[Result]) Source {
	r, err := MaxEl(source, selector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}
