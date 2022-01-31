go2linq
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/go2linq.svg)](https://pkg.go.dev/github.com/solsw/go2linq)
=======

**go2linq** is Go implementation of .NET's 
[LINQ to Objects](https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects).
(See also: [Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[LINQ](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/),
[Enumerable Class](https://docs.microsoft.com/dotnet/api/system.linq.enumerable).)

**go2linq** is inspired by Jon Skeet's [Edulinq series](https://codeblog.jonskeet.uk/category/edulinq/).

Since **go2linq** uses generics it requires at least Go 1.18.
Use [gotip tool](https://pkg.go.dev/golang.org/dl/gotip) or install [go1.18beta1](https://go.dev/dl/#go1.18beta1) to experiment with **go2linq**.

---

## Simple example

```go
//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

func main() {
	filter := go2linq.WhereMust(
		go2linq.NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8),
		func(i int) bool { return i > 6 || i%2 == 0 },
	)
	squares := go2linq.SelectMust(
		filter,
		func(i int) string { return fmt.Sprintf("%d: %d", i, i*i) },
	)
	enr := squares.GetEnumerator()
	for enr.MoveNext() {
		square := enr.Current()
		fmt.Println(square)
	}
}
```

The previous code prints the following:
```
2: 4
4: 16
6: 36
7: 49
8: 64
```
