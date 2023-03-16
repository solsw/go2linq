package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DefaultIfEmptyTest.cs

func TestDefaultIfEmptyMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "EmptySequenceNoDefaultValue",
			args: args{
				source: Empty[int](),
			},
			want: NewEnSlice(0),
		},
		{name: "NonEmptySequenceNoDefaultValue",
			args: args{
				source: NewEnSlice(3, 1, 4),
			},
			want: NewEnSlice(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultIfEmptyMust(tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DefaultIfEmptyMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestDefaultIfEmptyDefMust_int(t *testing.T) {
	type args struct {
		source       Enumerable[int]
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "EmptySequenceWithDefaultValue",
			args: args{
				source:       Empty[int](),
				defaultValue: 5,
			},
			want: NewEnSlice(5),
		},
		{name: "NonEmptySequenceWithDefaultValue",
			args: args{
				source:       NewEnSlice(3, 1, 4),
				defaultValue: 5,
			},
			want: NewEnSlice(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultIfEmptyDefMust(tt.args.source, tt.args.defaultValue)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DefaultIfEmptyDefMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the last example from Enumerable.DefaultIfEmpty help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func ExampleDefaultIfEmptyMust_ex1() {
	numbers := DefaultIfEmptyMust(NewEnSlice([]int{}...))
	enr := numbers.GetEnumerator()
	for enr.MoveNext() {
		number := enr.Current()
		fmt.Println(number)
	}
	// Output:
	// 0
}

// see DefaultIfEmptyEx1 example from Enumerable.DefaultIfEmpty help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func ExampleDefaultIfEmptyMust_ex2() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	enr := DefaultIfEmptyMust(pets).GetEnumerator()
	for enr.MoveNext() {
		pet := enr.Current()
		fmt.Println(pet.Name)
	}
	// Output:
	// Barley
	// Boots
	// Whiskers
}

// see DefaultIfEmptyEx2 example from Enumerable.DefaultIfEmpty help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func ExampleDefaultIfEmptyDefMust() {
	defaultPet := Pet{Name: "Default Pet", Age: 0}
	pets1 := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	enr1 := DefaultIfEmptyDefMust(pets1, defaultPet).GetEnumerator()
	for enr1.MoveNext() {
		pet := enr1.Current()
		fmt.Printf("Name: %s\n", pet.Name)
	}
	pets2 := NewEnSlice([]Pet{}...)
	enr2 := DefaultIfEmptyDefMust(pets2, defaultPet).GetEnumerator()
	for enr2.MoveNext() {
		pet := enr2.Current()
		fmt.Printf("\nName: %s\n", pet.Name)
	}
	// Output:
	// Name: Barley
	// Name: Boots
	// Name: Whiskers
	//
	// Name: Default Pet
}
