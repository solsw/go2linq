package go2linq

import (
	"fmt"
	"strings"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ExceptTest.cs

func TestExceptMust_int(t *testing.T) {
	i4 := NewEnSlice(1, 2, 3, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "IntWithoutComparer",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second: NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "IdenticalEnumerable",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4),
				second: NewEnSlice(1, 2, 3, 4),
			},
			want: Empty[int](),
		},
		{name: "IdenticalEnumerable2",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4),
				second: SkipMust(NewEnSlice(1, 2, 3, 4), 2),
			},
			want: NewEnSlice(1, 2),
		},
		{name: "SameEnumerable",
			args: args{
				first:  i4,
				second: SkipMust(i4, 2),
			},
			want: NewEnSlice(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestExceptMust_string(t *testing.T) {
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
				first:  NewEnSlice("A", "a", "b", "c", "b", "c"),
				second: NewEnSlice("b", "a", "d", "a"),
			},
			want: NewEnSlice("A", "c"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#except-and-exceptby
		{name: "Except",
			args: args{
				first:  NewEnSlice("Mercury", "Venus", "Earth", "Jupiter"),
				second: NewEnSlice("Mercury", "Earth", "Mars", "Jupiter"),
			},
			want: NewEnSlice("Venus"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestExceptEqMust_int(t *testing.T) {
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
		{name: "IntComparerSpecified",
			args: args{
				first:   NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second:  NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				equaler: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptEqMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestExceptEqMust_string(t *testing.T) {
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
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:   NewEnSlice("A", "a", "b", "c", "b"),
				second:  NewEnSlice("b", "a", "d", "a"),
				equaler: collate.CaseInsensitiveEqualer,
			},
			want: NewEnSlice("c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptEqMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestExceptCmpMust_int(t *testing.T) {
	i4 := NewEnSlice(1, 2, 3, 4)
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
		{name: "IntComparerSpecified",
			args: args{
				first:    NewEnSlice(1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
				second:   NewEnSlice(4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10),
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "SameEnumerable",
			args: args{
				first:    i4,
				second:   SkipMust(i4, 2),
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptCmpMust(tt.args.first, tt.args.second, tt.args.comparer)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestExceptCmpMust_string(t *testing.T) {
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
		{name: "CaseInsensitiveComparerSpecified",
			args: args{
				first:    NewEnSlice("A", "a", "b", "c", "b"),
				second:   NewEnSlice("b", "a", "d", "a"),
				comparer: collate.CaseInsensitiveComparer,
			},
			want: NewEnSlice("c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExceptCmpMust(tt.args.first, tt.args.second, tt.args.comparer)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ExceptCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the first example from Enumerable.Except help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExampleExceptMust() {
	numbers1 := NewEnSlice(2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5)
	numbers2 := NewEnSlice(2.2)
	except := ExceptMust(numbers1, numbers2)
	enr := except.GetEnumerator()
	for enr.MoveNext() {
		number := enr.Current()
		fmt.Println(number)
	}
	// Output:
	// 2
	// 2.1
	// 2.3
	// 2.4
	// 2.5
}

// see the last two examples from Enumerable.Except help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExampleExceptEqMust() {
	fruits1 := NewEnSlice(
		Product{Name: "apple", Code: 9},
		Product{Name: "orange", Code: 4},
		Product{Name: "lemon", Code: 12},
	)
	fruits2 := NewEnSlice(
		Product{Name: "APPLE", Code: 9},
	)
	var equaler collate.Equaler[Product] = collate.EqualerFunc[Product](
		func(p1, p2 Product) bool {
			return p1.Code == p2.Code && strings.EqualFold(p1.Name, p2.Name)
		},
	)
	//Get all the elements from the first array except for the elements from the second array.
	except := ExceptEqMust(fruits1, fruits2, equaler)
	enr := except.GetEnumerator()
	for enr.MoveNext() {
		product := enr.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// orange 4
	// lemon 12
}
