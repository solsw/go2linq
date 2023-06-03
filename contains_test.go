package go2linq

import (
	"fmt"
	"strings"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ContainsTest.cs

func TestContainsMust_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
		value  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchNoComparer",
			args: args{
				source: NewEnSlice("foo", "bar", "baz"),
				value:  "BAR",
			},
			want: false,
		},
		{name: "MatchNoComparer",
			args: args{
				source: NewEnSlice("foo", "bar", "baz"),
				value:  strings.ToLower("BAR"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsMust(tt.args.source, tt.args.value)
			if got != tt.want {
				t.Errorf("ContainsMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEqMust_string(t *testing.T) {
	type args struct {
		source  Enumerable[string]
		value   string
		equaler collate.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchWithCustomComparer",
			args: args{
				source:  NewEnSlice("foo", "bar", "baz"),
				value:   "gronk",
				equaler: collate.CaseInsensitiveEqualer,
			},
			want: false,
		},
		{name: "MatchWithCustomComparer",
			args: args{
				source:  NewEnSlice("foo", "bar", "baz"),
				value:   "BAR",
				equaler: collate.CaseInsensitiveEqualer,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsEqMust(tt.args.source, tt.args.value, tt.args.equaler)
			if got != tt.want {
				t.Errorf("ContainsEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEqMust_int(t *testing.T) {
	type args struct {
		source  Enumerable[int]
		value   int
		equaler collate.Equaler[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ImmediateReturnWhenMatchIsFound",
			args: args{
				source:  NewEnSlice(10, 1, 5, 0),
				value:   2,
				equaler: collate.EqualerFunc[int](func(i1, i2 int) bool { return i1 == 10/i2 }),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsEqMust(tt.args.source, tt.args.value, tt.args.equaler)
			if got != tt.want {
				t.Errorf("ContainsEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see the first example from Enumerable.Contains help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func ExampleContainsMust_ex1() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	fruit := "mango"
	hasMango := ContainsMust(NewEnSliceEn(fruits...), fruit)
	var what string
	if hasMango {
		what = "does"
	} else {
		what = "does not"
	}
	fmt.Printf("The array %s contain '%s'.\n", what, fruit)
	// Output:
	// The array does contain 'mango'.
}

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#query-expression-syntax-examples
// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#contains
func ExampleContainsMust_ex2() {
	markets := []Market{
		{Name: "Emily's", Items: []string{"kiwi", "cheery", "banana"}},
		{Name: "Kim's", Items: []string{"melon", "mango", "olive"}},
		{Name: "Adam's", Items: []string{"kiwi", "apple", "orange"}},
	}
	where := WhereMust(
		NewEnSliceEn(markets...),
		func(m Market) bool {
			return ContainsMust(NewEnSliceEn(m.Items...), "kiwi")
		},
	)
	names := SelectMust(where, func(m Market) string { return m.Name })
	enr := names.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Printf("%s market\n", name)
	}
	// Output:
	// Emily's market
	// Adam's market
}

// see the second example from Enumerable.Contains help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func ExampleContainsEqMust() {
	fruits := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "lemon", Code: 12},
	}
	apple := Product{Name: "apple", Code: 9}
	kiwi := Product{Name: "kiwi", Code: 8}
	var equaler collate.Equaler[Product] = collate.EqualerFunc[Product](
		func(p1, p2 Product) bool {
			return p1.Code == p2.Code && p1.Name == p2.Name
		},
	)
	hasApple := ContainsEqMust(NewEnSliceEn(fruits...), apple, equaler)
	hasKiwi := ContainsEqMust(NewEnSliceEn(fruits...), kiwi, equaler)
	fmt.Printf("Apple? %t\n", hasApple)
	fmt.Printf("Kiwi? %t\n", hasKiwi)
	// Output:
	// Apple? true
	// Kiwi? false
}
