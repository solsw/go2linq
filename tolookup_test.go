package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/generichelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToLookupTest.cs

func TestToLookup_string_int(t *testing.T) {
	lk := &Lookup[int, string]{KeyEqual: generichelper.DeepEqual[int]}
	lk.Add(3, "abc")
	lk.Add(3, "def")
	lk.Add(1, "x")
	lk.Add(1, "y")
	lk.Add(3, "ghi")
	lk.Add(1, "z")
	lk.Add(2, "00")
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want *Lookup[int, string]
	}{
		{name: "EmptySource",
			args: args{
				source:      Empty[string](),
				keySelector: func(s string) int { return len(s) },
			},
			want: &Lookup[int, string]{},
		},
		{name: "LookupWithNoComparerOrElementSelector",
			args: args{
				source:      VarAll("abc", "def", "x", "y", "ghi", "z", "00"),
				keySelector: func(s string) int { return len(s) },
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToLookup(tt.args.source, tt.args.keySelector)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLookup_string_string(t *testing.T) {
	lk := &Lookup[string, string]{KeyEqual: generichelper.DeepEqual[string]}
	lk.Add("abc", "abc")
	lk.Add("def", "def")
	lk.Add("ABC", "ABC")
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) string
	}
	tests := []struct {
		name string
		args args
		want *Lookup[string, string]
	}{
		{name: "LookupWithNilComparerButNoElementSelector",
			args: args{
				source:      VarAll("abc", "def", "ABC"),
				keySelector: Identity[string],
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToLookup(tt.args.source, tt.args.keySelector)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLookupEq_string_string(t *testing.T) {
	lk := Lookup[string, string]{KeyEqual: generichelper.DeepEqual[string]}
	lk.Add("abc", "abc")
	lk.Add("def", "def")
	lk.Add("abc", "ABC")
	type args struct {
		source      iter.Seq[string]
		keySelector func(string) string
		equal       func(string, string) bool
	}
	tests := []struct {
		name string
		args args
		want *Lookup[string, string]
	}{
		{name: "LookupWithComparerButNoElementSelector",
			args: args{
				source:      VarAll("abc", "def", "ABC"),
				keySelector: Identity[string],
				equal:       caseInsensitiveEqual,
			},
			want: &lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToLookupEq(tt.args.source, tt.args.keySelector, tt.args.equal)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookupEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLookupSel_string_int_string(t *testing.T) {
	lk := Lookup[int, string]{KeyEqual: generichelper.DeepEqual[int]}
	lk.Add(3, "a")
	lk.Add(3, "d")
	lk.Add(1, "x")
	lk.Add(1, "y")
	lk.Add(3, "g")
	lk.Add(1, "z")
	lk.Add(2, "0")
	type args struct {
		source          iter.Seq[string]
		keySelector     func(string) int
		elementSelector func(string) string
	}
	tests := []struct {
		name string
		args args
		want *Lookup[int, string]
	}{
		{name: "LookupWithElementSelectorButNoComparer",
			args: args{
				source:          VarAll("abc", "def", "x", "y", "ghi", "z", "00"),
				keySelector:     func(s string) int { return len(s) },
				elementSelector: func(s string) string { return string(s[0]) },
			},
			want: &lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToLookupSel(tt.args.source, tt.args.keySelector, tt.args.elementSelector)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookupSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LookupExample from
// https://learn.microsoft.com/dotnet/api/system.linq.Lookup-2#examples
func ExampleToLookupSel() {
	// Create a list of Packages to put into a Lookup data structure.
	packages := []Package{
		{Company: "Coho Vineyard", Weight: 25.2, TrackingNumber: 89453312},
		{Company: "Lucerne Publishing", Weight: 18.7, TrackingNumber: 89112755},
		{Company: "Wingtip Toys", Weight: 6.0, TrackingNumber: 299456122},
		{Company: "Contoso Pharmaceuticals", Weight: 9.3, TrackingNumber: 670053128},
		{Company: "Wide World Importers", Weight: 33.8, TrackingNumber: 4665518773},
	}
	// Create a Lookup to organize the packages.
	// Use the first character of Company as the key value.
	// Select Company appended to TrackingNumber for each element value in the Lookup.
	lookup, _ := ToLookupSel(
		SliceAll(packages),
		func(p Package) rune {
			return []rune(p.Company)[0]
		},
		func(p Package) string {
			return fmt.Sprintf("%s %d", p.Company, p.TrackingNumber)
		},
	)

	// Iterate through each Grouping in the Lookup and output the contents.
	for _, packageGroup := range lookup.groupings {
		// Print the key value of the Grouping.
		fmt.Println(string(packageGroup.Key()))
		// Iterate through each value in the Grouping and print its value.
		for str := range packageGroup.Values() {
			fmt.Printf("    %s\n", str)
		}
	}

	// Get the number of key-collection pairs in the Lookup.
	count := lookup.Count()
	fmt.Printf("\n%d\n", count)

	// Select a collection of Packages by indexing directly into the Lookup.
	cgroup := lookup.Item('C')
	// Output the results.
	fmt.Println("\nPackages that have a key of 'C'")
	for str := range cgroup {
		fmt.Println(str)
	}

	// Determine if there is a key with the value 'G' in the Lookup.
	hasG := lookup.Contains('G')
	fmt.Printf("\n%t\n", hasG)
	// Output:
	// C
	//     Coho Vineyard 89453312
	//     Contoso Pharmaceuticals 670053128
	// L
	//     Lucerne Publishing 89112755
	// W
	//     Wingtip Toys 299456122
	//     Wide World Importers 4665518773
	//
	// 3
	//
	// Packages that have a key of 'C'
	// Coho Vineyard 89453312
	// Contoso Pharmaceuticals 670053128
	//
	// false
}
