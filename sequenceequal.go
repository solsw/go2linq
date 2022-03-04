//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 34 - SequenceEqual
// https://codeblog.jonskeet.uk/2011/01/14/reimplementing-linq-to-objects-part-34-sequenceequal/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

// SequenceEqual determines whether two sequences are equal by comparing the elements using DeepEqual.
func SequenceEqual[Source any](first, second Enumerable[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	return SequenceEqualEq(first, second, nil)
}

// SequenceEqualMust is like SequenceEqual but panics in case of error.
func SequenceEqualMust[Source any](first, second Enumerable[Source]) bool {
	r, err := SequenceEqual(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// SequenceEqualEq determines whether two sequences are equal by comparing their elements using a specified Equaler.
// If 'equaler' is nil DeepEqual is used.
func SequenceEqualEq[Source any](first, second Enumerable[Source], equaler Equaler[Source]) (res bool, err error) {
	if first == nil || second == nil {
		return false, ErrNilSource
	}
	counter1, ok1 := first.(Counter)
	if ok1 {
		counter2, ok2 := second.(Counter)
		if ok2 && (counter1.Count() != counter2.Count()) {
			return false, nil
		}
	}
	if equaler == nil {
		equaler = DeepEqual[Source]{}
	}
	enr1 := first.GetEnumerator()
	enr2 := second.GetEnumerator()
	for enr1.MoveNext() {
		if !enr2.MoveNext() {
			return false, nil
		}
		if !equaler.Equal(enr1.Current(), enr2.Current()) {
			return false, nil
		}
	}
	if enr2.MoveNext() {
		return false, nil
	}
	return true, nil
}

// SequenceEqualEqMust is like SequenceEqualEq but panics in case of error.
func SequenceEqualEqMust[Source any](first, second Enumerable[Source], equaler Equaler[Source]) bool {
	r, err := SequenceEqualEq(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
