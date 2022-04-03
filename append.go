//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.append

// Append appends a value to the end of the sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.append)
func Append[Source any](source Enumerable[Source], element Source) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return Concat(source, RepeatMust(element, 1))
}

// AppendMust is like Append but panics in case of an error.
func AppendMust[Source any](source Enumerable[Source], element Source) Enumerable[Source] {
	r, err := Append(source, element)
	if err != nil {
		panic(err)
	}
	return r
}
