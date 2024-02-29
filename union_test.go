package go2linq

import (
	"cmp"
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/UnionTest.cs

func TestUnion_string(t *testing.T) {
	type args struct {
		first  iter.Seq[string]
		second iter.Seq[string]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
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
				second: VarAll("one", "two", "three", "four", "two", "three"),
			},
			want: VarAll("one", "two", "three", "four"),
		},
		{name: "SecondEmpty",
			args: args{
				first:  VarAll("one", "two", "three", "four", "three", "four"),
				second: Empty[string](),
			},
			want: VarAll("one", "two", "three", "four"),
		},
		{name: "UnionWithoutComparer",
			args: args{
				first:  VarAll("a", "b", "B", "c", "b"),
				second: VarAll("d", "e", "d", "a"),
			},
			want: VarAll("a", "b", "B", "c", "d", "e"),
		},
		{name: "UnionWithoutComparer2",
			args: args{
				first:  VarAll("a", "b"),
				second: VarAll("b", "a"),
			},
			want: VarAll("a", "b"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#union-and-unionby
		{name: "Union",
			args: args{
				first:  VarAll("Mercury", "Venus", "Earth", "Jupiter"),
				second: VarAll("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: VarAll("Mercury", "Venus", "Earth", "Jupiter", "Mars"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Union(tt.args.first, tt.args.second)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Union() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestUnion_int(t *testing.T) {
	ii1 := VarAll(1, 2, 3, 4)
	ii2 := VarAll(1, 2, 3, 4)
	ii3 := VarAll(1, 2, 3, 4)
	type args struct {
		first  iter.Seq[int]
		second iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "SameEnumerable1",
			args: args{
				first:  ii1,
				second: ii1,
			},
			want: VarAll(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:  errorhelper.Must(Take(ii2, 1)),
				second: errorhelper.Must(Skip(ii2, 3)),
			},
			want: VarAll(1, 4),
		},
		{name: "SameEnumerable3",
			args: args{
				first:  errorhelper.Must(Skip(ii3, 2)),
				second: ii3,
			},
			want: VarAll(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Union(tt.args.first, tt.args.second)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Union() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestUnionEq_int(t *testing.T) {
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
		{name: "UnionWithIntEquality",
			args: args{
				first:  VarAll(1, 2),
				second: VarAll(2, 3),
				equal:  generichelper.DeepEqual[int],
			},
			want: VarAll(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionEq(tt.args.first, tt.args.second, tt.args.equal)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("UnionEq() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestUnionEq_string(t *testing.T) {
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
		{name: "UnionWithCaseInsensitiveComparerEq",
			args: args{
				first:  VarAll("a", "b", "B", "c", "b"),
				second: VarAll("d", "e", "d", "a"),
				equal:  caseInsensitiveEqual,
			},
			want: VarAll("a", "b", "c", "d", "e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionEq(tt.args.first, tt.args.second, tt.args.equal)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("UnionEq() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestUnionCmp_int(t *testing.T) {
	ii1 := VarAll(1, 2, 3, 4)
	ii2 := VarAll(1, 2, 3, 4)
	ii3 := VarAll(1, 2, 3, 4)
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
		{name: "UnionWithIntComparer1",
			args: args{
				first:   VarAll(1, 2, 2),
				second:  Empty[int](),
				compare: cmp.Compare[int],
			},
			want: VarAll(1, 2),
		},
		{name: "UnionWithIntComparer2",
			args: args{
				first:   VarAll(1, 2),
				second:  VarAll(2, 3),
				compare: cmp.Compare[int],
			},
			want: VarAll(1, 2, 3),
		},
		{name: "SameEnumerable1",
			args: args{
				first:   ii1,
				second:  ii1,
				compare: cmp.Compare[int],
			},
			want: VarAll(1, 2, 3, 4),
		},
		{name: "SameEnumerable2",
			args: args{
				first:   errorhelper.Must(Skip(ii2, 2)),
				second:  errorhelper.Must(Take(ii2, 1)),
				compare: cmp.Compare[int],
			},
			want: VarAll(3, 4, 1),
		},
		{name: "SameEnumerable3",
			args: args{
				first:   errorhelper.Must(Skip(ii3, 2)),
				second:  ii3,
				compare: cmp.Compare[int],
			},
			want: VarAll(3, 4, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionCmp(tt.args.first, tt.args.second, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("UnionCmp() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestUnionCmp_string(t *testing.T) {
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
		{name: "UnionWithCaseInsensitiveComparerCmp",
			args: args{
				first:   VarAll("a", "b", "B", "c", "b"),
				second:  VarAll("d", "e", "d", "a"),
				compare: caseInsensitiveCompare,
			},
			want: VarAll("a", "b", "c", "d", "e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := UnionCmp(tt.args.first, tt.args.second, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("UnionCmp() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func ExampleUnion() {
	union, _ := Union(VarAll(5, 3, 9, 7, 5, 9, 3, 7), VarAll(8, 3, 6, 4, 4, 9, 1, 0))
	for num := range union {
		fmt.Printf("%d ", num)
	}
	// Output:
	// 5 3 9 7 8 6 4 1 0
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func ExampleUnionEq() {
	store1 := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
	}
	store2 := []Product{
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
	}
	//Get the products from the both arrays excluding duplicates.
	equal := func(p1, p2 Product) bool { return p1.Code == p2.Code && p1.Name == p2.Name }
	unionEq, _ := UnionEq(SliceAll(store1), SliceAll(store2), equal)
	for product := range unionEq {
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// apple 9
	// orange 4
	// lemon 12
}
