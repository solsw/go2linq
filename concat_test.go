package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ConcatTest.cs

func TestConcatMust_int(t *testing.T) {
	i4 := NewEnSlice(1, 2, 3, 4)
	rg := RangeMust(1, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "Empty",
			args: args{
				first:  Empty[int](),
				second: Empty[int](),
			},
			want: Empty[int](),
		},
		{name: "SemiEmpty",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4),
				second: Empty[int](),
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SimpleConcatenation",
			args: args{
				first:  NewEnSlice(1, 2, 3, 4),
				second: NewEnSlice(1, 2, 3, 4),
			},
			want: NewEnSlice(1, 2, 3, 4, 1, 2, 3, 4),
		},
		{name: "SimpleConcatenation2",
			args: args{
				first:  RangeMust(1, 2),
				second: RepeatMust(3, 1),
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "SameEnumerableInt",
			args: args{
				first:  i4,
				second: i4,
			},
			want: NewEnSlice(1, 2, 3, 4, 1, 2, 3, 4),
		},
		{name: "SameEnumerableInt2",
			args: args{
				first:  TakeMust(rg, 2),
				second: SkipMust(rg, 2),
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConcatMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ConcatMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestConcatMust_int2(t *testing.T) {
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "SecondSequenceIsntAccessedBeforeFirstUse",
			args: args{
				first: NewEnSlice(1, 2, 3, 4),
				second: SelectMust(
					Enumerable[int](NewEnSlice(0, 1)),
					func(x int) int { return 2 / x },
				),
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "NotNeededElementsAreNotAccessed",
			args: args{
				first: NewEnSlice(1, 2, 3),
				second: SelectMust(
					Enumerable[int](NewEnSlice(1, 0)),
					func(x int) int { return 2 / x },
				),
			},
			want: NewEnSlice(1, 2, 3, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeMust(ConcatMust(tt.args.first, tt.args.second), 4)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestConcatMust_string(t *testing.T) {
	rs := SkipMust(RepeatMust("q", 2), 1)
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "Empty",
			args: args{
				first:  Empty[string](),
				second: Empty[string](),
			},
			want: Empty[string](),
		},
		{name: "SemiEmpty",
			args: args{
				first:  Empty[string](),
				second: NewEnSlice("one", "two", "three", "four"),
			},
			want: NewEnSlice("one", "two", "three", "four"),
		},
		{name: "SimpleConcatenation",
			args: args{
				first:  NewEnSlice("a", "b"),
				second: NewEnSlice("c", "d"),
			},
			want: NewEnSlice("a", "b", "c", "d"),
		},
		{name: "SameEnumerableString",
			args: args{
				first:  rs,
				second: rs,
			},
			want: NewEnSlice("q", "q"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConcatMust(tt.args.first, tt.args.second)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ConcatMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see ConcatEx1 example from Enumerable.Concat help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.concat#examples
func ExampleConcatMust() {
	cats := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	dogs := []Pet{
		{Name: "Bounder", Age: 3},
		{Name: "Snoopy", Age: 14},
		{Name: "Fido", Age: 9},
	}
	query := ConcatMust(
		SelectMust(
			NewEnSliceEn(cats...),
			func(cat Pet) string { return cat.Name },
		),
		SelectMust(
			NewEnSliceEn(dogs...),
			func(dog Pet) string { return dog.Name },
		),
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Println(name)
	}
	// Output:
	// Barley
	// Boots
	// Whiskers
	// Bounder
	// Snoopy
	// Fido
}
