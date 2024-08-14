package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [SequenceEqual] determines whether two sequences are equal by comparing the elements using [generichelper.DeepEqual].
//
// [SequenceEqual]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func SequenceEqual[Source any](first, second iter.Seq[Source]) (bool, error) {
	if first == nil || second == nil {
		return false, errorhelper.CallerError(ErrNilSource)
	}
	r, err := SequenceEqualEq(first, second, generichelper.DeepEqual[Source])
	if err != nil {
		return false, errorhelper.CallerError(err)
	}
	return r, nil
}

// [SequenceEqualEq] determines whether two sequences are equal by comparing their elements using a specified 'equal'.
//
// [SequenceEqualEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func SequenceEqualEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (bool, error) {
	if first == nil || second == nil {
		return false, errorhelper.CallerError(ErrNilSource)
	}
	if equal == nil {
		return false, errorhelper.CallerError(ErrNilEqual)
	}
	next1, stop1 := iter.Pull(first)
	defer stop1()
	next2, stop2 := iter.Pull(second)
	defer stop2()
	for {
		s1, ok1 := next1()
		s2, ok2 := next2()
		if ok1 != ok2 {
			return false, nil
		}
		// here ok1 and ok2 are either both true or both false
		if !ok1 {
			break
		}
		if !equal(s1, s2) {
			return false, nil
		}
	}
	return true, nil
}

// SequenceEqual2 determines whether two sequence2s are equal by comparing the elements using [generichelper.DeepEqual].
func SequenceEqual2[K, V any](first, second iter.Seq2[K, V]) (bool, error) {
	if first == nil || second == nil {
		return false, errorhelper.CallerError(ErrNilSource)
	}
	r, err := SequenceEqual2Eq(first, second, generichelper.DeepEqual[K], generichelper.DeepEqual[V])
	if err != nil {
		return false, errorhelper.CallerError(err)
	}
	return r, nil
}

// SequenceEqual2Eq determines whether two sequence2s are equal by comparing their elements using specified equals.
func SequenceEqual2Eq[K, V any](first, second iter.Seq2[K, V], equalK func(K, K) bool, equalV func(V, V) bool) (bool, error) {
	if first == nil || second == nil {
		return false, errorhelper.CallerError(ErrNilSource)
	}
	if equalK == nil || equalV == nil {
		return false, errorhelper.CallerError(ErrNilEqual)
	}
	next1, stop1 := iter.Pull2(first)
	defer stop1()
	next2, stop2 := iter.Pull2(second)
	defer stop2()
	for {
		k1, v1, ok1 := next1()
		k2, v2, ok2 := next2()
		if ok1 != ok2 {
			return false, nil
		}
		// here ok1 and ok2 are either both true or both false
		if !ok1 {
			break
		}
		if !(equalK(k1, k2) && equalV(v1, v2)) {
			return false, nil
		}
	}
	return true, nil
}
