package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/EmptyTest.cs

func TestEmpty_int(t *testing.T) {
	tests := []struct {
		name string
		want Enumerable[int]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[int]()
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Empty() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestEmpty_string(t *testing.T) {
	tests := []struct {
		name string
		want Enumerable[string]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[string]()
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Empty() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the example from Enumerable.Empty help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.empty#examples
func ExampleEmpty() {
	names1 := NewEnSlice("Hartono, Tommy")
	names2 := NewEnSlice("Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	names3 := NewEnSlice("Solanki, Ajay", "Hoeing, Helge", "Andersen, Henriette Thaulow", "Potra, Cristina", "Iallo, Lucio")
	namesList := NewEnSlice(names1, names2, names3)
	allNames := AggregateSeedMust(namesList,
		Empty[string](),
		func(current, next Enumerable[string]) Enumerable[string] {
			// Only include arrays that have four or more elements
			if CountMust(next) > 3 {
				return UnionMust(current, next)
			}
			return current
		},
	)
	enr := allNames.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Println(name)
	}
	// Output:
	// Adams, Terry
	// Andersen, Henriette Thaulow
	// Hedlund, Magnus
	// Ito, Shu
	// Solanki, Ajay
	// Hoeing, Helge
	// Potra, Cristina
	// Iallo, Lucio
}
