//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 5 - Empty
// https://codeblog.jonskeet.uk/2010/12/24/reimplementing-linq-to-objects-part-5-empty/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.empty

// Empty returns an empty Enumerator that has the specified type argument.
func Empty[Result any]() Enumerator[Result] {
	return OnFunc[Result]{}
}
