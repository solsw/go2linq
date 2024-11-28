package go2linq

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
	"github.com/solsw/iterhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/IntersectTest.cs

func TestIntersect_int(t *testing.T) {
	ii1 := iterhelper.VarSeq(1, 2, 3, 4)
	ii2 := iterhelper.VarSeq(1, 2, 3, 4)
	ii3 := iterhelper.VarSeq(1, 2, 3, 4)
	type args struct {
		first  iter.Seq[int]
		second iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "1",
			args: args{
				first:  iterhelper.VarSeq(1, 2),
				second: iterhelper.VarSeq(2, 3),
			},
			want: iterhelper.VarSeq(2),
		},
		{name: "IntWithoutComparer",
			args: args{
				first:  iterhelper.VarSeq(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: iterhelper.VarSeq(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: iterhelper.VarSeq(4, 5, 6, 7, 8),
		},
		{name: "SameEnumerable1",
			args: args{
				first:  ii1,
				second: ii1,
			},
			want: iterhelper.VarSeq(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  ii2,
				second: errorhelper.Must(Skip(ii2, 1)),
			},
			want: iterhelper.VarSeq(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  errorhelper.Must(Skip(ii3, 3)),
				second: ii3,
			},
			want: iterhelper.VarSeq(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Intersect(tt.args.first, tt.args.second)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Intersect() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestIntersect_string(t *testing.T) {
	type args struct {
		first  iter.Seq[string]
		second iter.Seq[string]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "NoComparerSpecified",
			args: args{
				first:  iterhelper.VarSeq("A", "a", "b", "c", "b"),
				second: iterhelper.VarSeq("b", "a", "d", "a"),
			},
			want: iterhelper.VarSeq("a", "b"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#intersect-and-intersectby
		{name: "Intersect",
			args: args{
				first:  iterhelper.VarSeq("Mercury", "Venus", "Earth", "Jupiter"),
				second: iterhelper.VarSeq("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: iterhelper.VarSeq("Mercury", "Earth", "Jupiter"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Intersect(tt.args.first, tt.args.second)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Intersect() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestIntersectEq_int(t *testing.T) {
	type args struct {
		first  iter.Seq[int]
		second iter.Seq[int]
		equal  func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:  iterhelper.VarSeq(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: iterhelper.VarSeq(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				equal:  generichelper.DeepEqual[int]},
			want: iterhelper.VarSeq(4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IntersectEq(tt.args.first, tt.args.second, tt.args.equal)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("IntersectEq() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestIntersectEq_string(t *testing.T) {
	type args struct {
		first  iter.Seq[string]
		second iter.Seq[string]
		equal  func(string, string) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:  iterhelper.VarSeq("A", "a", "b", "c", "b"),
				second: iterhelper.VarSeq("b", "a", "d", "a"),
				equal:  caseInsensitiveEqual,
			},
			want: iterhelper.VarSeq("A", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IntersectEq(tt.args.first, tt.args.second, tt.args.equal)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("IntersectEq() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestIntersectCmp_int(t *testing.T) {
	ii1 := iterhelper.VarSeq(4, 3, 2, 1)
	ii2 := iterhelper.VarSeq(1, 2, 3, 4)
	ii3 := iterhelper.VarSeq(1, 2, 3, 4)
	type args struct {
		first   iter.Seq[int]
		second  iter.Seq[int]
		compare func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "IntComparerSpecified",
			args: args{
				first:   iterhelper.VarSeq(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second:  iterhelper.VarSeq(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				compare: cmp.Compare[int],
			},
			want: iterhelper.VarSeq(4, 5, 6, 7, 8),
		},
		{name: "SameEnumerable1",
			args: args{
				first:   ii1,
				second:  ii1,
				compare: cmp.Compare[int],
			},
			want: iterhelper.VarSeq(4, 3, 2, 1),
		},
		{name: "SameEnumerable2",
			args: args{
				first:   ii2,
				second:  errorhelper.Must(Skip(ii2, 1)),
				compare: cmp.Compare[int],
			},
			want: iterhelper.VarSeq(2, 3, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:   errorhelper.Must(Skip(ii3, 3)),
				second:  ii3,
				compare: cmp.Compare[int],
			},
			want: iterhelper.VarSeq(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IntersectCmp(tt.args.first, tt.args.second, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("IntersectCmp() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestIntersectCmp_string(t *testing.T) {
	type args struct {
		first   iter.Seq[string]
		second  iter.Seq[string]
		compare func(string, string) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:   iterhelper.VarSeq("A", "a", "b", "c", "b"),
				second:  iterhelper.VarSeq("b", "a", "d", "a"),
				compare: caseInsensitiveCompare,
			},
			want: iterhelper.VarSeq("A", "b"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IntersectCmp(tt.args.first, tt.args.second, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("IntersectCmp() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect
func ExampleIntersect() {
	id1 := iterhelper.VarSeq(44, 26, 92, 30, 71, 38)
	id2 := iterhelper.VarSeq(39, 59, 83, 47, 26, 4, 30)
	intersect, _ := Intersect(id1, id2)
	for id := range intersect {
		fmt.Println(id)
	}
	// Output:
	// 26
	// 30
}

// second and third examples from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect
func ExampleIntersectEq() {
	store1 := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
	}
	store2 := []Product{
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
	}
	// Get the products from the first array that have duplicates in the second array.
	equal := func(p1, p2 Product) bool {
		return p1.Name == p2.Name && p1.Code == p2.Code
	}
	intersectEq, _ := IntersectEq(slices.Values(store1), slices.Values(store2), equal)
	for product := range intersectEq {
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// apple 9
}
