//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/IntersectTest.cs

func Test_IntersectMust_int(t *testing.T) {
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
		{name: "1",
			args: args{
				first:  NewEnSlice(1, 2),
				second: NewEnSlice(2, 3),
			},
			want: NewEnSlice(2),
		},
		{name: "IntWithoutComparer",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: NewEnSlice(4, 5, 6, 7, 8),
		},
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  e2,
				second: SkipMust(e2, 1),
			},
			want: NewEnSlice(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 3),
				second: e3,
			},
			want: NewEnSlice(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_IntersectMust_string(t *testing.T) {
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "NoComparerSpecified",
			args: args{
				first:  NewEnSlice("A", "a", "b", "c", "b"),
				second: NewEnSlice("b", "a", "d", "a"),
			},
			want: NewEnSlice("a", "b"),
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#intersect-and-intersectby
		{name: "Intersect",
			args: args{
				first:  NewEnSlice("Mercury", "Venus", "Earth", "Jupiter"),
				second: NewEnSlice("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: NewEnSlice("Mercury", "Earth", "Jupiter"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_IntersectEqMust_int(t *testing.T) {
	type args struct {
		first   Enumerable[int]
		second  Enumerable[int]
		equaler Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:   NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second:  NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				equaler: Order[int]{}},
			want: NewEnSlice(4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectEqMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_IntersectEqMust_string(t *testing.T) {
	type args struct {
		first   Enumerable[string]
		second  Enumerable[string]
		equaler Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:   NewEnSlice("A", "a", "b", "c", "b"),
				second:  NewEnSlice("b", "a", "d", "a"),
				equaler: CaseInsensitiveEqualer,
			},
			want: NewEnSlice("A", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectEqMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_IntersectCmpMust_int(t *testing.T) {
	e1 := NewEnSlice(4, 3, 2, 1)
	e2 := NewEnSlice(1, 2, 3, 4)
	e3 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
		cmp    Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				cmp:    Order[int]{},
			},
			want: NewEnSlice(4, 5, 6, 7, 8),
		},
		{name: "SameEnumerable1",
			args: args{
				first:  e1,
				second: e1,
				cmp:    Order[int]{},
			},
			want: NewEnSlice(4, 3, 2, 1),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  e2,
				second: SkipMust(e2, 1),
				cmp:    Order[int]{},
			},
			want: NewEnSlice(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  SkipMust(e3, 3),
				second: e3,
				cmp:    Order[int]{},
			},
			want: NewEnSlice(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectCmpMust(tt.args.first, tt.args.second, tt.args.cmp)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_IntersectCmpMust_string(t *testing.T) {
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
		cmp    Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:  NewEnSlice("A", "a", "b", "c", "b"),
				second: NewEnSlice("b", "a", "d", "a"),
				cmp:    CaseInsensitiveComparer,
			},
			want: NewEnSlice("A", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectCmpMust(tt.args.first, tt.args.second, tt.args.cmp)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("IntersectCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the first example from Enumerable.Intersect help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect
func ExampleIntersectMust() {
	id1 := NewEnSlice(44, 26, 92, 30, 71, 38)
	id2 := NewEnSlice(39, 59, 83, 47, 26, 4, 30)
	intersect := IntersectMust(id1, id2)
	enr := intersect.GetEnumerator()
	for enr.MoveNext() {
		id := enr.Current()
		fmt.Println(id)
	}
	// Output:
	// 26
	// 30
}

// see the second and third examples from Enumerable.Intersect help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect
func ExampleIntersectEqMust() {
	store1 := NewEnSlice(
		Product{Name: "apple", Code: 9},
		Product{Name: "orange", Code: 4},
	)
	store2 := NewEnSlice(
		Product{Name: "apple", Code: 9},
		Product{Name: "lemon", Code: 12},
	)
	// Get the products from the first array that have duplicates in the second array.
	var equaler Equaler[Product] = EqualerFunc[Product](
		func(p1, p2 Product) bool {
			return p1.Name == p2.Name && p1.Code == p2.Code
		},
	)
	intersectEq := IntersectEqMust(store1, store2, equaler)
	enr := intersectEq.GetEnumerator()
	for enr.MoveNext() {
		product := enr.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// apple 9
}
