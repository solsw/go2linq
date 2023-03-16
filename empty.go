package go2linq

// Reimplementing LINQ to Objects: Part 5 - Empty
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-5-empty/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.empty

// [Empty] returns an empty [Enumerable] that has the specified type argument.
//
// [Empty]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.empty
func Empty[Result any]() Enumerable[Result] {
	return OnFactory[Result](nil)
}
