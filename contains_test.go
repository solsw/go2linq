package go2linq

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ContainsTest.cs

func TestContains_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
		value  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchNoComparer",
			args: args{
				source: VarAll("foo", "bar", "baz"),
				value:  "BAR",
			},
			want: false,
		},
		{name: "MatchNoComparer",
			args: args{
				source: VarAll("foo", "bar", "baz"),
				value:  strings.ToLower("BAR"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Contains(tt.args.source, tt.args.value)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEq_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
		value  string
		equal  func(string, string) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "NoMatchWithCustomComparer",
			args: args{
				source: VarAll("foo", "bar", "baz"),
				value:  "gronk",
				equal:  caseInsensitiveEqual,
			},
			want: false,
		},
		{name: "MatchWithCustomComparer",
			args: args{
				source: VarAll("foo", "bar", "baz"),
				value:  "BAR",
				equal:  caseInsensitiveEqual,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ContainsEq(tt.args.source, tt.args.value, tt.args.equal)
			if got != tt.want {
				t.Errorf("ContainsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEq_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
		value  int
		equal  func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ImmediateReturnWhenMatchIsFound",
			args: args{
				source: VarAll(10, 1, 5, 0),
				value:  2,
				equal:  func(i1, i2 int) bool { return 10/i1 == i2 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ContainsEq(tt.args.source, tt.args.value, tt.args.equal)
			if got != tt.want {
				t.Errorf("ContainsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func ExampleContains_ex1() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	fruit := "mango"
	hasMango, _ := Contains(SliceAll(fruits), fruit)
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
func ExampleContains_ex2() {
	markets := []Market{
		{Name: "Emily's", Items: []string{"kiwi", "cheery", "banana"}},
		{Name: "Kim's", Items: []string{"melon", "mango", "olive"}},
		{Name: "Adam's", Items: []string{"kiwi", "apple", "orange"}},
	}
	where, _ := Where(
		SliceAll(markets),
		func(m Market) bool {
			return errorhelper.Must(Contains(SliceAll(m.Items), "kiwi"))
		},
	)
	names, _ := Select(where, func(m Market) string { return m.Name })
	for name := range names {
		fmt.Printf("%s market\n", name)
	}
	// Output:
	// Emily's market
	// Adam's market
}

// second example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func ExampleContainsEq() {
	fruits := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "lemon", Code: 12},
	}
	apple := Product{Name: "apple", Code: 9}
	kiwi := Product{Name: "kiwi", Code: 8}
	var equal = func(p1, p2 Product) bool {
		return p1.Code == p2.Code && p1.Name == p2.Name
	}
	hasApple, _ := ContainsEq(VarAll(fruits...), apple, equal)
	hasKiwi, _ := ContainsEq(VarAll(fruits...), kiwi, equal)
	fmt.Printf("Apple? %t\n", hasApple)
	fmt.Printf("Kiwi? %t\n", hasKiwi)
	// Output:
	// Apple? true
	// Kiwi? false
}
