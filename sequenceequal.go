//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 34 – SequenceEqual
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-34-sequenceequal/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

// SequenceEqual determines whether two sequences are equal by comparing the elements using reflect.DeepEqual.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use SequenceEqualSelf instead.
// SequenceEqual panics if 'first' or 'second' is nil.
func SequenceEqual[Source any](first, second Enumerator[Source]) bool {
	return SequenceEqualEq(first, second, nil)
}

// SequenceEqualErr is like SequenceEqual but returns an error instead of panicking.
func SequenceEqualErr[Source any](first, second Enumerator[Source]) (res bool, err error) {
	defer func() {
		if x := recover(); x != nil {
			e, ok := x.(error)
			if ok {
				res = false
				err = e
			}
		}
	}()
	return SequenceEqual(first, second), nil
}

// SequenceEqualSelf determines whether two sequences are equal by comparing the elements using reflect.DeepEqual.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// SequenceEqualSelf panics if 'first' or 'second' is nil.
func SequenceEqualSelf[Source any](first, second Enumerator[Source]) bool {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return SequenceEqual(first, NewOnSlice(sl2...))
}

// SequenceEqualSelfErr is like SequenceEqualSelf but returns an error instead of panicking.
func SequenceEqualSelfErr[Source any](first, second Enumerator[Source]) (res bool, err error) {
	defer func() {
		if x := recover(); x != nil {
			e, ok := x.(error)
			if ok {
				res = false
				err = e
			}
		}
	}()
	return SequenceEqualSelf(first, second), nil
}

// SequenceEqualEq determines whether two sequences are equal by comparing their elements using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use SequenceEqualEqSelf instead.
// SequenceEqualEq panics if 'first' or 'second' is nil.
func SequenceEqualEq[Source any](first, second Enumerator[Source], eq Equaler[Source]) bool {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if eq == nil {
		eq = EqualerFunc[Source](DeepEqual[Source])
	}
	for first.MoveNext() {
		if !second.MoveNext() {
			return false
		}
		if !eq.Equal(first.Current(), second.Current()) {
			return false
		}
	}
	if second.MoveNext() {
		return false
	}
	return true
}

// SequenceEqualEqErr is like SequenceEqualEq but returns an error instead of panicking.
func SequenceEqualEqErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res bool, err error) {
	defer func() {
		if x := recover(); x != nil {
			e, ok := x.(error)
			if ok {
				res = false
				err = e
			}
		}
	}()
	return SequenceEqualEq(first, second, eq), nil
}

// SequenceEqualEqSelf determines whether two sequences are equal by comparing their elements using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// SequenceEqualEqSelf panics if 'first' or 'second' is nil.
func SequenceEqualEqSelf[Source any](first, second Enumerator[Source], eq Equaler[Source]) bool {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return SequenceEqualEq(first, NewOnSlice(sl2...), eq)
}

// SequenceEqualEqSelfErr is like SequenceEqualEqSelf but returns an error instead of panicking.
func SequenceEqualEqSelfErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res bool, err error) {
	defer func() {
		if x := recover(); x != nil {
			e, ok := x.(error)
			if ok {
				res = false
				err = e
			}
		}
	}()
	return SequenceEqualEqSelf(first, second, eq), nil
}
