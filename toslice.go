//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 20 – ToList
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-20-tolist/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolist
// Reimplementing LINQ to Objects: Part 24 – ToArray
// https://codeblog.jonskeet.uk/2011/01/02/reimplementing-linq-to-objects-part-24-toarray/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.toarray

// ToSlice creates a slice from an Enumerable.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.tolist,
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.toarray)
func ToSlice[Source any](source Enumerable[Source]) ([]Source, error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if slicer, ok := source.(Slicer[Source]); ok {
		return slicer.Slice(), nil
	}
	return enrToSlice(source.GetEnumerator()), nil
}

// ToSliceMust is like ToSlice but panics in case of error.
func ToSliceMust[Source any](source Enumerable[Source]) []Source {
	r, err := ToSlice(source)
	if err != nil {
		panic(err)
	}
	return r
}
