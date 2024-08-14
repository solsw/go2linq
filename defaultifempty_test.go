package go2linq

import (
	"fmt"
	"iter"
	"slices"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DefaultIfEmptyTest.cs

func TestDefaultIfEmpty_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "EmptySequenceNoDefaultValue",
			args: args{
				source: Empty[int](),
			},
			want: VarToSeq(0),
		},
		{name: "NonEmptySequenceNoDefaultValue",
			args: args{
				source: VarToSeq(3, 1, 4),
			},
			want: VarToSeq(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DefaultIfEmpty(tt.args.source)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DefaultIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultIfEmptyDef_int(t *testing.T) {
	type args struct {
		source       iter.Seq[int]
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "EmptySequenceWithDefaultValue",
			args: args{
				source:       Empty[int](),
				defaultValue: 5,
			},
			want: VarToSeq(5),
		},
		{name: "NonEmptySequenceWithDefaultValue",
			args: args{
				source:       VarToSeq(3, 1, 4),
				defaultValue: 5,
			},
			want: VarToSeq(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DefaultIfEmptyDef(tt.args.source, tt.args.defaultValue)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DefaultIfEmptyDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func ExampleDefaultIfEmpty_ex1() {
	numbers, _ := DefaultIfEmpty(slices.Values([]int{}))
	for number := range numbers {
		fmt.Println(number)
	}
	// Output:
	// 0
}

// DefaultIfEmptyEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func ExampleDefaultIfEmpty_ex2() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	defaultIfEmpty, _ := DefaultIfEmpty(slices.Values(pets))
	for pet := range defaultIfEmpty {
		fmt.Println(pet.Name)
	}
	// Output:
	// Barley
	// Boots
	// Whiskers
}

// DefaultIfEmptyEx2 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func ExampleDefaultIfEmptyDef() {
	defaultPet := Pet{Name: "Default Pet", Age: 0}
	pets1 := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	defaultIfEmptyDef1, _ := DefaultIfEmptyDef(slices.Values(pets1), defaultPet)
	for pet := range defaultIfEmptyDef1 {
		fmt.Printf("Name: %s\n", pet.Name)
	}
	pets2 := []Pet{}
	defaultIfEmptyDef2, _ := DefaultIfEmptyDef(slices.Values(pets2), defaultPet)
	for pet := range defaultIfEmptyDef2 {
		fmt.Printf("\nName: %s\n", pet.Name)
	}
	// Output:
	// Name: Barley
	// Name: Boots
	// Name: Whiskers
	//
	// Name: Default Pet
}
