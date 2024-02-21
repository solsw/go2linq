package go2linq

import (
	"cmp"
	"iter"
	"sort"
)

func orderByLsPrim[Source any](source iter.Seq[Source], less func(x, y Source) bool) iter.Seq[Source] {
	ss, _ := ToSlice(source)
	sort.SliceStable(ss, func(i, j int) bool {
		return less(ss[i], ss[j])
	})
	return SliceAll(ss)
}

// [OrderBy] sorts the elements of a sequence in ascending order.
//
// [OrderBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderby
func OrderBy[Source cmp.Ordered](source iter.Seq[Source]) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return orderByLsPrim(source, cmp.Less), nil
}

// [OrderByLs] sorts the elements of a sequence in ascending order using a specified 'less' function.
//
// [OrderByLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderby
func OrderByLs[Source any](source iter.Seq[Source], less func(x, y Source) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if less == nil {
		return nil, ErrNilLess
	}
	return orderByLsPrim(source, less), nil
}

// [OrderByDesc] sorts the elements of a sequence in descending order.
//
// [OrderByDesc]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
func OrderByDesc[Source cmp.Ordered](source iter.Seq[Source]) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return orderByLsPrim(source, reverseLess[Source](cmp.Less)), nil
}

// [OrderByDescLs] sorts the elements of a sequence in descending order using a specified 'less' function.
//
// [OrderByDescLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
func OrderByDescLs[Source any](source iter.Seq[Source], less func(x, y Source) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if less == nil {
		return nil, ErrNilLess
	}
	return orderByLsPrim(source, reverseLess[Source](less)), nil
}
