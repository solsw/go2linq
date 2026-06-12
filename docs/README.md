# go2linq
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/go2linq.svg)](https://pkg.go.dev/github.com/solsw/go2linq/v4)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/go2linq)

**go2linq** is a Go implementation of .NET's
[LINQ to Objects](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects).
It brings the familiar LINQ query operators — `Where`, `Select`, `GroupBy`,
`Join`, `OrderBy`, `Aggregate` and dozens more — to Go, implemented as generic
functions over sequences.

See also:
[Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[LINQ](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/),
[Enumerable Class](https://learn.microsoft.com/dotnet/api/system.linq.enumerable).

## Overview

**v4** of **go2linq** is built on Go's standard
[iterators](https://pkg.go.dev/iter#hdr-Iterators): a *sequence* is simply an
[`iter.Seq[T]`](https://pkg.go.dev/iter#Seq). Because operators consume and
produce plain `iter.Seq` values, go2linq composes naturally with the standard
library ([`slices`](https://pkg.go.dev/slices),
[`maps`](https://pkg.go.dev/maps)) and with any other code that speaks the
`iter.Seq` protocol — no custom collection type is required.

```go
package main

import (
	"fmt"

	"github.com/solsw/go2linq/v4"
	"github.com/solsw/iterhelper"
)

func main() {
	filter, _ := go2linq.Where(
		iterhelper.Var(1, 2, 3, 4, 5, 6, 7, 8),
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

Output:

```
2: 4
4: 16
6: 36
7: 49
8: 64
```

## Installation

```
go get github.com/solsw/go2linq/v4
```

```go
import "github.com/solsw/go2linq/v4"
```

Requires **Go 1.26** or later (for iterator support).

## Design

### Every operator returns an error

Each operator returns `(result, error)` rather than panicking. The error is
non-`nil` when an argument is invalid — a `nil` source, or a `nil`
predicate/selector/comparison function — and for operators that can fail at
evaluation time (e.g. `Single` on a sequence with more than one element, or
`ElementAt` with an out-of-range index). All error values are exported from the
package (see [Errors](#errors)), so they can be matched with `errors.Is`.

### Deferred vs. immediate execution

Following LINQ semantics, operators fall into two groups:

- **Deferred** operators (`Where`, `Select`, `Concat`, `OrderBy`, …) return a
  new `iter.Seq[T]` and do no work until the result is ranged over. They can be
  chained freely; nothing is materialized until you iterate.
- **Immediate** operators (`Count`, `First`, `Aggregate`, `ToMap`,
  `SequenceEqual`, …) consume the source and return a final value right away.

### Naming conventions

go2linq exposes a base operator plus suffixed variants. The suffixes signal what
extra argument the variant takes:

| Suffix  | Meaning                                                              | Example                       |
|---------|---------------------------------------------------------------------|-------------------------------|
| `Idx`   | the element's index is passed to the predicate/selector             | `WhereIdx`, `SelectIdx`       |
| `Pred`  | takes a predicate to filter which elements qualify                  | `FirstPred`, `CountPred`      |
| `Sel`   | takes a selector that projects each element                         | `AverageSel`, `MaxSel`        |
| `By`    | groups/compares elements by a key selector                          | `DistinctBy`, `ExceptBy`      |
| `Res`   | takes a result selector to shape the output                         | `GroupByRes`                  |
| `Eq`    | takes an equality function `func(T, T) bool`                         | `DistinctEq`, `ContainsEq`    |
| `Cmp`   | takes a comparison function `func(T, T) int`                         | `DistinctCmp`, `UnionCmp`     |
| `Ls`    | takes a "less" function `func(T, T) bool`                            | `OrderByLs`, `MaxLs`          |
| `Def`   | takes an explicit default value                                     | `DefaultIfEmptyDef`           |

Operators without an `Eq`/`Cmp`/`Ls` suffix use Go's built-in comparison
(`comparable` / `cmp.Ordered`) where one is required.

## Operators

### Filtering
`Where`, `WhereIdx`, `OfType`, `Cast`

### Projection
`Select`, `SelectIdx`, `SelectMany`, `SelectManyIdx`, `SelectManyColl`,
`SelectManyCollIdx`, `Zip`

### Partitioning
`Skip`, `SkipLast`, `SkipWhile`, `SkipWhileIdx`,
`Take`, `TakeLast`, `TakeWhile`, `TakeWhileIdx`, `Chunk`

### Concatenation
`Concat`, `ConcatMany`, `Append`, `Prepend`

### Ordering
`OrderBy`, `OrderByLs`, `OrderByDesc`, `OrderByDescLs`,
`OrderByKey`, `OrderByKeyLs`, `OrderByKeyDesc`, `OrderByKeyDescLs`,
`ThenLess`, `Reverse`, `ReverseLess`

### Grouping
`GroupBy`, `GroupByEq`, `GroupBySel`, `GroupBySelEq`,
`GroupByRes`, `GroupByResEq`, `GroupBySelRes`, `GroupBySelResEq`

### Joining
`Join`, `JoinEq`, `GroupJoin`, `GroupJoinEq`

### Set operations
`Distinct`, `DistinctCmp`, `DistinctEq`, `DistinctBy`, `DistinctByCmp`, `DistinctByEq`,
`Except`, `ExceptCmp`, `ExceptEq`, `ExceptBy`, `ExceptByCmp`, `ExceptByEq`,
`Intersect`, `IntersectCmp`, `IntersectEq`, `IntersectBy`, `IntersectByCmp`, `IntersectByEq`,
`Union`, `UnionCmp`, `UnionEq`, `UnionBy`, `UnionByCmp`, `UnionByEq`

### Aggregation
`Aggregate`, `AggregateSeed`, `AggregateSeedSel`,
`AggregateBy`, `AggregateByEq`, `AggregateBySel`, `AggregateBySelEq`,
`Average`, `AverageSel`, `Sum`, `SumSel`,
`Count`, `CountPred`, `CountBy`, `CountByEq`,
`Max`, `MaxLs`, `MaxSel`, `MaxSelLs`, `MaxBySel`, `MaxBySelLs`,
`Min`, `MinLs`, `MinSel`, `MinSelLs`, `MinBySel`, `MinBySelLs`

### Quantifiers
`All`, `Any`, `AnyPred`, `Contains`, `ContainsEq`, `SequenceEqual`, `SequenceEqualEq`

### Element operators
`First`, `FirstPred`, `FirstOrDefault`, `FirstOrDefaultPred`,
`Last`, `LastPred`, `LastOrDefault`, `LastOrDefaultPred`,
`Single`, `SinglePred`, `SingleOrDefault`, `SingleOrDefaultPred`,
`SingleOrZero`, `SingleOrZeroPred`,
`ElementAt`, `ElementAtOrDefault`, `DefaultIfEmpty`, `DefaultIfEmptyDef`

### Generation
`Empty`, `Range`, `Repeat`, `InfiniteSequence`

### Conversion & output
`ToMap`, `ToMapSel`,
`ToLookup`, `ToLookupEq`, `ToLookupSel`, `ToLookupSelEq`,
`Index`, `ApplyResultSelector`

### Helpers
`Identity` — selector that returns its argument unchanged.

For full signatures and per-operator documentation, see the
[Go Reference](https://pkg.go.dev/github.com/solsw/go2linq/v4).

## Errors

All errors returned by the package are exported and can be tested with
`errors.Is`:

| Error                 | Meaning                                                    |
|-----------------------|------------------------------------------------------------|
| `ErrNilSource`        | the source sequence is `nil`                               |
| `ErrNilPredicate`     | a required predicate function is `nil`                     |
| `ErrNilSelector`      | a required selector function is `nil`                      |
| `ErrNilAccumulator`   | a required accumulator function is `nil`                   |
| `ErrNilEqual`         | a required equality function is `nil`                      |
| `ErrNilCompare`       | a required comparison function is `nil`                    |
| `ErrNilLess`          | a required "less" function is `nil`                        |
| `ErrNilNext`          | a required "next" function is `nil`                        |
| `ErrEmptySource`      | the source is empty where a value is required              |
| `ErrMultipleElements` | `Single` found more than one element                       |
| `ErrMultipleMatch`    | `SinglePred` matched more than one element                 |
| `ErrNoMatch`          | a predicate matched no element                             |
| `ErrIndexOutOfRange`  | the requested index is out of range                        |
| `ErrNegativeCount`    | a negative count was supplied                              |
| `ErrSizeOutOfRange`   | a size argument (e.g. `Chunk`) is out of range             |
| `ErrDuplicateKeys`    | `ToMap` encountered duplicate keys                         |

## Examples

Runnable examples accompany every operator as `Example...` functions in the test
files; they are rendered in the
[package documentation](https://pkg.go.dev/github.com/solsw/go2linq/v4#pkg-examples).

## License

**go2linq** is distributed under the terms of the [LICENSE](https://github.com/solsw/go2linq?tab=MIT-1-ov-file#MIT-1-ov-file) file in the repository root.
