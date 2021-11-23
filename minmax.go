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
// Min panics if 'source' or 'selector' or 'lesser' is nil or 'source' is empty.
//
// To get the minimum element of the sequence itself pass Identity as 'selector'.
func Min[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) Result {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	if lesser == nil {
		panic(ErrNilLesser)
	}
	min, count := minMaxResPrim(source, selector, lesser, true)
	if count == 0 {
		panic(ErrEmptySource)
	}
	return min
}

// MinErr is like Min but returns an error instead of panicking.
func MinErr[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) (res Result, err error) {
	defer func() {
		catchErrPanic[Result](recover(), &res, &err)
	}()
	return Min(source, selector, lesser), nil
}

// MinEl invokes a transform function on each element of a sequence
// and returns the element which produces the minimum resulting value.
// MinEl panics if 'source' or 'selector' or 'lesser' is nil or 'source' is empty.
func MinEl[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	if lesser == nil {
		panic(ErrNilLesser)
	}
	min, count := minMaxElPrim(source, selector, lesser, true)
	if count == 0 {
		panic(ErrEmptySource)
	}
	return min
}

// MinElErr is like MinEl but returns an error instead of panicking.
func MinElErr[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return MinEl(source, selector, lesser), nil
}

// Max invokes a transform function on each element of a sequence and returns the maximum resulting value.
// Max panics if 'source' or 'selector' or 'lesser' is nil or 'source' is empty.
//
// To get the maximum element of the sequence itself pass Identity as 'selector'.
func Max[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) Result {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	if lesser == nil {
		panic(ErrNilLesser)
	}
	max, count := minMaxResPrim(source, selector, lesser, false)
	if count == 0 {
		panic(ErrEmptySource)
	}
	return max
}

// MaxErr is like Max but returns an error instead of panicking.
func MaxErr[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) (res Result, err error) {
	defer func() {
		catchErrPanic[Result](recover(), &res, &err)
	}()
	return Max(source, selector, lesser), nil
}

// MaxEl invokes a transform function on each element of a sequence
// and returns the element which produces the maximum resulting value.
// MaxEl panics if 'source' or 'selector' or 'lesser' is nil or 'source' is empty.
func MaxEl[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	if lesser == nil {
		panic(ErrNilLesser)
	}
	max, count := minMaxElPrim(source, selector, lesser, false)
	if count == 0 {
		panic(ErrEmptySource)
	}
	return max
}

// MaxElErr is like MaxEl but returns an error instead of panicking.
func MaxElErr[Source, Result any](source Enumerator[Source],
	selector func(Source) Result, lesser Lesser[Result]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return MaxEl(source, selector, lesser), nil
}
