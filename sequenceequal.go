//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 34 - SequenceEqual
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-34-sequenceequal/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

// SequenceEqual determines whether two sequences are equal by comparing the elements using reflect.DeepEqual.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use SequenceEqualSelf instead.
func SequenceEqual[Source any](first, second Enumerator[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	return SequenceEqualEq(first, second, nil)
}

// SequenceEqualMust is like SequenceEqual but panics in case of error.
func SequenceEqualMust[Source any](first, second Enumerator[Source]) bool {
	r, err := SequenceEqual(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// SequenceEqualSelf determines whether two sequences are equal by comparing the elements using reflect.DeepEqual.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func SequenceEqualSelf[Source any](first, second Enumerator[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return SequenceEqual(first, NewOnSlice(sl2...))
}

// SequenceEqualSelfMust is like SequenceEqualSelf but panics in case of error.
func SequenceEqualSelfMust[Source any](first, second Enumerator[Source]) bool {
	r, err := SequenceEqualSelf(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// SequenceEqualEq determines whether two sequences are equal by comparing their elements using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use SequenceEqualEqSelf instead.
func SequenceEqualEq[Source any](first, second Enumerator[Source], eq Equaler[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	if eq == nil {
		eq = EqualerFunc[Source](DeepEqual[Source])
	}
	for first.MoveNext() {
		if !second.MoveNext() {
			return false, nil
		}
		if !eq.Equal(first.Current(), second.Current()) {
			return false, nil
		}
	}
	if second.MoveNext() {
		return false, nil
	}
	return true, nil
}

// SequenceEqualEqMust is like SequenceEqualEq but panics in case of error.
func SequenceEqualEqMust[Source any](first, second Enumerator[Source], eq Equaler[Source]) bool {
	r, err := SequenceEqualEq(first, second, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// SequenceEqualEqSelf determines whether two sequences are equal by comparing their elements using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func SequenceEqualEqSelf[Source any](first, second Enumerator[Source], eq Equaler[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return SequenceEqualEq(first, NewOnSlice(sl2...), eq)
}

// SequenceEqualEqSelfMust is like SequenceEqualEqSelf but panics in case of error.
func SequenceEqualEqSelfMust[Source any](first, second Enumerator[Source], eq Equaler[Source]) bool {
	r, err := SequenceEqualEqSelf(first, second, eq)
	if err != nil {
		panic(err)
	}
	return r
}
