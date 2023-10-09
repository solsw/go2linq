package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectTest.cs

func TestSelect_int_int(t *testing.T) {
	var count int
	type args struct {
		source   Enumerable[int]
		selector func(int) int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceThrowsNullArgumentException",
			args: args{
				source:   nil,
				selector: func(x int) int { return x + 1 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullProjectionThrowsNullArgumentException",
			args: args{
				source:   NewEnSlice(1, 3, 7, 9, 10),
				selector: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "SimpleProjection",
			args: args{
				source:   NewEnSlice(1, 5, 2),
				selector: func(x int) int { return x * 2 },
			},
			want: NewEnSlice(2, 10, 4),
		},
		{name: "EmptySource",
			args: args{
				source:   Empty[int](),
				selector: func(x int) int { return x * 2 },
			},
			want: Empty[int](),
		},
		{name: "SideEffectsInProjection1",
			args: args{
				source:   NewEnSlice(3, 2, 1), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "SideEffectsInProjection2",
			args: args{
				source:   NewEnSlice(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewEnSlice(4, 5, 6),
		},
		{name: "SideEffectsInProjection3",
			args: args{
				source:   NewEnSlice(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: NewEnSlice(11, 12, 13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Select() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Select() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
		if tt.name == "SideEffectsInProjection2" {
			count = 10
		}
	}
}

func TestSelectMust_int_string(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SimpleProjectionToDifferentType",
			args: args{
				source:   NewEnSlice(1, 5, 2),
				selector: func(x int) string { return fmt.Sprint(x) },
			},
			want: NewEnSlice("1", "5", "2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectMust_string_string(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#select
		{name: "Select",
			args: args{
				source:   NewEnSlice("an", "apple", "a", "day"),
				selector: func(s string) string { return string([]rune(s)[0]) },
			},
			want: NewEnSlice("a", "a", "a", "d"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectIdx_int_int(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "WithIndexNullSourceThrowsNullArgumentException",
			args: args{
				source:   nil,
				selector: func(x, idx int) int { return x + idx },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "WithIndexNullSelectorThrowsNullArgumentException",
			args: args{
				source:   NewEnSlice(1, 3, 7, 9, 10),
				selector: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "WithIndexSimpleProjection",
			args: args{
				source:   NewEnSlice(1, 5, 2),
				selector: func(x, idx int) int { return x + idx*10 },
			},
			want: NewEnSlice(1, 15, 22),
		},
		{name: "WithIndexEmptySource",
			args: args{
				source:   Empty[int](),
				selector: func(x, idx int) int { return x + idx },
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectIdx(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SelectIdx() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectIdx() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the first example from Enumerable.Select help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.select
func ExampleSelectMust_ex1() {
	squares := SelectMust(
		RangeMust(1, 10),
		func(x int) int { return x * x },
	)
	enr := squares.GetEnumerator()
	for enr.MoveNext() {
		num := enr.Current()
		fmt.Println(num)
	}
	// Output:
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
	// 49
	// 64
	// 81
	// 100
}

func ExampleSelectMust_ex2() {
	numbers := []string{"one", "two", "three", "four", "five"}
	fmt.Println(ToStringDef(
		SelectMust(
			NewEnSlice(numbers...),
			func(s string) string {
				return string(s[0]) + string(s[len(s)-1])
			},
		),
	))
	fmt.Println(ToStringDef(
		SelectMust(
			NewEnSlice(numbers...),
			func(s string) string {
				runes := []rune(s)
				reversedRunes := ToSliceMust(
					ReverseMust(
						NewEnSlice(runes...),
					),
				)
				return string(reversedRunes)
			},
		),
	))
	// Output:
	// [oe to te fr fe]
	// [eno owt eerht ruof evif]
}

// see the last example from Enumerable.Select help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.select

type indexstr struct {
	index int
	str   string
}

func ExampleSelectIdxMust() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	query := SelectIdxMust(
		NewEnSlice(fruits...),
		func(fruit string, index int) indexstr {
			return indexstr{index: index, str: fruit[:index]}
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		obj := enr.Current()
		fmt.Printf("%+v\n", obj)
	}
	// Output:
	// {index:0 str:}
	// {index:1 str:b}
	// {index:2 str:ma}
	// {index:3 str:ora}
	// {index:4 str:pass}
	// {index:5 str:grape}
}
