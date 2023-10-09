# go2linq
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/go2linq.svg)](https://pkg.go.dev/github.com/solsw/go2linq/v3)

**go2linq** is Go implementation of .NET's 
[LINQ to Objects](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects).
(See also: [Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[LINQ](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/),
[Enumerable Class](https://learn.microsoft.com/dotnet/api/system.linq.enumerable).)

**go2linq** was initially inspired by Jon Skeet's [Edulinq series](https://codeblog.jonskeet.uk/category/edulinq/).

---

## Installation

```
go get github.com/solsw/go2linq/v3
```

## Examples

Examples of **go2linq** usage are in the `Example...` functions in test files
(see [Examples](https://pkg.go.dev/github.com/solsw/go2linq/v3#pkg-examples)).

### Quick and easy example:

```go
package main

import (
	"fmt"

	"github.com/solsw/go2linq/v3"
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

The previous code outputs the following:
```
2: 4
4: 16
6: 36
7: 49
8: 64
```
