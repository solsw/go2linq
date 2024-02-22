# go2linq
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/go2linq.svg)](https://pkg.go.dev/github.com/solsw/go2linq/v4)

**go2linq v4** is Go implementation of .NET's 
[LINQ to Objects](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects).
(See also: [Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[LINQ](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/),
[Enumerable Class](https://learn.microsoft.com/dotnet/api/system.linq.enumerable).)

**go2linq v4** is based on [iter.Seq](https://go.dev/wiki/RangefuncExperiment), so it requires setting *GOEXPERIMENT=rangefunc* when executing **go** commands.

---

## Installation

```
go get github.com/solsw/go2linq/v4
```

## Examples

Examples of **go2linq** usage are in the `Example...` functions in test files
(see [Examples](https://pkg.go.dev/github.com/solsw/go2linq/v4#pkg-examples)).

### Quick and easy example:

```go
package main

import (
	"fmt"

	"github.com/solsw/go2linq/v4"
)

func main() {
	filter, _ := go2linq.Where(
		go2linq.VarAll(1, 2, 3, 4, 5, 6, 7, 8),
		func(i int) bool { return i > 6 || i%2 == 0 },
	)
	squares, _ := go2linq.Select(
		filter,
		func(i int) string { return fmt.Sprintf("%d: %d", i, i*i) },
	)
	for square := range squares {
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
