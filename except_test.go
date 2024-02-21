package go2linq

import (
	"cmp"
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ExceptTest.cs

func TestExcept_int(t *testing.T) {
	i4 := VarAll(1, 2, 3, 4)
	type args struct {
		first  iter.Seq[int]
		second iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "IntWithoutComparer",
			args: args{
				first:  VarAll(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: VarAll(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: VarAll(1, 2, 3),
		},
		{name: "IdenticalEnumerable",
			args: args{
				first:  VarAll(1, 2, 3, 4),
				second: VarAll(1, 2, 3, 4),
			},
			want: Empty[int](),
		},
		{name: "IdenticalEnumerable2",
			args: args{
				first:  VarAll(1, 2, 3, 4),
				second: errorhelper.Must(Skip(VarAll(1, 2, 3, 4), 2)),
			},
			want: VarAll(1, 2),
		},
		{name: "SameEnumerable",
			args: args{
				first:  i4,
				second: errorhelper.Must(Skip(i4, 2)),
			},
			want: VarAll(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Except(tt.args.first, tt.args.second)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Except() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestExcept_string(t *testing.T) {
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
				first:  VarAll("A", "a", "b", "c", "b", "c"),
				second: VarAll("b", "a", "d", "a"),
			},
			want: VarAll("A", "c"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#except-and-exceptby
		{name: "Except",
			args: args{
				first:  VarAll("Mercury", "Venus", "Earth", "Jupiter"),
				second: VarAll("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: VarAll("Venus"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Except(tt.args.first, tt.args.second)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Except() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestExceptEq_int(t *testing.T) {
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
				first:  VarAll(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: VarAll(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				equal:  generichelper.DeepEqual[int],
			},
			want: VarAll(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ExceptEq(tt.args.first, tt.args.second, tt.args.equal)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("ExceptEq() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestExceptEq_string(t *testing.T) {
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
				first:  VarAll("A", "a", "b", "c", "b"),
				second: VarAll("b", "a", "d", "a"),
				equal:  CaseInsensitiveEqual,
			},
			want: VarAll("c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ExceptEq(tt.args.first, tt.args.second, tt.args.equal)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("ExceptEq() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestExceptCmp_int(t *testing.T) {
	i4 := VarAll(1, 2, 3, 4)
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
				first:   VarAll(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second:  VarAll(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				compare: cmp.Compare[int],
			},
			want: VarAll(1, 2, 3),
		},
		{name: "SameEnumerable",
			args: args{
				first:   i4,
				second:  errorhelper.Must(Skip(i4, 2)),
				compare: cmp.Compare[int],
			},
			want: VarAll(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ExceptCmp(tt.args.first, tt.args.second, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("ExceptCmp() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestExceptCmp_string(t *testing.T) {
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
				first:   VarAll("A", "a", "b", "c", "b"),
				second:  VarAll("b", "a", "d", "a"),
				compare: CaseInsensitiveCompare,
			},
			want: VarAll("c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ExceptCmp(tt.args.first, tt.args.second, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("ExceptCmp() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExampleExcept() {
	numbers1 := VarAll(2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5)
	numbers2 := VarAll(2.2)
	except, _ := Except(numbers1, numbers2)
	for number := range except {
		fmt.Println(number)
	}
	// Output:
	// 2
	// 2.1
	// 2.3
	// 2.4
	// 2.5
}

// last two examples from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExampleExceptEq() {
	fruits1 := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "lemon", Code: 12},
	}
	fruits2 := []Product{
		{Name: "APPLE", Code: 9},
	}
	var equal = func(p1, p2 Product) bool {
		return p1.Code == p2.Code && strings.EqualFold(p1.Name, p2.Name)
	}
	// Get all the elements from the first array exceptEq for the elements from the second array.
	exceptEq, _ := ExceptEq(
		VarAll(fruits1...),
		VarAll(fruits2...),
		equal,
	)
	for product := range exceptEq {
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// orange 4
	// lemon 12
}
