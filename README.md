go2linq
=======

**go2linq** is [generics-based](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
 Golang implementation of .NET's 
[LINQ to Objects](https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects).
(See also: [Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
 [LINQ](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/),
 [Enumerable Class](https://docs.microsoft.com/dotnet/api/system.linq.enumerable).
 **go2linq** is inspired by Jon Skeet's [Edulinq series](https://codeblog.jonskeet.uk/category/edulinq/).)

Since **go2linq** uses generics it cannot be compiled by [current version](https://golang.org/dl/) of Go.
 Use [go2go tool](https://blog.golang.org/generics-next-step) instead to experiment with **go2linq**.