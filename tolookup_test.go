//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToLookupTest.cs

func Test_ToLookupMust_string_int(t *testing.T) {
	lk1 := newLookup[int, string]()
	lk1.add(3, "abc")
	lk1.add(3, "def")
	lk1.add(1, "x")
	lk1.add(1, "y")
	lk1.add(3, "ghi")
	lk1.add(1, "z")
	lk1.add(2, "00")
	type args struct {
		source      Enumerable[string]
		keySelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want *Lookup[int, string]
	}{
		{name: "LookupWithNoComparerOrElementSelector",
			args: args{
				source:      NewEnSlice("abc", "def", "x", "y", "ghi", "z", "00"),
				keySelector: func(s string) int { return len(s) },
			},
			want: lk1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupMust(tt.args.source, tt.args.keySelector)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToLookupMust_string_string(t *testing.T) {
	lk2 := newLookup[string, string]()
	lk2.add("abc", "abc")
	lk2.add("def", "def")
	lk2.add("ABC", "ABC")
	type args struct {
		source      Enumerable[string]
		keySelector func(string) string
	}
	tests := []struct {
		name string
		args args
		want *Lookup[string, string]
	}{
		{name: "LookupWithNilComparerButNoElementSelector",
			args: args{
				source:      NewEnSlice("abc", "def", "ABC"),
				keySelector: Identity[string],
			},
			want: lk2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupMust(tt.args.source, tt.args.keySelector)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookupMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToLookupSelMust(t *testing.T) {
	lk := newLookup[int, string]()
	lk.add(3, "a")
	lk.add(3, "d")
	lk.add(1, "x")
	lk.add(1, "y")
	lk.add(3, "g")
	lk.add(1, "z")
	lk.add(2, "0")
	type args struct {
		source          Enumerable[string]
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
				source:          NewEnSlice("abc", "def", "x", "y", "ghi", "z", "00"),
				keySelector:     func(s string) int { return len(s) },
				elementSelector: func(s string) string { return string(s[0]) },
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupSelMust(tt.args.source, tt.args.keySelector, tt.args.elementSelector)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookupSelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToLookupEqMust(t *testing.T) {
	lk := newLookup[string, string]()
	lk.add("abc", "abc")
	lk.add("def", "def")
	lk.add("abc", "ABC")
	type args struct {
		source      Enumerable[string]
		keySelector func(string) string
		equaler     Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want *Lookup[string, string]
	}{
		{name: "LookupWithComparerButNoElementSelector",
			args: args{
				source:      NewEnSlice("abc", "def", "ABC"),
				keySelector: Identity[string],
				equaler:     CaseInsensitiveEqualer,
			},
			want: lk},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLookupEqMust(tt.args.source, tt.args.keySelector, tt.args.equaler)
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookupEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see LookupExample from Lookup Class help
// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2#examples
func ExampleToLookupSelMust() {
	// Create a list of Packages to put into a Lookup data structure.
	packages := NewEnSlice(
		Package{Company: "Coho Vineyard", Weight: 25.2, TrackingNumber: 89453312},
		Package{Company: "Lucerne Publishing", Weight: 18.7, TrackingNumber: 89112755},
		Package{Company: "Wingtip Toys", Weight: 6.0, TrackingNumber: 299456122},
		Package{Company: "Contoso Pharmaceuticals", Weight: 9.3, TrackingNumber: 670053128},
		Package{Company: "Wide World Importers", Weight: 33.8, TrackingNumber: 4665518773},
	)
	// Create a Lookup to organize the packages.
	// Use the first character of Company as the key value.
	// Select Company appended to TrackingNumber for each element value in the Lookup.
	lookup := ToLookupSelMust(packages,
		func(p Package) rune {
			return []rune(p.Company)[0]
		},
		func(p Package) string {
			return fmt.Sprintf("%s %d", p.Company, p.TrackingNumber)
		},
	)

	// Iterate through each Grouping in the Lookup and output the contents.
	enrLookup := lookup.GetEnumerator()
	for enrLookup.MoveNext() {
		packageGroup := enrLookup.Current()
		// Print the key value of the Grouping.
		fmt.Println(string(packageGroup.Key()))
		// Iterate through each value in the Grouping and print its value.
		enGr := packageGroup.GetEnumerator()
		for enGr.MoveNext() {
			str := enGr.Current()
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
	enrCgroup := cgroup.GetEnumerator()
	for enrCgroup.MoveNext() {
		str := enrCgroup.Current()
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
