package go2linq

import (
	"fmt"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/UnionTest.cs

func TestUnionMust_string(t *testing.T) {
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "UnionWithTwoEmptySequences",
			args: args{
				first:  Empty[string](),
				second: Empty[string](),
			},
			want: Empty[string](),
		},
		{name: "FirstEmpty",
			args: args{
				first:  Empty[string](),
				second: NewEnSlice("one", "two", "three", "four"),
			},
			want: NewEnSlice("one", "two", "three", "four"),
		},
		{name: "SecondEmpty",
			args: args{
				first:  NewEnSlice("one", "two", "three", "four"),
				second: Empty[string](),
			},
			want: NewEnSlice("one", "two", "three", "four"),
		},
		{name: "UnionWithoutComparer",
			args: args{
				first:  NewEnSlice("a", "b", "B", "c", "b"),
				second: NewEnSlice("d", "e", "d", "a"),
			},
			want: NewEnSlice("a", "b", "B", "c", "d", "e"),
		},
		{name: "UnionWithoutComparer2",
			args: args{
				first:  NewEnSlice("a", "b"),
				second: NewEnSlice("b", "a"),
			},
			want: NewEnSlice("a", "b"),
		},
		{name: "UnionWithEmptyFirstSequence",
			args: args{
				first:  Empty[string](),
				second: NewEnSlice("d", "e", "d", "a"),
			},
			want: NewEnSlice("d", "e", "a"),
		},
		{name: "UnionWithEmptySecondSequence",
			args: args{
				first:  NewEnSlice("a", "b", "B", "c", "b"),
				second: Empty[string](),
			},
			want: NewEnSlice("a", "b", "B", "c"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#union-and-unionby
		{name: "Union",
			args: args{
				first:  NewEnSlice("Mercury", "Venus", "Earth", "Jupiter"),
				second: NewEnSlice("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: NewEnSlice("Mercury", "Venus", "Earth", "Jupiter", "Mars"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestUnionMust_int(t *testing.T) {
	e1 := NewEnSlice(1, 2, 3, 4)
	e2 := NewEnSlice(1, 2, 3, 4)
	e3 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  TakeMust(e2, 1),
				second: SkipMust(e2, 3),
			},
			want: NewEnSlice(1, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 2),
				second: e3,
			},
			want: NewEnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestUnionEqMust_int(t *testing.T) {
	type args struct {
		first   Enumerable[int]
		second  Enumerable[int]
		equaler collate.Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "UnionWithIntEquality",
			args: args{
				first:   NewEnSlice(1, 2),
				second:  NewEnSlice(2, 3),
				equaler: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionEqMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestUnionEqMust_string(t *testing.T) {
	type args struct {
		first   Enumerable[string]
		second  Enumerable[string]
		equaler collate.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "UnionWithCaseInsensitiveComparerEq",
			args: args{
				first:   NewEnSlice("a", "b", "B", "c", "b"),
				second:  NewEnSlice("d", "e", "d", "a"),
				equaler: collate.CaseInsensitiveOrder,
			},
			want: NewEnSlice("a", "b", "c", "d", "e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnionEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionEqMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestUnionCmpMust_int(t *testing.T) {
	e1 := NewEnSlice(1, 2, 3, 4)
	e2 := NewEnSlice(1, 2, 3, 4)
	e3 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first    Enumerable[int]
		second   Enumerable[int]
		comparer collate.Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "UnionWithIntComparer1",
			args: args{
				first:    NewEnSlice(1, 2, 2),
				second:   Empty[int](),
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2),
		},
		{name: "UnionWithIntComparer2",
			args: args{
				first:    NewEnSlice(1, 2),
				second:   NewEnSlice(2, 3),
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "SameEnumerable1",
			args: args{
				first:    e1,
				second:   e1,
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:    SkipMust(e2, 2),
				second:   TakeMust(e2, 1),
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(3, 4, 1),
		},
		{name: "SameEnumerable3",
			args: args{
				first:    SkipMust(e3, 2),
				second:   e3,
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestUnionCmpMust_string(t *testing.T) {
	type args struct {
		first    Enumerable[string]
		second   Enumerable[string]
		comparer collate.Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "UnionWithCaseInsensitiveComparerCmp",
			args: args{
				first:    NewEnSlice("a", "b", "B", "c", "b"),
				second:   NewEnSlice("d", "e", "d", "a"),
				comparer: collate.CaseInsensitiveOrder,
			},
			want: NewEnSlice("a", "b", "c", "d", "e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionCmpMust(tt.args.first, tt.args.second, tt.args.comparer); !SequenceEqualMust(got, tt.want) {
				t.Errorf("UnionCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the first example from Enumerable.Union help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func ExampleUnionMust() {
	ints1 := []int{5, 3, 9, 7, 5, 9, 3, 7}
	ints2 := []int{8, 3, 6, 4, 4, 9, 1, 0}
	union := UnionMust(
		NewEnSlice(ints1...),
		NewEnSlice(ints2...),
	)
	enr := union.GetEnumerator()
	for enr.MoveNext() {
		num := enr.Current()
		fmt.Printf("%d ", num)
	}
	// Output:
	// 5 3 9 7 8 6 4 1 0
}

// see the last example from Enumerable.Union help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func ExampleUnionEqMust() {
	store1 := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
	}
	store2 := []Product{
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
	}
	//Get the products from the both arrays excluding duplicates.
	var equaler collate.Equaler[Product] = collate.EqualerFunc[Product](
		func(p1, p2 Product) bool {
			return p1.Code == p2.Code && p1.Name == p2.Name
		},
	)
	unionEq := UnionEqMust(
		NewEnSlice(store1...),
		NewEnSlice(store2...),
		equaler,
	)
	enr := unionEq.GetEnumerator()
	for enr.MoveNext() {
		product := enr.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// apple 9
	// orange 4
	// lemon 12
}
